package nhj

import (
	"net/http"
)

func MakeClient(connsCount int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost:     connsCount,
			MaxIdleConnsPerHost: connsCount,
		},
	}
}
