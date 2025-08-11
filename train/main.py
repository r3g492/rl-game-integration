import gymnasium as gym
from gymnasium import spaces
import numpy as np
import requests
from stable_baselines3 import PPO
from stable_baselines3.common.callbacks import BaseCallback

OBS_KEYS = [
    "car_x", "car_y", "car_z",
    "velocity", "yaw",
    "goal_x", "goal_y", "goal_z"
]

_OBS_LOW  = np.array([-100, -100, -100, -100, -np.pi, -100, -100, -100], dtype=np.float32)
_OBS_HIGH = np.array([ 100,  100,  100,  100,  np.pi,  100,  100,  100], dtype=np.float32)

class GoCarRemoteEnv(gym.Env):
    def __init__(self, server_url="http://localhost:8080"):
        super().__init__()
        self.server_url = server_url
        self.session = requests.Session()
        self.observation_space = spaces.Box(
            low=_OBS_LOW, high=_OBS_HIGH, dtype=np.float32
        )
        self.action_space = spaces.Box(
            low=np.array([-1.0, -1.0], dtype=np.float32),
            high=np.array([ 1.0,  1.0], dtype=np.float32),
            shape=(2,), dtype=np.float32
        )
        self.state = None

    def _clip_obs(self, obs_vec: np.ndarray) -> np.ndarray:
        return np.clip(obs_vec, self.observation_space.low, self.observation_space.high)

    def reset(self, *, seed=None, options=None):
        resp = self.session.post(f"{self.server_url}/reset", timeout=2.0)
        data = resp.json()
        obs_dict = data["observation"]
        obs = np.array([obs_dict[k] for k in OBS_KEYS], dtype=np.float32)
        obs = self._clip_obs(obs)
        self.state = obs
        print(f"[ENV] Reset observation: {obs}")
        return obs, {}

    def step(self, action):
        speed, rotation = float(action[0]), float(action[1])
        payload = {"speedGradient": speed, "rotationGradient": rotation}
        resp = self.session.post(f"{self.server_url}/step", json=payload, timeout=2.0)
        data = resp.json()

        obs_dict = data["observation"]
        obs1 = np.array([obs_dict[k] for k in OBS_KEYS], dtype=np.float32)
        obs1 = self._clip_obs(obs1)

        reward1     = float(data["reward"])
        terminated1 = bool(data.get("done", False))        # success/failure
        truncated1  = bool(data.get("truncated", False))   # timeout/step cap
        info1 = {"is_success": bool(data.get("is_success", False))}

        return obs1, reward1, terminated1, truncated1, info1

class EpisodeRewardPrinterCallback(BaseCallback):
    def __init__(self, verbose=0):
        super().__init__(verbose)
        self.episode_reward = 0.0

    def _on_step(self) -> bool:
        reward = self.locals["rewards"]
        done = self.locals["dones"]  # True for terminated OR truncated
        if isinstance(reward, (list, np.ndarray)):
            reward = float(reward[0])
        if isinstance(done, (list, np.ndarray)):
            done = bool(done[0])
        self.episode_reward += reward
        if done:
            print(f"[TRAIN] Episode done, total reward: {self.episode_reward:.2f}")
            self.episode_reward = 0.0
        return True

if __name__ == '__main__':
    env = GoCarRemoteEnv("http://localhost:8080")

    policy_kwargs = dict(net_arch=[1024, 512, 256, 128])
    reward_callback = EpisodeRewardPrinterCallback()

    model = PPO("MlpPolicy", env, verbose=1, policy_kwargs=policy_kwargs, seed=42)
    model.learn(total_timesteps=100000, callback=reward_callback)

    # ----- quick test -----
    N_EPISODES = 10
    success_count = 0
    max_steps = 5000

    for ep in range(N_EPISODES):
        obs, _ = env.reset()
        done = False
        truncated = False
        step_count = 0
        total_reward = 0.0
        print(f"\nEpisode {ep + 1} start")
        while not (done or truncated) and step_count < max_steps:
            action, _states = model.predict(obs, deterministic=False)
            obs, reward, done, truncated, info = env.step(action)
            total_reward += float(reward)
            step_count += 1

        if info.get("is_success", False):
            success_count += 1
            print(f"  -> SUCCESS in {step_count} steps, total reward {total_reward:.2f}")
        elif truncated:
            print(f"  -> TIMEOUT at {step_count} steps, total reward {total_reward:.2f}")
        else:
            print(f"  -> TERMINATED (fail) in {step_count} steps, total reward {total_reward:.2f}")

    print(f"\nReached success in {success_count}/{N_EPISODES} episodes")
