package client

import (
	"fmt"

	"github.com/dgraph-io/dgraphgoclient/graph"
)

// Req wraps the graph.Request so that we can define helper methods for the
// client around it.
type Req struct {
	gr graph.Request
}

// NewRequest initializes and returns a new request which can be used to query
// or perform set/delete mutations.
func NewRequest() Req {
	return Req{}
}

// Request returns the graph request object which is sent to the server to perform
// a query/mutation.
func (req *Req) Request() *graph.Request {
	return &req.gr
}

func checkNQuad(sub, pred, objId string, objVal Value) error {
	if len(sub) == 0 {
		return fmt.Errorf("Subject can't be empty")
	}
	if len(pred) == 0 {
		return fmt.Errorf("Predicate can't be empty")
	}
	hasVal := objVal != nil && objVal.Val != nil
	if len(objId) == 0 && !hasVal {
		return fmt.Errorf("Both objectId and objectValue can't be nil")
	}
	if len(objId) > 0 && hasVal {
		return fmt.Errorf("Only one out of objectId and objectValue can be set")
	}
	return nil
}

// SetQuery sets a query as part of the request.
func (req *Req) SetQuery(q string) {
	req.gr.Query = q
}

// SetMutation adds a set mutation operation.
func (req *Req) SetMutation(sub, pred, objId string, value Value, label string) error {
	if err := checkNQuad(sub, pred, objId, value); err != nil {
		return err
	}

	if req.gr.Mutation == nil {
		req.gr.Mutation = new(graph.Mutation)
	}

	req.gr.Mutation.Set = append(req.gr.Mutation.Set, &graph.NQuad{
		Sub:   sub,
		Pred:  pred,
		ObjId: objId,
		Value: value,
		Label: label,
	})
	return nil
}

// DelMutation adds a delete mutation operation.
func (req *Req) DelMutation(sub, pred, objId string, value Value, label string) error {
	if err := checkNQuad(sub, pred, objId, value); err != nil {
		return err
	}

	if req.gr.Mutation == nil {
		req.gr.Mutation = new(graph.Mutation)
	}

	req.gr.Mutation.Del = append(req.gr.Mutation.Del, &graph.NQuad{
		Sub:   sub,
		Pred:  pred,
		ObjId: objId,
		Value: value,
		Label: label,
	})
	return nil
}
