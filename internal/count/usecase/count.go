package usecase

func (u *Usecase) FetchCount() (int, error) {
	msg, err := u.p.FetchCount()
	if err != nil {
		return 0, err
	}

	if msg == 0 {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) IncrementCount(count int) error {
	isExist, err := u.p.CheckCountExist()
	if err != nil {
		return err
	}

	if !isExist {
		return nil
	}

	err = u.p.UpdateCount(count)
	if err != nil {
		return err
	}

	return nil
}
