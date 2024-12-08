package usecase

func (u *Usecase) FetchQuery(name string) (string, error) {
	msg, err := u.p.FetchQuery(name)
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) InsertQuery(name string) error {
	isExist, err := u.p.CheckQueryExist(name)
	if err != nil {
		return err
	}

	if isExist {
		return nil
	}

	err = u.p.InsertQuery(name)
	if err != nil {
		return err
	}

	return nil
}
