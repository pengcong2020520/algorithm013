/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    if l1 == nil {
        return l2
    } else if l2 == nil {
        return l1
    }
    if l1.Val > l2.Val {
        dummy.Val = l2.Val
        dummy.Next = mergeTwoLists(l1, l2.Next)
    } else {
        dummy.Val = l1.Val
        dummy.Next = mergeTwoLists(l1.Next, l2)
    }
    return dummy
}