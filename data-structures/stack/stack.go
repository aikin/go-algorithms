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

func (stack *Stack) Push(element interface{}) {

	prepend := make([]interface{}, 1)
	prepend[0] = element

	stack.stack = append(prepend, stack.stack...)

	stack.len++
}

func (stack *Stack) Pop() (element interface{}) {
	element, stack.stack = stack.stack[0], stack.stack[1:]
	stack.len--
	return element
}

func (stack *Stack) Peek() interface{} {
	return stack.stack[0]
}