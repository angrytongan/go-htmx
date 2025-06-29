package weather

import (
	"errors"
	"fmt"
)

var weatherItems = []WeatherPartialItem{
	{City: "City 0", Temperature: "Temperature 0."},
	{City: "City 1", Temperature: "Temperature 1."},
	{City: "City 2", Temperature: "Temperature 2."},
	{City: "City 3", Temperature: "Temperature 3."},
	{City: "City 4", Temperature: "Temperature 4."},
	{City: "City 5", Temperature: "Temperature 5."},
	{City: "City 6", Temperature: "Temperature 6."},
	{City: "City 7", Temperature: "Temperature 7."},
	{City: "City 8", Temperature: "Temperature 8."},
	{City: "City 9", Temperature: "Temperature 9."},
	{City: "City 10", Temperature: "Temperature 10."},
	{City: "City 11", Temperature: "Temperature 11."},
	{City: "City 12", Temperature: "Temperature 12."},
}

var ErrOffsetInvalid = errors.New("offset is invalid")

func Partial(offset, limit int) (WeatherPartial, error) {
	var weatherPartial WeatherPartial

	if offset < 0 || offset > len(weatherItems)/limit {
		return weatherPartial, fmt.Errorf("offset: %d: %w", offset, ErrOffsetInvalid)
	}

	start := offset * limit

	end := start + limit
	if end >= len(weatherItems) {
		end = len(weatherItems)
	}

	for ; start < end; start++ {
		weatherPartial.Items = append(weatherPartial.Items, weatherItems[start])
	}

	if offset > 0 {
		weatherPartial.HasPrev = true
		weatherPartial.PrevPage = offset - 1
	}

	if end < len(weatherItems) {
		weatherPartial.HasNext = true
		weatherPartial.NextPage = offset + 1
	}

	weatherPartial.Limit = limit

	return weatherPartial, nil
}
