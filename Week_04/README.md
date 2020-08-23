# 第四周学习总结

**本章主要学习的内容有：深度优先与广度优先算法、贪心算法以及二分查找**



## 深度优先搜索

+ 概念
深度优先搜索，顾名思义，优先对一条你需要搜索的路径“一探到底”，简称dfs。
+ 步骤
> 1.取root节点
2.根据root节点取节点的左孩子节点：`root.left` ,一直取到左孩子节点的 `node.left == null`
3.再访问上一层的右孩子节点，一直递归上去，直到所有的结点全部访问完
+ 代码模板（go语言）
```
func dfs(res *[type], temp [type], trees [type]) {
	if trees == nil { return res} 
		for trees != nil {
		process(*res, temp, trees)
		dfs(res,temp,trees.left)
		dfs(res,temp,trees.right)
		}
}
```
上面代码模板中，
	* res代表的是存储最终结果的一个变量，
	* [type]代表变量对应的类型
	* temp 代表的是一个“跑龙套”的变量，一般是用来存储res的一个子变量
	* trees 代表要处理的数据结构
注意：通常情况下,res为一个指针，通过引用类型来改变对应的值存储的结果。

## 广度优先搜索
+ 概念
广度优先搜索，顾名思义，按层对整个数据结构进行搜索，简称bfs。
+ 步骤
> 1.取root节点
2.根据root节点取节点的第一级孩子节点：`root.Child(1)` ,一直取到左孩子节点的 `root.Child(n-1) == nil`
3.取到最后一级的孩子节点，访问宣告结束
+ 代码模板（go语言）
```
func bfs(res *[type], temp [type], trees [type], level int) {
	if level > MAX { return res} 
		for level < MAX {
		process(*res, temp, trees, level)
		dfs(res,temp,trees,level+1)
		}
}
```
上面代码模板中，
	* res代表的是存储最终结果的一个变量，
	* [type]代表变量对应的类型
	* temp 代表的是一个“跑龙套”的变量，一般是用来存储res的一个子变量
	* trees 代表要处理的数据结构
	* level 代表访问的层数
注意：通常情况下,res为一个指针，通过引用类型来改变对应的值存储的结果。

### 做题中总结的一些解法思路
+ 在一些二维数组题中的解法，通常会设两个数组
```
    var gx = [4]int{0, 1, 0, -1}
    var gy = [4]int{1, 0, -1, 0}
	for i := 0; i < 4; i++ {
		(x-gx[i],y-gy[i])
	}
```
对于上述操作，其实可以通过一个上北下南左西右东的方向盘来解释，即按顺时针方向依次访问周围的节点，设计较为巧妙
## 贪心算法
+ 概念
贪心算法，即取当下最优解。一般来说，贪心算法不算复杂，只要想用就一定能写出来
+ 贪心算法其实主要是一种思想，就是永远把最优的解法放在最前面执行。
例如 柠檬水找零 这道题，在找零钱时，如果是20元，优先找一个10块一个5块的思想，就是贪心法的思想。
相关代码：
```
 if v == 20 {
 	if five > 0 && ten > 0 {
 	five--
 	ten--
 	} else if five >= 3 {
 	five -= 3
 	}
 }
```
学习笔记
