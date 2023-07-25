package util

func EmptySetOf[T comparable]() map[T]bool {
	return map[T]bool{}
}

func SetOf[T comparable](items ...T) map[T]bool {
	s := map[T]bool{}
	for _, item := range items {
		s[item] = true
	}
	return s
}
