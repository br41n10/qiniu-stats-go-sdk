package stats

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/qiniu/go-sdk/v7/auth"
)

//global variables

var (
	ak     = os.Getenv("QINIU_TEST_ACCESS_KEY")
	sk     = os.Getenv("QINIU_TEST_SECRET_KEY")
	bucket = os.Getenv("QINIU_TEST_BUCKET")
	region = os.Getenv("QINIU_TEST_REGION")

	layout    = "2006-01-02"
	now       = time.Now()
	startDate = now.AddDate(0, 0, -1).Format(layout)
	endDate   = now.AddDate(0, 0, -1).Format(layout)
)

var mac *auth.Credentials
var kodoStatsManager *StatsManager

func init() {
	if ak == "" || sk == "" {
		panic("ak/sk should not be empty")
	}
	mac = auth.New(ak, sk)

	kodoStatsManager = NewStatManager(mac)
}

// TestGetGetFileCountData
func TestGetGetFileCountData(t *testing.T) {
	fmt.Println(123)
	type args struct {
		startDate   string
		endDate     string
		granularity string
		bucket      string
		region      string
	}

	testCases := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "KodoStatsManager_TestGetGetFileCountData",
			args: args{
				startDate,
				endDate,
				"day",
				bucket,
				region,
			},
			wantCode: 200,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ret, err := kodoStatsManager.GetFileCount(0, tc.args.startDate, tc.args.endDate, tc.args.granularity, tc.args.bucket, "")
			fmt.Printf("%+v\n", ret)
			if err != nil {
				t.Errorf("GetFileCount() error = %v", err)
				return
			}
		})
	}
}
