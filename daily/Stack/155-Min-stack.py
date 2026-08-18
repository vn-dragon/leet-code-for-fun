class MinStack:
    def __init__(self):
        self.stacks = []
        self.min_stacks = []

    def push(self, value: int) -> None:
        self.stacks.append(value)
        if not self.min_stacks or value <= self.min_stacks[-1]:
            self.min_stacks.append(value)

    def pop(self) -> None:
        topElement = self.stacks.pop()
        if topElement == self.min_stacks[-1]:
            self.min_stacks.pop()

    def top(self) -> int:
        return self.stacks[-1]

    def getMin(self) -> int:
        return self.min_stacks[-1]
