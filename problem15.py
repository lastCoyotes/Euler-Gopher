# PE 15 Lattice Paths
# assuming square grid, since the problem only asks for 20x20 grid
# can ONLY move RIGHT and DOWN
# can easily solve this with a formula from math but i want to utilize a backtracking approach
# all set of moves can only be length 2n.

def latticePaths(n: int) -> int:
    result = []  # keeping track of all options of moves jic i want it for somethin later
    # no actual grid is being made so we'll keep track with xy indices

    def bfs(path: list[str], x: int, y: int):
        if len(path) == 2*n:
            result.append(path.copy())
            print(path)
            return
        if x < n: # go right
            x += 1
            path.append('R')
            bfs(path, x, y)
            path.pop()
            x -= 1
        if y < n: # go down
            y += 1
            path.append('D')
            bfs(path, x, y)
            path.pop()
            y -= 1

    bfs([], 0, 0)
    return len(result)

# 40 choose 20 = 137_846_528_820

if __name__ == '__main__':
    print(latticePaths(4)) # 8 choose 4 = 70
