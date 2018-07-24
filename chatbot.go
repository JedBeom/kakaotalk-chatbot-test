package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Keyboard struct {
	Type    string   `json:"type"`
	Buttons []string `json:"buttons"`
}

type PostedMessage struct {
	User_key string `json:"user_key"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

type MessageInside struct {
	Text string `json:"text"`
	//Photo          Photo         `json:"photo"`
	//Message_button MessageButton `json:"message_button"`
}

type MessageBody struct {
	Message  MessageInside `json:"message"`
	Keyboard Keyboard      `json:"keyboard"`
}

type MessageButton struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Photo struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func timestampNow() (timenow string) {
	t := time.Now()
	timenow = fmt.Sprintf("%v-%v-%v %v:%v:%v", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
	return
}

func handlerKeyboard(w http.ResponseWriter, r *http.Request) {
	buttonarray := []string{"월요일", "화요일", "수요일", "목요일", "금요일"}
	keyboard := Keyboard{Type: "buttons", Buttons: buttonarray}
	js, err := json.MarshalIndent(&keyboard, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//log.Println(keyboard)
	w.Write(js)
}

func handlerMessage(w http.ResponseWriter, r *http.Request) {
	var postedMessage PostedMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&postedMessage)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	//log.Println("MessageHandler")
	//log.Println("type:", postedMessage.Type)
	//log.Println("content:", postedMessage.Content)
	if postedMessage.Type != "text" {
		w.WriteHeader(400)
		return
	}
	//임시 답장
	text := fmt.Sprintf("%v의 급식이 없는 것 같네요...ㅠ", postedMessage.Content)
	message := MessageInside{Text: text}
	buttonarray := []string{"월요일", "화요일", "수요일", "목요일", "금요일"}
	keyboard := Keyboard{Type: "buttons", Buttons: buttonarray}
	messageBody := MessageBody{Message: message, Keyboard: keyboard}
	//log.Println("Message:", messageBody)
	js, err2 := json.MarshalIndent(&messageBody, "", "\t")
	//log.Println("Err:", err)
	if err != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(js)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/keyboard", handlerKeyboard)
	http.HandleFunc("/message", handlerMessage)
	http.HandleFunc("/", sayHello)
	/*
		server := http.Server{
			Addr:    ":443",
			Handler: nil,
		}
		server.ListenAndServeTLS("/etc/letsencrypt/live/inmull.xyz/fullchain.pem", "/etc/letsencrypt/live/inmull.xyz/privkey.pem")
	*/
	http.ListenAndServe(":80", nil)
}
