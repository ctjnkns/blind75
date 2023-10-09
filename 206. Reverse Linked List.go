/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverse(cur *ListNode, prev *ListNode) *ListNode {
    if cur == nil {
        return prev
    } else {
        next := cur.Next
        cur.Next = prev
        return reverse(next,cur)
    }
}
func reverseList(head *ListNode) *ListNode {
    //var prev *ListNode
    //cur := head
    //for cur != nil {
    //    nxt := cur.Next
    //    cur.Next = prev
    //    prev = cur
    //    cur = nxt
    //}
    //return prev
    return reverse(head,nil)

}
