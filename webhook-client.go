package main

import (
	"fmt"
	"net/http"
)

// SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
//         Region: "cn-north-03",
//         ResourceType: "uhost",
//         ResourceId: "uhost-xxxx",
//         MetricName: "MemUsage",
//         AlarmTime: 1458733318,
//         RecoveryTime: 0
type WarnMessage struct {
	SessionID    string
	Region       string
	ResourceType string
	ResourceId   string
	MetricName   string
	AlarmTime    int64
	RecoveryTime int64
}

func main() {
	http.Post(url, bodyType, body)
}
