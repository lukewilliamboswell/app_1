// Code generated by goa v3.1.2, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen backend-api/design

package client

import (
	calc "backend-api/gen/calc"
	calcpb "backend-api/gen/grpc/calc/pb"
	"encoding/json"
	"fmt"
)

// BuildAddPayload builds the payload for the calc add endpoint from CLI flags.
func BuildAddPayload(calcAddMessage string) (*calc.AddPayload, error) {
	var err error
	var message calcpb.AddRequest
	{
		if calcAddMessage != "" {
			err = json.Unmarshal([]byte(calcAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 8399553735696626949,\n      \"b\": 360622074634248926\n   }'")
			}
		}
	}
	v := &calc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}