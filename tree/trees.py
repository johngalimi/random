class Node:
    def __init__(self, value):
        self.left = None
        self.right = None
        self.value = value

    def insert(self, value):
        print("*******", value)
        if self.value:
            print('yep value')
            if value < self.value:
                print('less than current value, go left', value, self.value)
                if self.left is None:
                    self.left = Node(value)
                else:
                    self.left.insert(value)
            elif value >= self.value:
                print('more than current value, go right', value, self.value)
                if self.right is None:
                    self.right = Node(value)
                else:
                    self.right.insert(value)
        else:
            self.value = value

    def visualize(self):
        if self.left:
            self.left.visualize()
        print(self.value)
        if self.right:
            self.right.visualize()

    def traverse_ordered(self, root_node):
        
        values = []
        
        if root_node:
            values = self.traverse_ordered(root_node.left)
            values.append(root_node.value)
            values += self.traverse_ordered(root_node.right)
        
        return values


if __name__ == '__main__':
    root = Node(28)
    root.insert(15)
    root.insert(35)
    root.insert(11)
    root.insert(18)
    root.insert(30)
    root.insert(41)
    print(root.traverse_ordered(root))
    root.visualize()