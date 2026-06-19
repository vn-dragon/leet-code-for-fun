class Solution:
    def connect(self, root: 'Optional[Node]') -> 'Optional[Node]':
        if not root:
            return None
        
        def bfs(level_nodes):
            if not level_nodes:
                return
            next_level = []
            for i in range(len(level_nodes)):
                n = level_nodes[i]

                if n.left: 
                    next_level.append(n.left)
                if n.right: 
                    next_level.append(n.right)

                for i in range(len(next_level) - 1):
                    next_level[i].next = next_level[i + 1]

            bfs(next_level)
        bfs([root])
        return root