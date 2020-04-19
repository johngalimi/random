class Store:
    def __init__(self):
        self.kv = {}

    def get(self, key):
        return self.kv.get(key)

    def set(self, key, value):
        self.kv[key] = value
        return True

    def delete(self, key):
        if key in self.kv:
            del self.kv[key]
            return True
        return False

    def flush(self):
        self.kv.clear()
        return True

    def multi_get(self, keys):
        return [self.get(key) for key in keys]
 
    def multi_set(self, kv_pairs):
        for key, value in kv_pairs.items():
            self.set(key, value)
        return True

