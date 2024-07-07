package commons

import (
	"sort"
	"strings"
)

func NewMapQueries() MapQueries {
	return make(MapQueries)
}

type MapQueries map[string]Query

func (m MapQueries) Has(name string) bool {
	var exist bool
	_, exist = m[name]
	return exist
}

func (m MapQueries) SetQuery(name string, value string) {

	if _, exist := m[name]; exist {
		m[name].SetValue(value)
		return
	}

	m[name] = newStringQuery(m.Length(), name, value)

}

func (m MapQueries) DelQuery(name string) {

	if _, exist := m[name]; !exist {
		return
	}

	delete(m, name)

	m.reSort()

}

func (m MapQueries) GetQueryValue(name string) string {

	var (
		query     Query
		nameExist bool
	)

	if query, nameExist = m[name]; !nameExist {
		return ""
	}

	return query.GetValue()
}

func (m MapQueries) GetQueriesAsArray() []Query {

	var queries []Query

	for _, query := range m {
		queries = append(queries, query)
	}

	sort.Slice(queries, func(i int, j int) bool {
		return queries[i].GetSortNum() < queries[j].GetSortNum()
	})

	return queries
}

func (m MapQueries) ToArray() []string {

	var (
		queries []Query = m.GetQueriesAsArray()
		arr     []string
	)

	for _, query := range queries {
		arr = append(arr, query.GetQueryString())
	}

	return arr
}

func (m MapQueries) ToString() string {

	var queries []string = m.ToArray()

	return strings.Join(queries, "&")
}

func (m MapQueries) ToStringWithTail(tailQueries ...string) string {

	var queries []string = m.ToArray()
	queries = append(queries, tailQueries...)
	return strings.Join(queries, "&")

}

func (m MapQueries) Length() int {
	return len(m)
}

func (m MapQueries) reSort() {

	var (
		queries []Query = m.GetQueriesAsArray()
	)

	for i, query := range queries {
		m[query.GetName()].SetSortNum(i)
	}

}
