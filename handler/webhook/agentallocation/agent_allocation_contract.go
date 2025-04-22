package agentallocation

import (
	"context"
	"time"
)

var (
	RecoveryRequestEmailTemplate = "Hello %s,\n\nWe received a request to reset your password for your Digital Sekuriti Indonesia account. If you made this request, please click the link below to reset your password:\n\n🔗 %s\n\nThis link is valid for %s. If you did not request a password reset, please ignore this email or contact our support team.\n\nFor security reasons, do not share this link with anyone.\n\nBest regards,\nDigital Sekuriti Indonesia\nadmin@digitalsekuriti.id | https://digitalsekuriti.id/"
)

// Request struct is used to get request value.
type Request struct {
	AppId          string `json:"app_id"`
	AvatarUrl      string `json:"avatar_url"`
	CandidateAgent struct {
		AvatarUrl    interface{} `json:"avatar_url"`
		CreatedAt    time.Time   `json:"created_at"`
		Email        string      `json:"email"`
		ForceOffline bool        `json:"force_offline"`
		Id           int         `json:"id"`
		IsAvailable  bool        `json:"is_available"`
		IsVerified   bool        `json:"is_verified"`
		LastLogin    interface{} `json:"last_login"`
		Name         string      `json:"name"`
		SdkEmail     string      `json:"sdk_email"`
		SdkKey       string      `json:"sdk_key"`
		Type         int         `json:"type"`
		TypeAsString string      `json:"type_as_string"`
		UpdatedAt    time.Time   `json:"updated_at"`
	} `json:"candidate_agent"`
	Email         string      `json:"email"`
	Extras        string      `json:"extras"`
	IsNewSession  bool        `json:"is_new_session"`
	IsResolved    bool        `json:"is_resolved"`
	LatestService interface{} `json:"latest_service"`
	Name          string      `json:"name"`
	RoomId        string      `json:"room_id"`
	Source        string      `json:"source"`
}

// Response struct is used to get response value.
type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

// transformToResponse is a function to transform user into response value.
func (co *Controller) transformToResponse(ctx context.Context, message string) Response {
	var response Response
	response.Message = message
	response.Data = nil

	return response
}
