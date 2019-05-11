package unionfn

import "fmt"

type Interface interface {
	IsEqual(Interface) bool
}
type argError struct {
	arg Interface
	msg string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%#v - %s", e.arg, e.msg)
}

type union struct {
	items map[Interface]uint32
	roots []Interface
	ranks []uint8
}

func Init(items []Interface) *union {
	u := new(union)
	u.items = make(map[Interface]uint32, len(items))
	u.roots = make([]Interface, len(items))
	u.ranks = make([]uint8, len(items))
	for i, val := range items {
		u.items[val] = uint32(i)
		u.roots[i] = val
		u.ranks[i] = 0
	}
	return u
}

func (u *union) Connected(a, b Interface) (bool, error) {
	aRoot, err := u.find(a)
	if err != nil {
		return false, err
	}
	bRoot, err := u.find(b)
	if err != nil {
		return false, err
	}
	return aRoot.IsEqual(bRoot), nil
}
func (u *union) FindRoot(a Interface) (Interface, error) {
	return u.find(a)
}

func (u *union) Merge(a, b Interface) error {

	aRoot, err := u.find(a)
	if err != nil {
		return err
	}
	bRoot, err := u.find(b)
	if err != nil {
		return err
	}
	if aRoot.IsEqual(bRoot) {
		return &argError{b, " already has the same root"}
	}
	aRootPos := u.items[aRoot]
	bRootPos := u.items[bRoot]
	if u.ranks[aRootPos] > u.ranks[bRootPos] {
		u.roots[bRootPos] = aRoot
	} else {
		if u.ranks[aRootPos] == u.ranks[bRootPos] {
			u.ranks[bRootPos]++
		}
		u.roots[aRootPos] = bRoot
	}
	return nil
}

func (u *union) compresPaths(items []Interface, root Interface) {
	for _, val := range items {
		u.roots[u.items[val]] = root
	}
}

func (u *union) find(i Interface) (Interface, error) {
	var itemsToRemap []Interface
	item := i
	rootPos, has := u.items[item]
	if has == false {
		return nil, &argError{i, " is not in universe passed to union-find"}
	}
	for !u.roots[rootPos].IsEqual(item) {
		itemsToRemap = append(itemsToRemap, item)
		item = u.roots[rootPos]
		rootPos = u.items[item]
	}
	u.compresPaths(itemsToRemap, item)
	return item, nil
}
