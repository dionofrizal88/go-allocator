package auth

import "fmt"

// InputAuth is actor that used for performing know auth.
type InputAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SetInput is uses to set required data.
func (p *Auth) SetInput(a interface{}) error {
	input, ok := a.(InputAuth)
	if !ok {
		return fmt.Errorf("invalid input")
	}

	p.InputAuth = InputAuth{
		Email:    input.Email,
		Password: input.Password,
	}

	return nil
}

// SetRequestID is for setting request id.
func (p *Auth) SetRequestID(requestID string) {
	p.RequestID = requestID
}
