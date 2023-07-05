/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    asSlice := make([]*ListNode, 0, 30)
    head2 := head
    for head != nil {
        asSlice = append(asSlice, head)
        head = head.Next
    }
    if len(asSlice) == 1 {
        return nil
    } else if n == 1 {
        node := asSlice[len(asSlice)-2]
        node.Next = nil
    } else {
        node := asSlice[len(asSlice)-n]
        *node = *node.Next
    }
    return head2
}