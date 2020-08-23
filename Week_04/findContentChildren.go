func findContentChildren(g []int, s []int) int {
    //1. 先对数组进行排序
    sort.Ints(g)
    sort.Ints(s)
    var i, j int
    for i < len(g) && j < len(s) {
        //2. 对每一步进行决策，如果满足胃口即匹配
        if g[i] <= s[j] {
            i++
        }
        j++
    }
    return i
}