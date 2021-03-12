package other

type schmidt struct {
	sex      string
	fullName string
	strong   bool
}

func (sch schmidt) makeSituationAwkward() string {
	return "Can someone please get my towel? Its in my room next to the Irish walking cape."
}

func (sch schmidt) partOfSociety() bool {
	return true
}

func (sch schmidt) lieButBelieveItTrue() bool {
	return false
}
