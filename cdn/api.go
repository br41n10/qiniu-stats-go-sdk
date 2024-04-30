package cdn

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/conf"
)

var (
	HostApi = "https://api.qiniu.com"
)

type CdnManager struct {
	mac   *auth.Credentials
	Debug bool
}

func NewCdnManager(mac *qbox.Mac) *CdnManager {
	return &CdnManager{mac, false}
}

// TODO
func (m *CdnManager) GetDomains() ([]string, error) {
	return []string{}, nil
}

// sendGetRequest 发送签名请求，如果 response status code 不是 2xx 则返回错误
func sendGetRequest(mac *qbox.Mac, path string, query string) (resp []byte, err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     HostApi,
		Path:     path,
		RawQuery: query,
	}
	//fmt.Println(u.String())
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest("GET", u.String(), nil); err != nil {
		return
	}

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
