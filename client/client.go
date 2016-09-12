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

func NewRequest() Req {
	return Req{}
}
func (req *Req) Request() *graph.Request {
	return &req.gr
}

func checkNQuad(sub, pred, objId, objVal string) error {
	if len(sub) == 0 {
		return fmt.Errorf("Subject can't be empty")
	}
	if len(pred) == 0 {
		return fmt.Errorf("Predicate can't be empty")
	}
	if len(objId) == 0 && len(objVal) == 0 {
		return fmt.Errorf("Both objectId and objectValue can't be nil")
	}
	if len(objId) > 0 && len(objVal) > 0 {
		return fmt.Errorf("Only one out of objectId and objectValue can be set")
	}
	return nil
}

func (req *Req) SetQuery(q string) {
	req.gr.Query = q
}

func (req *Req) SetMutation(sub, pred, objId, objVal, label string) error {
	if err := checkNQuad(sub, pred, objId, objVal); err != nil {
		return err
	}

	if req.gr.Mutation == nil {
		req.gr.Mutation = new(graph.Mutation)
	}

	req.gr.Mutation.Set = append(req.gr.Mutation.Set, &graph.NQuad{
		Sub:    sub,
		Pred:   pred,
		ObjId:  objId,
		ObjVal: []byte(objVal),
		Label:  label,
	})
	return nil
}

func (req *Req) DelMutation(sub, pred, objId, objVal, label string) error {
	if err := checkNQuad(sub, pred, objId, objVal); err != nil {
		return err
	}

	if req.gr.Mutation == nil {
		req.gr.Mutation = new(graph.Mutation)
	}

	req.gr.Mutation.Del = append(req.gr.Mutation.Del, &graph.NQuad{
		Sub:    sub,
		Pred:   pred,
		ObjId:  objId,
		ObjVal: []byte(objVal),
		Label:  label,
	})
	return nil
}
