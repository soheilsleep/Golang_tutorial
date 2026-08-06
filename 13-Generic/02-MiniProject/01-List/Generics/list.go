package Generics

import "github.com/google/go-cmp/cmp"

type List[T any] struct {
	Items []T
}

func main() {

}
func (list *List[T]) InsertAt(index int, item T) {
	// 1,2,3,4,5
	// 1,2,30,3,4,5
	list.Items = append(list.Items, item) // 1,2,3,4,5,30
	copy(list.Items[index+1:], list.Items[index:])
	list.Items[index] = item
}
func (list *List[T]) RemoveAt(index int) {
	list.Items = append(list.Items[:index], list.Items[index+1:]...)
}
func (list *List[T]) Remove(item T) {
	index := list.Find(item)
	if index != -1 {
		list.RemoveAt(index)
	}
}
func (list *List[T]) get(index int) T {
	return list.Items[index]
}
func (list *List[T]) Add(item T) {
	list.Items = append(list.Items, item)
}
func (list *List[T]) Find(item T) int {
	for i, v := range list.Items {
		if cmp.Equal(item, v) {
			return i
		}
	}
	return -1
}
