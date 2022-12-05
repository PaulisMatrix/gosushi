package snippets

type Stack struct {
	Elements []rune //rune slice
}

func NewStack() *Stack {
	mystack := &Stack{}
	return mystack
}

func (s *Stack) Push(v []rune) {
	s.Elements = append(s.Elements, v...)
}

func (s *Stack) Pop(n int) []rune {
	v := s.Elements[len(s.Elements)-n:]

	//resize
	s.Elements = s.Elements[:len(s.Elements)-n]
	return v
}

func (s *Stack) AddToBottom(r rune) {
	s.Elements = append([]rune{r}, s.Elements...)
}

//convert to string
func (s Stack) String() string {
	var str string
	for _, r := range s.Elements {
		str += string(r) + " "
	}
	return str
}
