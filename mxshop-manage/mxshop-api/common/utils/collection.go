package utils

// Set Set[T comparable] 是一个泛型集合，适用于任何可比较类型（string, int 等）
type Set[T comparable] struct {
	set map[T]struct{}
}

// NewSet 创建一个 Set 实例，并将 slice 转换为 map
func NewSet[T comparable](slice []T) *Set[T] {
	set := make(map[T]struct{})
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return &Set[T]{set: set}
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.set[value]
	return exists
}
