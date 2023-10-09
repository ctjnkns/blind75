/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    slow := head
    fast := head.Next

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    second := slow.Next
    slow.Next = nil
    prev := slow.Next

    for second != nil {
        nxt := second.Next
        second.Next = prev
        prev = second
        second = nxt
    }

    first := head
    second = prev
    for second != nil {
        nxt1 := first.Next
        nxt2 := second.Next
        first.Next = second
        second.Next = nxt1
        first = nxt1
        second = nxt2
    }

}
