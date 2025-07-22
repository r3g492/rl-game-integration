import gymnasium as gym
from gymnasium import spaces
import numpy as np
import requests
from stable_baselines3 import PPO

OBS_KEYS = [
    "car_x", "car_y", "car_z",
    "velocity", "yaw",
    "goal_x", "goal_y", "goal_z"
]

class GoCarRemoteEnv(gym.Env):
    def __init__(self, server_url="http://localhost:8080"):
        super().__init__()
        self.server_url = server_url
        # Use the correct observation and action spaces!
        self.observation_space = spaces.Box(
            low=-np.inf, high=np.inf, shape=(8,), dtype=np.float32
        )
        self.action_space = spaces.Box(
            low=np.array([-1.0, -1.0]),  # Example values, adjust as needed
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
        return obs, {}

    def step(self, action):
        # action should be a 2-element array-like
        speed, rotation = float(action[0]), float(action[1])
        payload = {"speedGradient": speed, "rotationGradient": rotation}
        resp = requests.post(f"{self.server_url}/step", json=payload)
        data = resp.json()
        obs_dict = data["observation"]
        obs = np.array([obs_dict[k] for k in OBS_KEYS], dtype=np.float32)
        reward = data["reward"]
        done = data["done"]
        info = {}
        return obs, reward, done, False, info
if __name__ == '__main__':
    env = GoCarRemoteEnv("http://localhost:8080")
    model = PPO("MlpPolicy", env, verbose=1)
    model.learn(total_timesteps=100000)
    # (Optional: Export the model if needed)

    # Evaluation: Run N episodes and track success
    N_EPISODES = 10
    success_count = 0
    max_steps = 200  # Set max steps per episode to avoid infinite loops

    for ep in range(N_EPISODES):
        obs, _ = env.reset()
        done = False
        step_count = 0
        total_reward = 0
        print(f"\nEpisode {ep + 1} start")
        while not done and step_count < max_steps:
            action, _states = model.predict(obs, deterministic=True)
            obs, reward, done, truncated, info = env.step(action)
            total_reward += reward
            step_count += 1

            # Print every step (optional)
            print(f"  Step {step_count}: obs={obs}, reward={reward:.2f}, done={done}")
            # Optional: Add a delay for debugging/visualization
            # time.sleep(0.05)

        if done:
            success_count += 1
            print(
                f"  -> Episode {ep + 1} finished (success or done) in {step_count} steps, total reward {total_reward:.2f}")
        else:
            print(f"  -> Episode {ep + 1} timed out at {max_steps} steps, total reward {total_reward:.2f}")

    print(f"\nReached done in {success_count}/{N_EPISODES} episodes")