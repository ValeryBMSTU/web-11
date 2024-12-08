package usecase

type Provider interface {
	FetchQuery(name string) (string, error)
	InsertQuery(name string) error
	CheckQueryExist(name string) (bool, error)
}
