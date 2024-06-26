package stats

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
	"github.com/qiniu/go-sdk/v7/storage"
)

type StatsManager struct {
	mac *auth.Credentials
}

func NewStatManager(mac *qbox.Mac) *StatsManager {
	return &StatsManager{mac: mac}
}

// sendGetRequest 发送签名请求，如果 response status code 不是 2xx 则返回错误
func sendGetRequest(mac *qbox.Mac, path string, query string) (resp []byte, err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     storage.DefaultAPIHost,
		Path:     path,
		RawQuery: query,
	}
	//fmt.Println(u.String())
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}
	request.Header.Set("Host", storage.DefaultAPIHost)
	request.Header.Set("Content-Type", conf.CONTENT_TYPE_FORM)
	if _, err = mac.SignRequest(request); err != nil {
		return
	}

	if err = mac.AddToken(auth.TokenQBox, request); err != nil {
		return
	}

	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode/100 != 2 {
		err = fmt.Errorf("response status code %d", response.StatusCode)
		return
	}

	if resp, err = io.ReadAll(response.Body); err != nil {
		return
	}

	return
}
