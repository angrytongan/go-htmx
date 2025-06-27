package news

type NewsPartialItem struct {
	Headline string
	Content  string
}

type NewsPartial struct {
	Items    []NewsPartialItem
	NextPage int
	PrevPage int
	HasNext  bool
	HasPrev  bool
	Limit    int
}
