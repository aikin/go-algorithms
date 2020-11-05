package stack

type Stack struct {

	stack []interface{}

	len int
}

func NewStack() *Stack {
	stack := &Stack{}
	stack.stack = make([]interface{}, 0)
	stack.len = 0
	return stack
}

func (stack *Stack) Len() int {
	return stack.len
}

func (stack *Stack) IsEmpty() bool {
	return stack.len == 0
}