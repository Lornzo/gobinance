package commons

type Queries interface {
	SetQuery(name string, value string)
	DelQuery(name string)
	GetQueryValue(name string) string
	ToArray() []string
	ToString() string
}
