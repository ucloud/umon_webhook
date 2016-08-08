package restfulAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

//const
const (
	HContentType = "Content-Type"
	ContentHTML  = "text/html;charset=UTF-8"
	ContentJSON  = "application/json;charset=UTF-8"
	ContentPLAIN = "application/plain;charset=UTF-8"
	ContentMD    = "application/markdown;charset=UTF-8"
	DOCName      = "doc/user_guide.html"
)

const (
	ErrorFormat = "Response for request: %s %s is failed, error : %s\n"
)

func SendJsonResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	nrjson, err := json.Marshal(v)
	if err != nil {
		SendErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set(HContentType, ContentJSON)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(nrjson)
	if err != nil {
		log.Printf(ErrorFormat, r.Method, r.URL, err.Error())
	}
}

func SendNormalResopnse(w http.ResponseWriter, r *http.Request,
	sessionID uuid.UUID) {
	nr := &NormalResopnse{
		SessionID: sessionID,
		RetCode:   0,
	}

	SendJsonResponse(w, r, nr)
}

func SendErrorResponse(w http.ResponseWriter, r *http.Request,
	errStatusCode int,
	error string) {
	if error == "" || strings.TrimSpace(error) == "" {
		w.WriteHeader(errStatusCode)
		log.Printf(ErrorFormat, r.Method, r.URL, "nil")
		return
	}

	rsp := &NormalResopnse{
		RetCode: -1,
		Error:   error,
	}

	rspj, err := json.Marshal(rsp)
	if err != nil {
		log.Println("Internal Error, Marshal error : ", err.Error())
		log.Println("Original http status: ",
			errStatusCode,
			", but change to ",
			http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(errStatusCode)
	w.Header().Set(HContentType, ContentJSON)
	_, err = w.Write(rspj)
	if err != nil {
		log.Printf(ErrorFormat, r.Method, r.URL, err.Error())
	} else {
		log.Printf(ErrorFormat, r.Method, r.URL, rsp.Error)
	}
}

func UserGuide(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile(DOCName)
	if err != nil {
		SendErrorResponse(w, r, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set(HContentType, ContentHTML)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		log.Printf(ErrorFormat, r.Method, r.URL, err.Error())
	}
}

func GetCurrentWarn(w http.ResponseWriter, r *http.Request) {
	wgs, err := GetAllWarnMessage()
	if err != nil {
		SendErrorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	RespMap := make(map[string]interface{})
	RespMap["Result"] = wgs
	RespMap["RetCode"] = 0

	SendJsonResponse(w, r, RespMap)
}

func PostMonitorWarn(w http.ResponseWriter, r *http.Request) {
	//get request body

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var ws WarnMessage
	err = json.Unmarshal(b, &ws)
	if err != nil {
		SendErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//check valid

	//save to db or file
	err = SaveWarnMessage(&ws)
	if err != nil {
		SendErrorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	//write response
	SendNormalResopnse(w, r, ws.SessionID)
}
