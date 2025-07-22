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