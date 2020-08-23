type st struct {}
func robotSim(commands []int, obstacles [][]int) int {
    var dx = [4]int{0, 1, 0, -1}
    var dy = [4]int{1, 0, -1, 0}
    var curdir int   //0代表北方向  1代表东 2代表南 3代表西
    var nx, ny, curx, cury, res int
    obsMap := make(map[string]st, len(obstacles))
	for i:=0; i < len(obstacles); i++{
		tmp := fmt.Sprintf("%v %v", obstacles[i][0], obstacles[i][1])
		obsMap[tmp] = st{}
	}
    for i := 0; i < len(commands); i++ {
        if commands[i] == -1 {
            curdir = (curdir + 1)%4
        } else if commands[i] == -2 {
            curdir = (curdir + 3)%4
        } else {
            for j := 0; j < commands[i]; j++ {
                nx, ny = curx + dx[curdir], cury + dy[curdir]
                nxy := fmt.Sprintf("%v %v", nx, ny)
                if _, ok := obsMap[nxy]; ok {
                    break
                }
                curx = nx
                cury = ny
            }
        }
    res = max(res, curx*curx + cury*cury)    
    }
    return res
}


func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
