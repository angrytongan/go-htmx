package weather

type Repository interface {
	Partial(offset, limit int) (WeatherPartial, error)
}
