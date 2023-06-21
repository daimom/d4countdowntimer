# d4countdowntimer
---

本專案搭配 [d4notify](https://github.com/daimom/d4notify)使用，
自動抓取暗黑編年史的世界王開始時間後，倒數計時

# Usage

設定了API接入點，
呼叫 /line/boss?boss=XXX&unixtime=1687368758 
會開始倒數30分時，發line notify ，
需到 controller/countdownController.go 的 31行，修改 line notify token


