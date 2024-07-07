package commons

type Query interface {
	SetSortNum(num int)
	GetSortNum() int
	SetValue(value string)
	GetValue() string
	GetName() string
	GetQueryString() string
}
