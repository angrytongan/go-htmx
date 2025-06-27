package news

import (
	"errors"
	"fmt"
)

var newsItems = []NewsPartialItem{
	{Headline: "Headline 0", Content: "This is content 0."},
	{Headline: "Headline 1", Content: "This is content 1."},
	{Headline: "Headline 2", Content: "This is content 2."},
	{Headline: "Headline 3", Content: "This is content 3."},
	{Headline: "Headline 4", Content: "This is content 4."},
	{Headline: "Headline 5", Content: "This is content 5."},
	{Headline: "Headline 6", Content: "This is content 6."},
	{Headline: "Headline 7", Content: "This is content 7."},
	{Headline: "Headline 8", Content: "This is content 8."},
	{Headline: "Headline 9", Content: "This is content 9."},
	{Headline: "Headline 10", Content: "This is content 10."},
	{Headline: "Headline 11", Content: "This is content 11."},
	{Headline: "Headline 12", Content: "This is content 12."},
}

var ErrOffsetInvalid = errors.New("offset is invalid")

func Partial(offset, limit int) (NewsPartial, error) {
	var newsPartial NewsPartial

	if offset < 0 || offset > len(newsItems)/limit {
		return newsPartial, fmt.Errorf("offset: %d: %w", offset, ErrOffsetInvalid)
	}

	start := offset * limit

	end := start + limit
	if end >= len(newsItems) {
		end = len(newsItems)
	}

	for ; start < end; start++ {
		newsPartial.Items = append(newsPartial.Items, newsItems[start])
	}

	if offset > 0 {
		newsPartial.HasPrev = true
		newsPartial.PrevPage = offset - 1
	}

	if end < len(newsItems) {
		newsPartial.HasNext = true
		newsPartial.NextPage = offset + 1
	}

	newsPartial.Limit = limit

	return newsPartial, nil
}
