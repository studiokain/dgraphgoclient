package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNQuad(t *testing.T) {
	if err := checkNQuad("", "name", "", "Alice"); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "", "", "Alice"); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "name", "", ""); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "name", "id", "Alice"); err == nil {
		t.Fatal(err)
	}
}

func TestSetMutation(t *testing.T) {
	req := NewRequest()

	if err := req.SetMutation("alice", "name", "", "Alice", ""); err != nil {
		t.Fatal(err)
	}
	if err := req.SetMutation("alice", "falls.in", "", "rabbithole", ""); err != nil {
		t.Fatal(err)
	}
	if err := req.DelMutation("alice", "falls.in", "", "rabbithole", ""); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(req.gr.Mutation.Set), 2, "Set should have 2 entries")
	assert.Equal(t, len(req.gr.Mutation.Del), 1, "Del should have 1 entry")
}
