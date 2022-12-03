package snippets

import "fmt"

type Set struct {
	list map[string]struct{} //empty struct takes 0 memory as compared to using a bool and doing s.list[v] = True
}

func NewSet() *Set {
	set := &Set{}
	set.list = make(map[string]struct{})
	return set

}

func (set *Set) Add(v string) {
	set.list[v] = struct{}{}
}

func (set *Set) AddMulti(list []string) {
	for _, v := range list {
		set.Add(v)
	}

}

func (set *Set) Has(v string) bool {
	_, ok := set.list[v]

	return ok
}

func (set *Set) Size() int {
	return len(set.list)
}

func (set *Set) Remove(v string) {
	delete(set.list, v)
}

func (set *Set) Clear() {
	set.list = make(map[string]struct{})
}

func (set *Set) Display() {
	fmt.Print("Your Set: ")
	for k := range set.list {
		fmt.Print(k)
	}
	fmt.Printf("\n")
}

//set1.Intersection(set2)
func (set1 *Set) Intersection(set2 *Set) *Set {
	result := NewSet()
	for v := range set1.list {
		if !set2.Has(v) {
			continue
		} else {
			result.Add(v)
		}
	}
	return result
}

//set2.Union(set2)
func (set1 *Set) Union(set2 *Set) *Set {
	result := NewSet()

	for v := range set1.list {
		result.Add(v)
	}

	for v := range set2.list {
		result.Add(v)
	}
	return result
}

//set1.Difference(set2), present in set1 but not in set2
func (set1 *Set) Difference(set2 *Set) *Set {
	result := NewSet()

	for v := range set1.list {
		if set2.Has(v) {
			continue
		}
		result.Add(v)
	}
	return result
}
