func isValid(s string) bool {
	var st Stack
	for _,b := range s {
		switch b {
		case '(','[','{':
			st.Push(b)
		case ')':
			v, err:= st.Pop()
			if err != nil || v.(rune) != '(' {
				return false
			}
		case ']':
			v, err:= st.Pop()
			if err != nil || v.(rune) != '[' {
				return false
			}
		case '}':
			v, err:= st.Pop()
			if err != nil || v.(rune) != '{' {
				return false
			}
        default:
			return false
		}
	}
    return st.IsEmpty()
}


type Stack []interface {}

func (stack *Stack) Push(value interface{})  {
	*stack = append(*stack, value)
}

func (stack *Stack) Pop() (interface{}, error)  {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack) - 1]
	*stack = theStack[:len(theStack) - 1]
	return value, nil
}

func (stack Stack) IsEmpty() bool  {
	return len(stack) == 0
}