func nthUglyNumber(n int) int {
    var res int
    i, j, k := 1, 1, 1
    var temp []int
    temp = append(temp, 1)
    if n == 0 {
        return -1
    }
    if n == 1 {
        return 1
    }
    for count := 1; count < n; count++ {
        res = min(2*temp[i-1], 3*temp[j-1], 5*temp[k-1])
        temp = append(temp, res)
        if res == 2*temp[i-1] {
            i++
        }
        if res == 3*temp[j-1] {
            j++
        }
        if res == 5*temp[k-1] {

            k++
        }
    }
    return res
}

func min(x int, y int, z int) int {
    if x < y {
        if x < z {
            return x
        }
        return z
    }
    if  y > z {
        return z
    }
    return y
}