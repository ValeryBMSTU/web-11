package usecase

type Provider interface {
	FetchCount() (int, error)
	UpdateCount(int) error
	CheckCountExist() (bool, error)
}
