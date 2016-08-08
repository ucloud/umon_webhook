package restfulAPI

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

// SessionID: "xxxxxxxxxxxxxxxxxxxxxxx",
//         Region: "cn-north-03",
//         ResourceType: "uhost",
//         ResourceId: "uhost-xxxx",
//         MetricName: "MemUsage",
//         AlarmTime: 1458733318,
//         RecoveryTime: 0
type WarnMessage struct {
	SessionID    uuid.UUID `json:",string"`
	Region       string
	ResourceType string
	ResourceId   string
	MetricName   string
	AlarmTime    int64
	RecoveryTime int64
}

type WarnMsgs []WarnMessage

type NormalResopnse struct {
	SessionID uuid.UUID `json:",omitempty,string"`
	RetCode   int
	Error     string `json:",omitempty"`
}

const (
	Insert_STMT        = "INSERT INTO warn_message VALUES (?, ?, ?, ?, ?, ?, ?)"
	Select_All_STMT    = "SELECT * FROM warn_message"
	Select_Filter_STMT = "SELECT * FROM warn_message WHERE SessionID=?"
)

//Insert
func SaveWarnMessage(ws *WarnMessage) error {
	db := GetDB()
	if db == nil {
		log.Fatal("Database can not access!")
	}

	Ins, err := db.Prepare(Insert_STMT)
	if err != nil {
		log.Println("Can not save warn message : ", err.Error())
		return err
	}

	defer Ins.Close()

	_, err = Ins.Exec(ws.SessionID, ws.Region, ws.ResourceType, ws.ResourceId, ws.MetricName, ws.AlarmTime, ws.RecoveryTime)
	if err != nil {
		log.Println("Save warn message error : ", err.Error())
		return err
	}

	return nil
}

//Get all warn message
func GetAllWarnMessage() (WarnMsgs, error) {
	db := GetDB()
	if db == nil {
		log.Fatal("Database can not access!")
	}

	rows, err := db.Query(Select_All_STMT)
	if err != nil {
		log.Println("Can not get warn message : ", err.Error())
		return nil, err
	}

	defer rows.Close()

	var msg WarnMsgs
	for rows.Next() {
		var wm WarnMessage
		err = rows.Scan(&wm.SessionID,
			&wm.Region,
			&wm.ResourceType,
			&wm.ResourceId,
			&wm.MetricName,
			&wm.AlarmTime,
			&wm.RecoveryTime)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("Scan database error : ", err.Error())
				return nil, err
			} else {
				msg = append(msg, wm)
				break
			}
		}

		msg = append(msg, wm)
	}

	return msg, nil
}
