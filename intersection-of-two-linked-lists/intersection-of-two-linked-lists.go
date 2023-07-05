/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    nodeCounts := map[*ListNode]int{}
    for headA != nil || headB != nil {
        if headA != nil {
            nodeCounts[headA]++
            headA = headA.Next
        }
       if headB != nil {
            nodeCounts[headB]++
            headB = headB.Next
        }
        for k, v := range nodeCounts {
            if v == 2 {
                return k
            }
        }
    }
    return nil
}

