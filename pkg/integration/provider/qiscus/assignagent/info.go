package assignagent

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dionofrizal88/go-allocator/pkg/rest"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// InfoRequest is a request contract.
type InfoRequest struct {
	RoomID  string `json:"room_id"`
	AgentID int    `json:"agent_id"`
}

// DataResponse is a data response contract.
type DataResponse struct {
	AddedAgent struct {
		AvatarUrl    interface{} `json:"avatar_url"`
		CreatedAt    string      `json:"created_at"`
		Email        string      `json:"email"`
		ForceOffline bool        `json:"force_offline"`
		Id           interface{} `json:"id"`
		IsAvailable  bool        `json:"is_available"`
		IsVerified   bool        `json:"is_verified"`
		LastLogin    time.Time   `json:"last_login"`
		Name         string      `json:"name"`
		SdkEmail     string      `json:"sdk_email"`
		SdkKey       string      `json:"sdk_key"`
		Type         int         `json:"type"`
		TypeAsString string      `json:"type_as_string"`
		UpdatedAt    string      `json:"updated_at"`
	} `json:"added_agent"`
	Service struct {
		CreatedAt             string      `json:"created_at"`
		FirstCommentId        string      `json:"first_comment_id"`
		FirstCommentTimestamp interface{} `json:"first_comment_timestamp"`
		IsResolved            bool        `json:"is_resolved"`
		LastCommentId         string      `json:"last_comment_id"`
		Notes                 interface{} `json:"notes"`
		ResolvedAt            interface{} `json:"resolved_at"`
		RetrievedAt           time.Time   `json:"retrieved_at"`
		RoomId                string      `json:"room_id"`
		UpdatedAt             string      `json:"updated_at"`
		UserId                interface{} `json:"user_id"`
	} `json:"service"`
}

// InfoResponse is a response contract.
type InfoResponse struct {
	Meta struct {
		After      interface{} `json:"after"`
		Before     interface{} `json:"before"`
		PerPage    int         `json:"per_page"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Status int          `json:"status"`
	Data   DataResponse `json:"data"`
}

// Info is for getting info of actor certificate detail.
func (p *AssignAgent) Info(ctx context.Context) (interface{}, error) {
	req := rest.NewHTTPRequest()

	req.Data.Header = http.Header{}
	req.Data.Header.Set("Content-Type", "application/json")
	req.Data.Header.Set("Qiscus-App-Id", p.AppConfig.QiscusAppID)
	req.Data.Header.Set("Qiscus-Secret-Key", p.AppConfig.QiscusSecretKey)
	req.Method = rest.HTTPMethodPOST
	req.Data.URL = fmt.Sprintf("%s%s", p.AppConfig.QiscusBaseURL, "api/v1/admin/service/assign_agent")

	data := InfoRequest{
		RoomID:  p.InputAssignAgent.RoomID,
		AgentID: p.InputAssignAgent.AgentID,
	}
	jsonData, err := json.Marshal(&data)
	if err != nil {
		log.Printf("error marshal json request, err: %v", err)
		return nil, err
	}
	req.Data.Body = jsonData

	res, err := req.Exec(ctx)
	if err != nil {
		log.Printf("Error send assign agent through to qiscus provider, err: %v", err)

		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error read body, err: %v", err)

		return nil, err
	}

	var response InfoResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error decoding response, err: %v", err)
		log.Printf("response: %v", string(body))

		return nil, err
	}

	return &response, err
}
