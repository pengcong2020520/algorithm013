func maxProfit(prices []int) int {
    //  单笔交易收益最高 ： 遍历该数组比前一个数小就跳过，大就与之前的min数相比
    // length := len(prices)
    // res := 0
    // min := prices[0] 
    // pre := prices[0]
    // for i := 1; i < length; i++ {
    //     if prices[i] <= pre {
    //         pre = prices[i]
    //         continue
    //     }
    //     if min > pre {
    //         min, pre = pre, prices[i]
    //     } 
    //     res = prices[i] - min
    //     pre = prices[i]     
    // }
    // return res

    //多笔交易收益最高 ：只要前面的数比后面的大，就加上差值
    length := len(prices)
    res := 0
    pre := prices[0]
    for i := 1; i < length; i++ {
        if prices[i] >= pre {
            res += prices[i] - pre 
        }
        pre = prices[i]
    }
    return res
}