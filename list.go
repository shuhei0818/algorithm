package algorithm

type List[T any] interface {
	Get() (T, error)
	Add(d T) error
}
