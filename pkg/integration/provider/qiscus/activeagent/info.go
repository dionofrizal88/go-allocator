package activeagent

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dionofrizal88/go-allocator/pkg/rest"
	"io/ioutil"
	"log"
	"net/http"
)

// AgentResponse is an agent response contract.
type AgentResponse struct {
	AvatarUrl            string      `json:"avatar_url"`
	CreatedAt            string      `json:"created_at"`
	CurrentCustomerCount int         `json:"current_customer_count"`
	Email                string      `json:"email"`
	ForceOffline         bool        `json:"force_offline"`
	Id                   int         `json:"id"`
	IsAvailable          bool        `json:"is_available"`
	LastLogin            interface{} `json:"last_login"`
	Name                 string      `json:"name"`
	SdkEmail             string      `json:"sdk_email"`
	SdkKey               string      `json:"sdk_key"`
	Type                 int         `json:"type"`
	TypeAsString         string      `json:"type_as_string"`
	UserChannels         []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user_channels"`
	UserRoles []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"user_roles"`
}

// DataResponse is an data response contract.
type DataResponse struct {
	Agents []AgentResponse `json:"agents"`
}

// InfoResponse is a response contract.
type InfoResponse struct {
	Meta struct {
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Status int          `json:"status"`
	Data   DataResponse `json:"data"`
}

// Info is for getting info of actor certificate detail.
func (p *ActiveAgent) Info(ctx context.Context) (interface{}, error) {
	req := rest.NewHTTPRequest()

	req.Data.Header = http.Header{}
	req.Data.Header.Set("Content-Type", "application/json")
	req.Data.Header.Set("Authorization", p.Authorization)
	req.Data.Header.Set("Qiscus-App-Id", p.AppConfig.QiscusAppID)
	req.Method = rest.HTTPMethodGET
	req.Data.URL = fmt.Sprintf("%s%s", p.AppConfig.QiscusBaseURL, "api/v2/admin/agents?page=1&limit=100")

	res, err := req.Exec(ctx)
	if err != nil {
		log.Printf("Error send get active agent through to qiscus provider, err: %v", err)

		return nil, err
	}

	//if res.StatusCode != http.StatusOK {
	//
	//}

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
