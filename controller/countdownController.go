package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

// type d4event struct {
// 	boss     string `json:"name"`
// 	unixtime int    `json:unixtime`
// }

//	type flagger struct {
//		Name      string `json:"name"`
//		Namespace string `json:"namespace"`
//		Phase     string `json:"phase"`
//		Metadata  struct {
//			EventMessage string `json:"eventMessage"`
//			EventType    string `json:"eventType"`
//			Timestamp    string `json:"timestamp"`
//			Channel      string `json:"channel"`
//		}
//	}
func lineNotify(msg string) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	payload := strings.NewReader("message=testabc")

	URL := "https://notify-api.line.me/api/notify"
	req, _ := http.NewRequest("POST", URL, payload)
	req.Header.Add("Authorization", "Bearer <token>")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(respErr)
	}
	fmt.Println(req.Body)

	defer resp.Body.Close()

}
func CheckTimer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("=======Notify=======")
	queryParams := req.URL.Query()
	d4event := make(map[string]string)
	d4event["boss"] = queryParams.Get("boss")
	d4event["unixtime"] = queryParams.Get("unixtime")

	// fmt.Println("boss:" + d4event["boss"] + ",time" + d4event["unixtime"])
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		panic(err)
	}

	unixtime, err := strconv.ParseInt(d4event["unixtime"], 10, 64)
	if err != nil {
		fmt.Println("轉換失敗:", err)
		return
	}

	c := cron.New()
	//* 每分鐘執行一次
	c.AddFunc("* * * * *", func() {

		t := time.Unix(unixtime, 0)
		now := time.Now()
		fmt.Println("now :" + now.Format("2006-01-02 15:04:05"))
		fmt.Println("boss time:" + t.In(location).Format("2006-01-02 15:04:05"))

		firstStopTime := t.Add(-30 * time.Hour)
		// secondStopTime := t.Add(-20*time.Minute)

		//stopTime := time.Date(2023, 6, 21, 16, 48, 0, 0, time.Local)

		if now.After(firstStopTime) {

			subMin := t.Sub(now)
			fmt.Println("還有" + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")

			switch {
			case subMin%10 == 0, subMin%10 == 2, subMin == 3, subMin == 1:
				lineNotify("還有" + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")
			default:
				fmt.Println(subMin)
			}
		}
		if now.After(t) {
			c.Stop()
		}
	})

	c.Start()
}
