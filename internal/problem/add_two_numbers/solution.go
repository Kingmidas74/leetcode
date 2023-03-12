package add_two_numbers

import (
	"github.com/kingmidas74/leetcode/internal/model"
)

func (s *Solution) Solve(arr1 []int, arr2 []int) []int {
	l1 := model.NewListFromArray(arr1)
	l2 := model.NewListFromArray(arr2)
	return s.solve(l1, l2).ToArray()
}

// Solve for https://leetcode.com/problems/add-two-numbers/
func (s *Solution) solve(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	l3 := &model.ListNode{}
	head := l3

	car := 0
	sum := 0
	for {
		one, two := 0, 0
		if l1 != nil {
			one = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			two = l2.Val
			l2 = l2.Next
		}

		sum = one + two + car
		car = sum / 10
		sum = sum % 10

		l3.Val = sum
		if l1 == nil && l2 == nil {
			break
		}

		l3.Next = &model.ListNode{}
		l3 = l3.Next

	}

	if car == 0 {
		return head
	}
	l3.Next = &model.ListNode{}
	l3 = l3.Next
	l3.Val = car
	return head
}
