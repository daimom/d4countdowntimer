package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	_ "github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

type Job struct {
	T         time.Time `json:"unix"`
	Boss_time string    `json:"boss_time"`
	Boss      string    `json:"boss"`
}

func lineNotify(msg string) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	payload := strings.NewReader("message=" + msg)

	URL := "https://notify-api.line.me/api/notify"
	req, _ := http.NewRequest("POST", URL, payload)

	req.Header.Add("Authorization", "Bearer <token>")

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(respErr)
	}
	// fmt.Println(req.Body)

	defer resp.Body.Close()

}

func CheckTimer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("=======Notify=======")
	fmt.Println(req.URL.RawQuery)

	queryParams := req.URL.Query()
	// queryParams := req.URL.Query()
	d4event := make(map[string]string)
	d4event["boss"], _ = url.QueryUnescape(queryParams.Get("boss"))
	d4event["unixtime"] = queryParams.Get("unixtime")

	// fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "boss:" + d4event["boss"] + ",time" + d4event["unixtime"])
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		panic(err)
	}

	unixtime, err := strconv.ParseInt(d4event["unixtime"], 10, 64)
	if err != nil {
		fmt.Println("轉換失敗:", err)
		return
	}
	t := time.Unix(unixtime, 0)
	boss_time := t.In(location).Format("2006-01-02 15:04:05")
	fmt.Println("(" + time.Now().Format("2006-01-02 15:04:05") + ")boss:" + d4event["boss"] + ",time:" + boss_time)

	lineJob := Job{
		T:         t,
		Boss_time: boss_time,
		Boss:      d4event["boss"],
	}
	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
	c.AddJob("* * * * *", &lineJob)
	// now := time.Now()
	// if now.After(t) {
	// 	dec.Stop()
	// }
	// 啟動執行任務
	c.Start()

	// 退出時關閉工作排程
	// defer c.Stop()

	// c := cron.New()

	// //* 每分鐘執行一次
	// c.AddFunc("* * * * *", func() {

	// 	now := time.Now()
	// 	// fmt.Println("now :" + now.Format("2006-01-02 15:04:05"))
	// 	// fmt.Println("boss time:" + boss_time)

	// 	firstStopTime := t.Add(-30 * time.Hour)

	// 	//stopTime := time.Date(2023, 6, 21, 16, 48, 0, 0, time.Local)
	// 	msg := "距離世界王" + d4event["boss"] + " 出現時間，" + boss_time + "，還有"
	// 	if now.After(firstStopTime) {

	// 		subMin := t.Sub(now)
	// 		lineNotify(msg + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")

	// 		countdownMin := int(subMin.Minutes())
	// 		switch {
	// 		case countdownMin%10 == 0, countdownMin == 3, countdownMin == 1:
	// 			//fmt.Println("in")
	// 			lineNotify(msg + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")
	// 		default:
	// 			//fmt.Println(subMin)
	// 		}
	// 	}
	// 	if now.After(t) {
	// 		c.Stop()
	// 	}
	// })

	// c.Start()
}

func (j *Job) Run() {
	now := time.Now()
	fmt.Println("now :" + now.Format("2006-01-02 15:04:05"))
	fmt.Println("boss time:" + j.Boss_time)
	firstStopTime := j.T.Add(-30 * time.Hour)

	//stopTime := time.Date(2023, 6, 21, 16, 48, 0, 0, time.Local)
	msg := "距離世界王" + j.Boss + " 出現時間，" + j.Boss_time + "，還有"
	if now.After(firstStopTime) {

		subMin := j.T.Sub(now)
		lineNotify(msg + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")

		countdownMin := int(subMin.Minutes())
		switch {
		case countdownMin%10 == 0, countdownMin == 3, countdownMin == 1:
			lineNotify(msg + strconv.FormatFloat(subMin.Minutes(), 'f', 0, 64) + "分鐘")
		default:
			//fmt.Println(subMin)
		}
	}
	if now.After(j.T) {
		return
	}
}
