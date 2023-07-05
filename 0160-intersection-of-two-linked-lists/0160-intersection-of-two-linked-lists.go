/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodeCounts := map[*ListNode]bool{}
    for headA != nil || headB != nil {
        if headA != nil {
            if nodeCounts[headA] {
                return headA
            } else {
                nodeCounts[headA] = true
                headA = headA.Next
            }
        }
         if headB != nil {
            if nodeCounts[headB] {
                return headB
            } else {
                nodeCounts[headB] = true
                headB = headB.Next
            }
        }
    }
    return nil
}

