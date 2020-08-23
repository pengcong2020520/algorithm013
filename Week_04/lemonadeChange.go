func lemonadeChange(bills []int) bool {
    var five, ten int
    for _, v := range bills {
        if v == 5 {
            five++
        } else if (v == 10) {
            if five > 0 {
                ten++
                five--
            } else {
                return false  
            }
        } else if v == 20 {
            if five > 0 && ten > 0 {
                five--
                ten--
            } else if five >= 3 {
                five -= 3
            } else {
                return false
            }
        }
    }
    return true
}