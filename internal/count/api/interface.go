package api

type Usecase interface {
	FetchCount() (int, error)
	IncreaseCount(cnt int) error
}
