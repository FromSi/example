package requests

type SortableRequest interface {
	GetData() map[string]string
}
