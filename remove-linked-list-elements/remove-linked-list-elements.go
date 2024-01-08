/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    begin := &ListNode{0,head}
    prev := begin
    node := head
    for node != nil {
        if node.Val == val {
            prev.Next = node.Next
            node = node.Next
        } else {
            prev = node
            node = node.Next
        }
    }
    return begin.Next
}