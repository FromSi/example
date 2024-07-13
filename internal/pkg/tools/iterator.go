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
	return len(iterator.order) > iterator.index && len(iterator.order) > iterator.index
}

func (iterator *MapStringIterator) GetNext() (string, string) {
	key := ""
	value := ""

	if iterator.HasNext() {
		if item, exists := iterator.data[iterator.order[iterator.index]]; exists {
			key = iterator.order[iterator.index]
			value = item

			iterator.index++
		}
	}

	return key, value
}
