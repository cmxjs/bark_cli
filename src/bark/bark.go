package bark

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Bark struct {
	Host              string
	Key               string
	IsArchive         int // 1 indicate auto archive
	AutomaticallyCopy int // 1 indicate automatically copy
}

func (b Bark) Send(body string, title string, group string) (statusCode int, err error) {
	p := url.Values{}
	p.Add("isArchive", strconv.Itoa(b.IsArchive))
	p.Add("automaticallyCopy", strconv.Itoa((b.AutomaticallyCopy)))
	if group != "" {
		p.Add("group", url.QueryEscape(group))
	}

	url := (strings.TrimRight(b.Host, "/") + "/" +
		strings.TrimRight(b.Key, "/") + "/" +
		url.QueryEscape(title) + "/" +
		url.QueryEscape(body) + "?" +
		p.Encode())
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}
