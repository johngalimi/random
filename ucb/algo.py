import numpy as np
import random

"""
UCB = mean(xj) + sqrt(2log(t) / nj)

mean(xj) == avg payoff at jth step
t = total number of contacts that have entered the experiment (what round we're on)
nj = total number of contacts allocated to the jth variant

general flow -->
1. try each arm once
2. in round t, pick action = max(UCB)
"""

class Arm:
    def __init__(self, id):
        self.id = id
        self.payoff_probability = random.random()

    def pull(self):
        return int(random.random() < self.payoff_probability)

class UCB:
    def __init__(self):
        self.payoffs = []
        self.average_payoff = None
        self.allocation_quantity = None

    def update_payoffs(self, is_win):
        self.payoffs.append(is_win)

    def update_mean(self):
        self.average_payoff = sum(self.payoffs) / len(self.payoffs)
        
class Simulation:
    def __init__(self, rounds):
        self.rounds = rounds


if __name__ == '__main__':
    arm_a = Arm('a')
    arm_b = Arm('b')
    print('a ', arm_a.payoff_probability)
    print('b ', arm_b.payoff_probability)
    for i in range(10):
        print(arm_a.pull(), arm_b.pull())
