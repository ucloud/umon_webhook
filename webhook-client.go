package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var PostUrl = flag.String("u", "http://localhost/add", "WebHook Url")

// SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
//         Region: "cn-north-03",
//         ResourceType: "uhost",
//         ResourceId: "uhost-xxxx",
//         MetricName: "MemUsage",
//         AlarmTime: 1458733318,
//         RecoveryTime: 0
type WarnMessage struct {
	SessionID    uuid.UUID
	Region       string
	ResourceType string
	ResourceId   string
	MetricName   string
	AlarmTime    int64
	RecoveryTime int64
}

func HttpSend(url string, body []byte) bool {
	hc := &http.Client{
	//		Timeout: time.Duration(5),
	}

	rq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	rq.Header.Set("Context-Type", "application/json")

	rsp, err := hc.Do(rq)
	if err != nil {
		fmt.Println("Post error : ", err.Error())
		return false
	}

	defer rsp.Body.Close()
	b, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(b))

	return true
}

func SendWarnMessage(url string) bool {
	uid, _ := uuid.NewUUID()
	wmsg := &WarnMessage{
		SessionID:    uid,
		Region:       "cn-north-03",
		ResourceType: "uhost",
		ResourceId:   "uhost-xxxx",
		MetricName:   "MemUsage",
		AlarmTime:    time.Now().Unix(),
		RecoveryTime: 0,
	}

	jmsg, err := json.Marshal(wmsg)
	if err != nil {
		//log error
		fmt.Println("Marshal object wmsg failed : ", err.Error())
		return false
	}

	fmt.Println(string(jmsg))

	return HttpSend(url, jmsg)
}

func main() {
	flag.Parse()

	for i := 0; i < 10; i++ {
		time.Sleep(50)
		fmt.Println("Post request for webhook: ", i)
		SendWarnMessage(*PostUrl)
	}
}
