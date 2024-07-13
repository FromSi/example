package requests

type PageableRequest interface {
	GetPage() int
	GetLimit() int
}
