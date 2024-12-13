package usecase

func (u *Usecase) FetchCount() (int, error) {
	msg, err := u.p.GetCount()
	if err != nil {
		return 0, err
	}

	if msg == 0 {
		return u.defaultCnt, nil
	}

	return msg, nil
}

func (u *Usecase) IncreaseCount(cnt int) error {
	err := u.p.AddCount(cnt)
	if err != nil {
		return err
	}

	return nil
}
