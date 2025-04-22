package activeagent

import "fmt"

// InputActiveAgent is actor that used for performing know active agent.
type InputActiveAgent struct {
}

// SetInput is uses to set required data.
func (p *ActiveAgent) SetInput(a interface{}) error {
	_, ok := a.(InputActiveAgent)
	if !ok {
		return fmt.Errorf("invalid input")
	}

	return nil
}

// SetRequestID is for setting request id.
func (p *ActiveAgent) SetRequestID(requestID string) {
	p.RequestID = requestID
}
