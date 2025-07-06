import gymnasium as gym
from gymnasium import spaces
import numpy as np
from stable_baselines3 import PPO
import torch
import requests

class GoRPSRemoteEnv(gym.Env):
    def __init__(self, server_url="http://localhost:8080"):
        super().__init__()
        self.server_url = server_url
        self.action_space = spaces.Discrete(3)
        self.observation_space = spaces.Box(low=0, high=2, shape=(1,), dtype=np.int32)
        self.state = None

    def reset(self, *, seed=None, options=None):
        resp = requests.post(f"{self.server_url}/reset")
        obs = np.array([resp.json()["observation"]], dtype=np.int32)
        self.state = obs
        return obs, {}

    def step(self, action):
        resp = requests.post(f"{self.server_url}/step", json={"action": int(action)})
        data = resp.json()
        obs = np.array([data["observation"]], dtype=np.int32)
        reward = data["reward"]
        done = data["done"]
        info = {}
        return obs, reward, done, False, info


if __name__ == '__main__':
    # env = VisibleRPS()
    # model = PPO("MlpPolicy", env, verbose=1)
    # print("Training agent on VisibleRPS...\n")
    # model.learn(total_timesteps=10000)
    print("\nTesting the trained agent...")
    env = GoRPSRemoteEnv("http://localhost:8080")
    model = PPO("MlpPolicy", env, verbose=1)
    model.learn(total_timesteps=10000)

    obs, _ = env.reset()
    wins, draws, losses = 0, 0, 0
    action_names = ['Rock', 'Paper', 'Scissors']

    for i in range(20):
        action, _ = model.predict(obs, deterministic=True)
        obs, reward, done, truncated, info = env.step(int(action))
        opponent_action = obs[0]  # because env returns the new opponent move as obs after step

        # Compute result string for printout
        if reward == 1:
            result = "Win"
            wins += 1
        elif reward == 0:
            result = "Draw"
            draws += 1
        else:
            result = "Lose"
            losses += 1

        print(f"Round {i + 1}: Agent played {action_names[int(action)]}, "
              f"Opponent played {action_names[int(opponent_action)]}, "
              f"Result: {result}")

        obs, _ = env.reset()

    print(f"\nAgent results in 20 rounds: {wins} Wins, {draws} Draws, {losses} Losses")


    # After training
    device = next(model.policy.parameters()).device
    dummy_input = torch.as_tensor(np.array([[0]], dtype=np.int32)).to(device)
    torch.onnx.export(
        model.policy,
        dummy_input,
        "rps_agent_visible.onnx",
        input_names=["input"],
        output_names=["output"],
        dynamic_axes={"input": {0: "batch_size"}, "output": {0: "batch_size"}},
        opset_version=11,
    )
    print("\nModel exported as rps_agent_visible.onnx")
