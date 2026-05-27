package datastructures

type Stack[K any] struct {
	values []*K
}

func (stack *Stack[K]) Push(value *K) { _ = "STUB: not implemented"; return }

func (stack *Stack[K]) Pop() *K { _ = "STUB: not implemented"; return nil }

func (stack *Stack[K]) Top() *K { _ = "STUB: not implemented"; return nil }

func (stack *Stack[K]) Size() int { _ = "STUB: not implemented"; return 0 }

func (stack *Stack[K]) Clear() { _ = "STUB: not implemented"; return }

func (stack *Stack[K]) String() string { _ = "STUB: not implemented"; return "" }
