package commons

type Set struct {
	data map[interface{}]bool
}

func NewSet() *Set {
	return &Set{
		data: map[interface{}]bool{},
	}
}

func (set Set) PutAllInt64s(iArr *[]int64) {
	for _, key := range *iArr {
		set.Put(key)
	}
}

func (set Set) PutAll(iArr []interface{}) {
	for _, key := range iArr {
		set.Put(key)
	}
}

func (set Set) Put(key interface{}) {
	set.data[key] = true
}

func (set Set) Contains(key interface{}) bool {
	v := set.data[key]
	return v == true
}

func (set Set) All() []interface{} {
	var keys []interface{}
	for key, _ := range set.data {
		keys = append(keys, key)
	}
	return keys
}
