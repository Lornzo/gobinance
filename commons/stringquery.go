package commons

import "fmt"

func newStringQuery(sortNum int, name string, value string) *StringQuery {
	return &StringQuery{
		SortNum: sortNum,
		Name:    name,
		Value:   value,
	}
}

type StringQuery struct {
	SortNum int
	Name    string
	Value   string
}

func (s *StringQuery) GetName() string {
	return s.Name
}

func (s *StringQuery) SetSortNum(num int) {
	s.SortNum = num
}

func (s *StringQuery) GetSortNum() int {
	return s.SortNum
}

func (s *StringQuery) SetValue(value string) {
	s.Value = value
}

func (s *StringQuery) GetValue() string {
	return s.Value
}

func (s *StringQuery) GetQueryString() string {
	return fmt.Sprint(s.Name, "=", s.Value)
}
