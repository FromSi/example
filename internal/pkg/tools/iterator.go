package tools

type MapStringIterator struct {
	index int
	order []string
	data  map[string]string
}

func NewMapStringIterator(data map[string]string) *MapStringIterator {
	var order []string

	for key := range data {
		order = append(order, key)
	}

	return &MapStringIterator{
		index: 0,
		order: order,
		data:  data,
	}
}

func (iterator *MapStringIterator) HasNext() bool {
	isNextIndex := len(iterator.order) > iterator.index

	if !isNextIndex {
		return false
	}

	_, isNextData := iterator.data[iterator.order[iterator.index]]

	return isNextData
}

func (iterator *MapStringIterator) GetNext() (string, string) {
	key := ""
	value := ""

	if iterator.HasNext() {
		key = iterator.order[iterator.index]
		item := iterator.data[key]
		value = item

		iterator.index++
	}

	return key, value
}
