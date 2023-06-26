package utils

type KeyType interface {
	int | string
}

type ValueType interface {
	int | string | []byte
}

func GetKey[T KeyType, V ValueType](origin map[T]V) []T {
	ans := make([]T, 0)
	for k := range origin {
		ans = append(ans, k)
	}
	return ans
}
