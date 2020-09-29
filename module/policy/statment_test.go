package policy

import "testing"

func TestNothingToMerge(t *testing.T) {
	reqs := []Request{
		Request{
			apiKey: "EC2/StartInstance",
			body:   "",
		},
		Request{
			apiKey: "s3/CreateBucket",
			body:   "",
		},
	}

	stmts := NewStatements()
	err := stmts.AddRequests(reqs)
	if err != nil {
		t.Fatalf("adding failed: %v", err)
	}
	s := stmts.List()

	if len(s) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(s))
	}
}

func TestMergeTwoStatements(t *testing.T) {
	reqs := []Request{
		Request{
			apiKey: "EC2/StartInstance",
			body:   "",
		},
		Request{
			apiKey: "EC2/CreateInstance",
			body:   "",
		},
	}

	stmts := NewStatements()
	err := stmts.AddRequests(reqs)
	if err != nil {
		t.Fatalf("adding failed: %v", err)
	}

	s := stmts.List()
	if len(s) != 1 {
		t.Fatalf("Expected 1 statements, got %d", len(s))
	}
	if len(s[0].Actions) != 2 {
		t.Fatalf("Expected to have two actions, got %d", len(s[0].Actions))
	}
}

func TestTwoOfTheSameActions(t *testing.T) {
	reqs := []Request{
		Request{
			apiKey: "EC2/StartInstance",
			body:   "",
		},
		Request{
			apiKey: "EC2/StartInstance",
			body:   "",
		},
	}

	stmts := NewStatements()
	err := stmts.AddRequests(reqs)
	if err != nil {
		t.Fatalf("adding failed: %v", err)
	}
	s := stmts.List()
	if len(s[0].Actions) != 1 {
		t.Fatalf("Expected 1 action, got %d", len(s[0].Actions))
	}
}
