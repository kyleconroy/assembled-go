// Code generated by riza; DO NOT EDIT.

package assembled

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Agents can have statuses (sourced from external systems) that indicate what
// they are currently working on. For example, in a phone system an agent may
// be “available for calls”, “busy on a call”, or “wrapping up on a
// call".
type AgentStatus struct {
	// Identifier for the associated event with the current status.
	EventID string `json:"event_id,omitempty"`

	// Agent's current status from upstream system e.g. 'ready', 'away', or
	// 'busy'. Values are not validated.
	Status string `json:"status,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Identifier for the corresponding agent.
	AgentName string    `json:"agent_name,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

func (r AgentStatus) MarshalJSON() ([]byte, error) {
	type Alias AgentStatus
	return json.Marshal(&struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}{
		Alias:     (Alias)(r),
		EndTime:   timestamp(r.EndTime),
		StartTime: timestamp(r.StartTime),
	})
}

func (r *AgentStatus) UnmarshalJSON(b []byte) error {
	type Alias AgentStatus
	var a struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	*r = AgentStatus(a.Alias)
	r.EndTime = time.Unix(a.EndTime, 0)
	r.StartTime = time.Unix(a.StartTime, 0)
	return nil
}

type CreateAgentStatusRequest struct {
	// Identifier for the associated event with the current status.
	EventID string `json:"event_id,omitempty"`

	// Agent's current status from upstream system e.g. 'ready', 'away', or
	// 'busy'. Values are not validated.
	Status string `json:"status,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Identifier for the corresponding agent.
	AgentName string    `json:"agent_name,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

func (r CreateAgentStatusRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateAgentStatusRequest
	return json.Marshal(&struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}{
		Alias:     (Alias)(r),
		EndTime:   timestamp(r.EndTime),
		StartTime: timestamp(r.StartTime),
	})
}

func (r *CreateAgentStatusRequest) UnmarshalJSON(b []byte) error {
	type Alias CreateAgentStatusRequest
	var a struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	*r = CreateAgentStatusRequest(a.Alias)
	r.EndTime = time.Unix(a.EndTime, 0)
	r.StartTime = time.Unix(a.StartTime, 0)
	return nil
}

type GetAgentStatusRequest struct {
	ID string `json:"id,omitempty"`
}

// Returns GetAgentStatusRequest with ID set to the empty string so that it's
// not included in the JSON request body.
func (r *GetAgentStatusRequest) body() interface{} {
	if r == nil {
		return r
	}
	req := *r
	req.ID = ""
	return &req
}

func (c *Client) CreateAgentStatus(ctx context.Context, r *CreateAgentStatusRequest) (*AgentStatus, error) {
	var resp AgentStatus
	if err := c.request(ctx, "POST", "/v0/agents/status", nil, r, &resp); err != nil {
		return nil, fmt.Errorf("CreateAgentStatus: %w", err)
	}
	return &resp, nil
}

func (c *Client) GetAgentStatus(ctx context.Context, r *GetAgentStatusRequest) (*AgentStatus, error) {
	var resp AgentStatus
	if err := c.request(ctx, "GET", fmt.Sprintf("/v0/agents/%s/status", r.ID), nil, nil, &resp); err != nil {
		return nil, fmt.Errorf("GetAgentStatus: %w", err)
	}
	return &resp, nil
}
