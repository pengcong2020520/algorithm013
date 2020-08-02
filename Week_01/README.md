学习笔记
Go中Queue的实现方式:
        队列的特性较为单一，基本操作即初始化、获取大小、添加元素、移除元素等。最重要的特性就是满足先进先出。
        /*
	在go中源码没有queue的数据结构，故在实现queue时，通常使用以下两种方式进行实现
	*/
    第一种方式：切片  (针对array结构下的queue)
        /* 
         先定义queue为一个切片
         */
        var queue []interface{}

        /* 
        实现queue的peek操作 
        Peek操作只需要获取队列队头的元素即可，不用删除。
        */
        func Peek(queue []interface{}) interface{} {
            if len(queue) == 0 {
                fmt.Println("queue is empty")
                return -1
            }
            return queue[1:]
        }

        /* 
        实现queue的Add操作 
        Add操作添加元素，通过append添加以后，新元素刚好在queue的尾部
        */
        func Add(queue []interface{}, value interface{}) []interface{} {
            return append(queue, value)
        }

        /* 
        实现queue的Remove操作 
        Remove操作删除元素，通过append将要删除元素前后的切片重新组合起来
        */
        func Remove(queue []interface{}, value interface{}) []interface {
            for i, v := range queue {
                if v == value {
                    return append(queue[:i-1], queue[i+1:]...)
                }
            }
        }

    第二种方式：结构体+指针  (针对linklist结构下的queue)
        /* 
            先定义LinkedQueue的结构体
         */
        type LinkedQueue struct {
            head *node
            tail *node
            size int    
        }

        type node struct {
            value interface{}
            prev *node
            next *node
        }
        /* 
            实现queue的peek操作 
            Peek操作只需要获取队列队头的元素即可，不用删除。
         */
        func (queue *LinkedQueue) Peek() interface{} {
            if queue.head == nil {
                panic("Empty queue.")
            }
            return queue.head.value
        }
        /* 
            实现queue的Add操作 
            Add操作添加元素，通过append添加以后，新元素刚好在queue的尾部
         */
        func (queue *LinkedQueue) Add(value interface{}) {
            new_node := &node{value, queue.tail, nil}
            if queue.tail == nil {
                queue.head = new_node
                queue.tail = new_node
            } else {
                queue.tail.next = new_node
                queue.tail = new_node
            }
            queue.size++
            new_node = nil
        } 

        /* 
            实现queue的Remove操作 
            Remove操作删除元素，通过append将要删除元素前后的切片重新组合起来
        */    
        func (queue *LinkedQueue) Remove() {
            if queue.head == nil {
                panic("Empty queue.")
            }
            first_node := queue.head
            queue.head = first_node.next
            first_node.next = nil
            first_node.value = nil
            queue.size--
            first_node = nil
        }





    

