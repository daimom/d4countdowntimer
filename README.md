# d4countdowntimer

---

本專案建議搭配 [d4notify](https://github.com/daimom/d4notify)使用，
自動抓取暗黑編年史的世界王開始時間，呼叫api後開始倒數計時

## 安裝說明

請先建立 appsetting.json , 
- Token 為倒數計時通知用。
- noneScrapyToken 為第一次通知世界王出現時間用

```json
{
    "Line":{
        "Token":[ 
            "<token>",
            "<token>",
            "<token>"
        ],
        "noneScrapyToken":[ 
            "<token>",
            "<token>"
        ]
    }
}
```

---

## Usage

設定了API接入點，
呼叫 /line/boss?boss=XXX&unixtime=1687368758，

> XXX 為url encode的字串
> unixtime 為unix時間

會開始倒數30分時，發line notify 


