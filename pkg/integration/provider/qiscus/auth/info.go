package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dionofrizal88/go-allocator/pkg/rest"
	"io/ioutil"
	"log"
	"net/http"
)

// InfoRequest is a request contract.
type InfoRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse is a user response contract.
type UserResponse struct {
	Id                  interface{} `json:"id"`
	Name                string      `json:"name"`
	Email               string      `json:"email"`
	AuthenticationToken string      `json:"authentication_token"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	SdkEmail            string      `json:"sdk_email"`
	SdkKey              string      `json:"sdk_key"`
	IsAvailable         bool        `json:"is_available"`
	Type                int         `json:"type"`
	AvatarUrl           string      `json:"avatar_url"`
	AppId               interface{} `json:"app_id"`
	IsVerified          bool        `json:"is_verified"`
	NotificationsRoomId interface{} `json:"notifications_room_id"`
	BubbleColor         interface{} `json:"bubble_color"`
	QismoKey            string      `json:"qismo_key"`
	DirectLoginToken    string      `json:"direct_login_token"`
	LastLogin           string      `json:"last_login"`
	ForceOffline        bool        `json:"force_offline"`
	DeletedAt           interface{} `json:"deleted_at"`
	IsTocAgree          bool        `json:"is_toc_agree"`
	TotpToken           string      `json:"totp_token"`
	IsReqOtpReset       interface{} `json:"is_req_otp_reset"`
	LastPasswordUpdate  string      `json:"last_password_update"`
	TypeAsString        string      `json:"type_as_string"`
	AssignedRules       interface{} `json:"assigned_rules"`
}

// DataResponse is a data response contract.
type DataResponse struct {
	User UserResponse `json:"user"`
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

// Info is for getting info of actor detail.
func (p *Auth) Info(ctx context.Context) (interface{}, error) {
	req := rest.NewHTTPRequest()

	req.Data.Header = http.Header{}
	req.Data.Header.Set("Content-Type", "application/json")
	req.Method = rest.HTTPMethodPOST
	req.Data.URL = fmt.Sprintf("%s%s", p.AppConfig.QiscusBaseURL, "api/v1/auth")

	data := InfoRequest{
		Email:    p.InputAuth.Email,
		Password: p.InputAuth.Password,
	}
	jsonData, err := json.Marshal(&data)
	if err != nil {
		log.Printf("error marshal json request, err: %v", err)
		return nil, err
	}
	req.Data.Body = jsonData

	res, err := req.Exec(ctx)
	if err != nil {
		log.Printf("Error send auth through to qiscus provider, err: %v", err)

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
