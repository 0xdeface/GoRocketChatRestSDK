package rocketchat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID       string `json:"userId"`
	Token    string `json:"authToken"`
	Username string `json:"username"`
}
type Message struct {
	ID string `json:"_id"`
	//rid	:	wojkjJkSQMFRCNngq
	Msg       string `json:"msg"`
	Ts        string `json:"ts"`
	ReplyID   string `json:"tmid"`
	User      User   `json:"u"`
	UpdatedAt string `json:"_updatedAt"`
}
type Messages struct {
	Messages []Message `json:"messages"`
}

type requestSettings struct {
	Method  string
	ApiUrl  string
	Payload interface{}
}

type GroupCreateSettings struct {
	Members  []string
	ReadOnly bool
}

type RocketChat struct {
	host        string
	email       string
	password    string
	currentUser *User
	Cancel      context.CancelFunc
}

const (
	loginUrl        = "api/v1/login"
	groupListUrl    = "api/v1/groups.list"
	postMessageUrl  = "api/v1/chat.postMessage"
	imHistoryUrl    = "api/v1/im.history"
	groupHistoryUrl = "api/v1/groups.history"
	groupCreateUrl  = "api/v1/groups.create"
	groupDeleteUrl  = "api/v1/groups.delete"
	successStatus   = "success"
)

func CreateRocketChat(host, email, password string) *RocketChat {
	rc := &RocketChat{host, email, password, &User{}, nil}
	rc.login()
	return rc
}

// login - this method gets token from RocketChat server
// it call implicitly when you create RocketChat object
func (r *RocketChat) login() {
	fmt.Println(r.host + "/" + loginUrl)
	requestBody, err := json.Marshal(map[string]string{"user": r.email, "password": r.password})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post(r.host+"/"+loginUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]json.RawMessage
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	err = json.Unmarshal(result["status"], &status)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(result["data"], r.currentUser)
	if err != nil {
		log.Fatal(err)
	}

}

// prepareRequest - internal function for making api request
func (r *RocketChat) prepareRequest(settings requestSettings) *http.Request {
	body, err := json.Marshal(settings.Payload)
	if err != nil {
		log.Fatal(err)
	}
	request, err := http.NewRequest(settings.Method, fmt.Sprintf("%v/%v", r.host, settings.ApiUrl), bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("X-Auth-Token", r.currentUser.Token)
	request.Header.Add("X-User-Id", r.currentUser.ID)
	return request
}

// GroupList - Lists all of the private groups the calling user has joined.
func (r *RocketChat) GroupList() {
	request := r.prepareRequest(requestSettings{Method: "GET", ApiUrl: groupListUrl})
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// ChatSendMessage - send message to one channel
func (r *RocketChat) ChatSendMessage(to, text string) {
	data := map[string]string{"channel": to, "text": text}
	request := r.prepareRequest(requestSettings{Method: "POST", ApiUrl: postMessageUrl, Payload: data})

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// GroupHistory Retrieves the messages from a private group,
// only if you're part of the group.
func (r *RocketChat) GroupHistory() *Messages {
	request := r.prepareRequest(requestSettings{Method: "GET", ApiUrl: groupHistoryUrl + "?roomId=wojkjJkSQMFRCNngq"})
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	messages := &Messages{}
	json.Unmarshal(body, messages)
	return messages
}

// GroupCreate - Create a private channel.
func (r *RocketChat) GroupCreate(channelName string, optional *GroupCreateSettings) string {
	data := map[string]interface{}{"name": channelName, "members": optional.Members, "readOnly": optional.ReadOnly}
	request := r.prepareRequest(requestSettings{Method: "POST", ApiUrl: groupCreateUrl, Payload: data})
	client := &http.Client{}
	b, err := request.GetBody()
	body, err := ioutil.ReadAll(b)
	fmt.Println(string(body))
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return channelName
}

// GroupDelete Remove a private channel.
func (r *RocketChat) GroupDelete() {

}
