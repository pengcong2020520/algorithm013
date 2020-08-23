func numIslands(grid [][]byte) int {

    res := 0
    if len(grid) == 0 {
        return 0
    }
    for row := 0; row < len(grid); row++ {
        for col := 0; col < len(grid[0]); col++ {
            if grid[row][col] == '1' {
                res += 1
                dfs(row, col, &grid)
            }

        }
    }
    return res
}

func dfs(row, col int, grid *[][]byte) {
    var gx = [4]int{0, 1, 0, -1}
    var gy = [4]int{1, 0, -1, 0}
    if row > len(*grid)-1 || col > len((*grid)[0])-1{
        return
    }
    for i := 0; i < 4; i++ {
        if row + gx[i] >= 0 && col + gy[i] >= 0 && row+gx[i] < len(*grid) && col+gy[i] < len((*grid)[0]) && (*grid)[row+gx[i]][col+gy[i]] == '1' {
            (*grid)[row+gx[i]][col+gy[i]] = '0'
            dfs(row+gx[i], col+gy[i], grid)
        }
    }
}

