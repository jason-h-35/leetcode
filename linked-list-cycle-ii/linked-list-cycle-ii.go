/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    node := getCycleNode(head)
    index := 0
    for head != nil {
        if head == node {
            return head
        } else {
            index++
            head = head.Next
        }
    }
    return nil
}

func getCycleNode(head *ListNode) *ListNode {
    seen := map[*ListNode]bool{}
    for head != nil {
        if seen[head] {
            return head
        } else {
            seen[head] = true
        }
        head = head.Next
    }
    return nil
}