package codes

type Codes string

func (code Codes) String() string {
	return string(code)
}

func (code Codes) IsSame(compareWith string) bool {
	return string(code) == compareWith
}

func (code Codes) IsNotSame(compareWith string) bool {
	return !code.IsSame(compareWith)
}
