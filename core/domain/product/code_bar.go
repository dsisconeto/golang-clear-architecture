package product

type CodeBar struct {
	Code string
}

func NewCodeBar(code string) *CodeBar {
	return &CodeBar{Code: code}
}
