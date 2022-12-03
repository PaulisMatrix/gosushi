package snippets

import "fmt"

type Set struct {
	List map[string]struct{} //empty struct takes 0 memory as compared to using a bool and doing s.list[v] = True
}

func NewSet() *Set {
	set := &Set{}
	set.List = make(map[string]struct{})
	return set

}

func (set *Set) Add(v string) {
	set.List[v] = struct{}{}
}

func (set *Set) AddMulti(list []string) {
	for _, v := range list {
		set.Add(v)
	}

}

func (set *Set) Has(v string) bool {
	_, ok := set.List[v]

	return ok
}

func (set *Set) Size() int {
	return len(set.List)
}

func (set *Set) Remove(v string) {
	delete(set.List, v)
}

func (set *Set) Clear() {
	set.List = make(map[string]struct{})
}

func (set *Set) Display() {
	fmt.Print("Your Set: ")
	for k := range set.List {
		fmt.Print(k)
	}
	fmt.Printf("\n")
}

//set1.Intersection(set2)
func (set1 *Set) Intersection(set2 *Set) *Set {
	result := NewSet()
	for v := range set1.List {
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

	for v := range set1.List {
		result.Add(v)
	}

	for v := range set2.List {
		result.Add(v)
	}
	return result
}

//set1.Difference(set2), present in set1 but not in set2
func (set1 *Set) Difference(set2 *Set) *Set {
	result := NewSet()

	for v := range set1.List {
		if set2.Has(v) {
			continue
		}
		result.Add(v)
	}
	return result
}
