package model

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListFromArray(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}

	l := &ListNode{}
	result := l
	for i, v := range input {
		l.Val = v
		if i < len(input)-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}

	return result
}

func (l *ListNode) ToArray() []int {
	result := make([]int, 0)

	s := l
	for s != nil {
		result = append(result, s.Val)
		s = s.Next
	}

	return result
}
