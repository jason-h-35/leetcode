/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    asSlice := make([]*ListNode, 0, 30)
    node := head
    for node != nil {
        asSlice = append(asSlice, node)
        node = node.Next
    }
    if len(asSlice) == 1 {
        return nil
    } else if n == 1 {
        nth := asSlice[len(asSlice)-2]
        nth.Next = nil
    } else {
        nth := asSlice[len(asSlice)-n]
        *nth = *nth.Next
    }
    return head
}