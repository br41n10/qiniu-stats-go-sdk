package cdn

import (
	"os"
	"testing"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
)

//global variables

var (
	ak     = os.Getenv("accessKey")
	sk     = os.Getenv("secretKey")
	domain = os.Getenv("QINIU_TEST_DOMAIN")

	layout    = "2006-01-02"
	now       = time.Now()
	startDate = now.AddDate(0, 0, -2).Format(layout)
	endDate   = now.AddDate(0, 0, -1).Format(layout)
)

var mac *auth.Credentials
var cdnManager *CdnManager

func init() {
	if ak == "" || sk == "" {
		panic("ak/sk should not be empty")
	}
	mac = auth.New(ak, sk)
	cdnManager = NewCdnManager(mac)
}

// TestGetDynReqCount
func TestGetDomains(t *testing.T) {
	type args struct {
		startDate   string
		endDate     string
		granularity string
		domainList  []string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "DcdnManager_TestGetDynReqCount",
			args: args{
				startDate,
				endDate,
				"day",
				[]string{domain},
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := cdnManager.GetDomains()
			t.Log(ret)
			if err != nil {
				t.Errorf("GetDomains() error = %v", err)
				return
			}
		})
	}
}
