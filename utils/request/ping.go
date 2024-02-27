package request

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

func Ping(endpoint, host string) (t int, err error) {
	t1 := time.Now()
	c := req.C().SetTimeout(10 * time.Second)
	resp, err := c.R().SetHeader(
		"Host", host,
	).Get(endpoint)
	if err != nil || resp.StatusCode != http.StatusOK {
		return 9999, err
	}
	t2 := time.Now()
	t = int(t2.Sub(t1) / time.Millisecond)
	return
}
