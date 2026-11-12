from typing import Optional
class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        clones = {}
        def dfs(node):
            if node is None:
                return
            if node in clones:
                return clones[node]
            newNode = Node(node.val)
            clones[node] = newNode

            for nb in node.neighbors:
                newNode.neighbors.append(dfs(nb))
            return newNode
        return dfs(node)