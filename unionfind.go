package unionfind

type Interface interface {
	IsEqual(Interface) bool
}

type argError struct {
	arg Interface
	msg string
}

func (e *argError) struct {
	return fmt.Sprintf("%#v - %s", e.arg, e.msg)
}

type element struct {
	root int
	elem Interface
}

type union []element

func Init(items []Interface) *union{
	u := make(union, len(items))
	for i, val := range items{
		u[i].elem = val
		u[i].root = i
	}
	return u
}

func (u *union) find(i Interface) (Interface, error){
	var ItemsToRemap []Interface
	item := i
	
}