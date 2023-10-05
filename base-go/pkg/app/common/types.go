package common

type InfiniteScrollResult[T any] struct {
	nextCursor string
	data       []T
	total      int
}

func NewInfiniteScrollResult[T any](nextCursor string, data []T, total int) InfiniteScrollResult[T] {
	return InfiniteScrollResult[T]{nextCursor: nextCursor, data: data, total: total}
}

func (r InfiniteScrollResult[T]) NextCursor() string {
	return r.nextCursor
}

func (r InfiniteScrollResult[T]) Data() []T {
	return r.data
}

func (r InfiniteScrollResult[T]) Total() int {
	return r.total
}

type TableAdvancedResult[T any] struct {
	startAfter string
	endBefore  string
	data       []T
	total      int
	itemCount  int
}

func NewTableAdvancedResult[T any](startAfter string, endBefore string, data []T, total int, itemCount int) TableAdvancedResult[T] {
	return TableAdvancedResult[T]{startAfter: startAfter, endBefore: endBefore, data: data, total: total, itemCount: itemCount}
}

func (r TableAdvancedResult[T]) StartAfter() string {
	return r.startAfter
}

func (r TableAdvancedResult[T]) EndBefore() string {
	return r.endBefore
}

func (r TableAdvancedResult[T]) Data() []T {
	return r.data
}

func (r TableAdvancedResult[T]) Total() int {
	return r.total
}

func (r TableAdvancedResult[T]) ItemCount() int {
	return r.itemCount
}
