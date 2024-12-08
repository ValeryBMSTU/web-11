package api

type Usecase interface {
	FetchCount() (int, error)
	IncrementCount(count int) error
}
