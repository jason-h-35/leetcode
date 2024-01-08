/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// import "github.com/emirpasic/gods/sets/linkedliststack"
import (
    "fmt"
)

func isPalindrome(head *ListNode) bool {
    stack := linkedliststack.New()
    node := head
    size := 0
    // gimme a len so i know when to build and tear down stack
    for node != nil {
        size += 1
        node = node.Next
    }
    node = head
    index := 0
    for node != nil {
        // if less than center push vals to stack
        if index < size / 2 {
            stack.Push(node.Val)
        }
        // if exactly, then do nothing and skip. don't add to stack
        // if greater, then if node val matches stack pop, continue. else fail.
        if index > size / 2 || (index == size / 2 && size % 2 == 0) {
            val, ok := stack.Pop()
            if ok {
                if val != node.Val {
                    fmt.Println("value short circuit\n")
                    return false
                }
            } else {
                return false 
            }
        }
        index += 1
        node = node.Next
    }
    fmt.Println(stack)
    if stack.Empty() {
        return true
    }
    return false
}