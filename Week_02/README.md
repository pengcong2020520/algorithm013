学习笔记
### golang Set的实现方式
	在golang中，hash的方式是通过map的方式（key-value键值对这样的形式）实现：
/*var hash map[key类型]value类型*/

	且对于map类型来说，golang中的map自带一个删除重复性的特点，也就是key的值是唯一的，value值可以由多个key值指向，例如查找map中的某个元素：
/*value, ok := map[key]
	//其中ok代表的就是map中是否包含该元素。*/
	对于hash的实现方式，其实在底层主要是通过哈希算法来实现，也就是对于每一个key通过哈希算法以后都会得到一个哈希值，最终将不同的键映射到不同的索引上，但是哈希算法中输入的键的数量级往往是大于输出的值的数量级的，故在大量的例子中，必然会出现不同的key映射到同一个位置的方式，golang在这个地方沿用了C语言的溢出桶设计方式来进行实现（毕竟同一个祖先嘛），这样可以减少扩容的频次。
	而对于其它语言的set，在golang中的源码里没有这种定义，但是看大部分网友的实现方式，均是采用map来实现，但是由于map的value是会消耗内存的，所以对于set的定义大多数是使用以下方式
/*var set map[interface{}]struct{} value用一个空的结构体来表示*/

对于golang中的空结构体是不需要消耗内存的，因而很好的实现了其他语言的set。

