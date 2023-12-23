package model

import (
	"fmt"
	"net/url"
	"sort"

	helper "url-shortener/helper"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	baseURL = "http://localhost:3000/short/"
)

type Metrics struct {
	Host  string `json:"host"`
	Count int    `json:"count"`
}

type UrlShortener struct {
	data    map[string]string
	metrics map[string]int
}

func NewURLShortener() *UrlShortener {
	return &UrlShortener{
		data:    make(map[string]string),
		metrics: make(map[string]int),
	}
}

func (u *UrlShortener) ShortURL(request string) (string, error) {
	shortString, err := helper.GenerateRandomStringWithKey(request, 4)
	if err != nil {
		return "", err
	}

	if u.data[shortString] != "" {
		return fmt.Sprintf(baseURL + shortString), nil
	}

	newURL := fmt.Sprintf(baseURL + shortString)
	u.data[shortString] = request
	err = u.addMetrics(request)
	if err != nil {
		return "", err
	}

	return newURL, nil
}

func (u *UrlShortener) GetMetrics() []Metrics {
	var response []Metrics
	for k, v := range u.metrics {
		response = append(response, Metrics{k, v})
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].Count > response[j].Count
	})

	if len(response) > 3 {
		return response[:3]
	}

	return response
}

func (u *UrlShortener) addMetrics(request string) error {
	parsed, err := url.Parse(request)
	if err != nil {
		return err
	}

	count := u.metrics[parsed.Host]
	u.metrics[parsed.Host] = count + 1

	return nil
}

func (u *UrlShortener) GetMainURL(shortURL string) string {
	return u.data[shortURL]
}
