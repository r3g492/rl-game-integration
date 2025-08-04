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

class GoCarRemoteEnv(gym.Env):
    def __init__(self, server_url="http://localhost:8080"):
        super().__init__()
        self.server_url = server_url
        self.observation_space = spaces.Box(
            low=-np.inf, high=np.inf, shape=(8,), dtype=np.float32
        )
        self.action_space = spaces.Box(
            low=np.array([-1.0, -1.0]),
            high=np.array([1.0, 1.0]),
            shape=(2,),
            dtype=np.float32
        )
        self.state = None

    def reset(self, *, seed=None, options=None):
        resp = requests.post(f"{self.server_url}/reset")
        data = resp.json()
        obs_dict = data["observation"]
        obs = np.array([obs_dict[k] for k in OBS_KEYS], dtype=np.float32)
        self.state = obs
        print(f"[ENV] Reset observation: {obs}")
        return obs, {}

    def step(self, action):
        speed, rotation = float(action[0]), float(action[1])
        payload = {"speedGradient": speed, "rotationGradient": rotation}
        resp = requests.post(f"{self.server_url}/step", json=payload)
        data = resp.json()
        obs_dict = data["observation"]
        obs1 = np.array([obs_dict[k] for k in OBS_KEYS], dtype=np.float32)
        reward1 = data["reward"]
        done1 = data["done"]
        info1 = {}
        # print(f"[ENV] Step | Action: {action}, Reward: {reward1}, Done: {done1}")
        return obs1, reward1, done1, False, info1

# ----------- CALLBACK TO PRINT EPISODE REWARD -----------
class EpisodeRewardPrinterCallback(BaseCallback):
    def __init__(self, verbose=0):
        super().__init__(verbose)
        self.episode_reward = 0

    def _on_step(self) -> bool:
        reward = self.locals["rewards"]
        done = self.locals["dones"]

        # SB3 always wraps scalar envs as lists
        if isinstance(reward, (list, np.ndarray)):
            reward = reward[0]
        if isinstance(done, (list, np.ndarray)):
            done = done[0]

        self.episode_reward += reward
        if done:
            print(f"[TRAIN] Episode done, total reward: {self.episode_reward:.2f}")
            self.episode_reward = 0
        return True

if __name__ == '__main__':
    env = GoCarRemoteEnv("http://localhost:8080")

    policy_kwargs = dict(
        net_arch=[1024, 512, 256, 128]
    )

    # Create the callback instance
    reward_callback = EpisodeRewardPrinterCallback()

    # Train the model with the callback
    model = PPO("MlpPolicy", env, verbose=1, policy_kwargs=policy_kwargs, seed=42)
    model.learn(total_timesteps=1000000, callback=reward_callback)

    # ----- Optionally: test runs after training (unchanged) -----
    N_EPISODES = 10
    success_count = 0
    max_steps = 1000

    for ep in range(N_EPISODES):
        obs, _ = env.reset()
        done = False
        step_count = 0
        total_reward = 0
        print(f"\nEpisode {ep + 1} start")
        while not done and step_count < max_steps:
            action, _states = model.predict(obs, deterministic=False)
            obs, reward, done, truncated, info = env.step(action)
            total_reward += reward
            step_count += 1
            # Comment or remove the next line if you don't want per-step logs:
            # print(f"  Step {step_count}: obs={obs}, reward={reward:.2f}, done={done}")
        if done:
            success_count += 1
            print(
                f"  -> Episode {ep + 1} finished (success or done) in {step_count} steps, total reward {total_reward:.2f}")
        else:
            print(f"  -> Episode {ep + 1} timed out at {max_steps} steps, total reward {total_reward:.2f}")

    print(f"\nReached done in {success_count}/{N_EPISODES} episodes")