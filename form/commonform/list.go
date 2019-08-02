package commonform

//StringList string list
type StringList map[string]bool

//Has check if given value  in list
func (l *StringList) Has(value string) bool {
	return (*l)[value]
}

//NewStringList create new string list
func NewStringList(values ...string) *StringList {
	m := StringList(make(map[string]bool, len(values)))
	for k := range values {
		m[values[k]] = true
	}
	return &m
}

//IntList string list
type IntList map[int]bool

//Has check if given value  in list
func (l *IntList) Has(value int) bool {
	return (*l)[value]
}

//NewIntList create new string list
func NewIntList(values ...int) *IntList {
	m := IntList(make(map[int]bool, len(values)))
	for k := range values {
		m[values[k]] = true
	}
	return &m
}

//Int64List string list
type Int64List map[int64]bool

//Has check if given value  in list
func (l *Int64List) Has(value int64) bool {
	return (*l)[value]
}

//NewInt64List create new string list
func NewInt64List(values ...int64) *Int64List {
	m := Int64List(make(map[int64]bool, len(values)))
	for k := range values {
		m[values[k]] = true
	}
	return &m
}
