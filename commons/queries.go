package commons

type Queries interface {
	SetQuery(name string, value string)
	DelQuery(name string)
	ToArray() []string
	ToString() string
}
