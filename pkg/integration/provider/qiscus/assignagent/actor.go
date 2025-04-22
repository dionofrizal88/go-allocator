package assignagent

import "fmt"

// InputAssignAgent is actor that used for performing know auth.
type InputAssignAgent struct {
	RoomID  string `json:"room_id"`
	AgentID int    `json:"agent_id"`
}

// SetInput is uses to set required data.
func (p *AssignAgent) SetInput(a interface{}) error {
	input, ok := a.(InputAssignAgent)
	if !ok {
		return fmt.Errorf("invalid input")
	}

	p.InputAssignAgent = InputAssignAgent{
		RoomID:  input.RoomID,
		AgentID: input.AgentID,
	}

	return nil
}

// SetRequestID is for setting request id.
func (p *AssignAgent) SetRequestID(requestID string) {
	p.RequestID = requestID
}
