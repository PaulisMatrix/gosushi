package snippets

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func iterate(head *Node) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}

}

func getNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}

}

func middle(head *Node) *Node {
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *Node) *Node {
	result := &Node{}
	temp := result

	for left != nil && right != nil {
		if left.Val < right.Val {
			temp.Next = left
			left = left.Next
		} else {
			temp.Next = right
			right = right.Next
		}
		temp = temp.Next
	}

	if left != nil {
		temp.Next = left
	}

	if right != nil {
		temp.Next = right
	}

	return result.Next

}

func sortList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	m := middle(head)
	unlink := m
	m = m.Next
	unlink.Next = nil

	left := sortList(head)
	right := sortList(m)

	result := merge(left, right)

	return result

}
func detectAndRemove(head *Node) *Node {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		if fast.Next == slow {
			fmt.Println("Loop starts at", fast.Val)
			return fast
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return nil

}

func newSLL() {
	head := getNode(0)
	n1 := getNode(20)
	n2 := getNode(30)
	n3 := getNode(40)
	n4 := getNode(50)
	n5 := getNode(10)

	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n3

	temp := head
	//detect and remove the cycle and then sort
	result := detectAndRemove(head.Next)
	if result != nil {
		for temp != result {
			temp = temp.Next
		}
		temp.Next = nil
		iterate(sortList(head.Next))
	} else {
		iterate(sortList(head.Next))
	}
}
