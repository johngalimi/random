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
