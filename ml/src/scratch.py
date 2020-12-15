import operator
import numpy as np

inputs = [[1, 2, 3, 2.5], [2, 5, -1, 2], [-1.5, 2.7, 3.3, -0.8]]
weights = [[0.2, 0.8, -0.5, 1.0], [0.5, -0.91, 0.26, -0.5], [-0.26, -0.27, 0.17, 0.87]]
biases = [2, 3, 0.5]


def dot_product(vector_a, vector_b):
    if len(vector_a) != len(vector_b):
        raise Exception

    return sum(map(operator.mul, vector_a, vector_b))


# raw_dot = [
#     dot_product(inputs, neuron_weights) + neuron_bias
#     for neuron_weights, neuron_bias in zip(weights, biases)
# ]

# numpy determines how to index your return based on first arg
output = np.dot(weights, inputs) + biases
print(output)
