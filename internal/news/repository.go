package news

type Repository interface {
	Partial(offset, limit int) (NewsPartial, error)
}
