package weather

type WeatherPartialItem struct {
	City        string
	Temperature string
}

type WeatherPartial struct {
	Items    []WeatherPartialItem
	NextPage int
	PrevPage int
	HasNext  bool
	HasPrev  bool
	Limit    int
}
