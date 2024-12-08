package api

type Usecase interface {
	FetchQuery(name string) (string, error)
	InsertQuery(name string) error
}
