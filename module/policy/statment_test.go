package policy

import "testing"

func TestNothingToMerge(t *testing.T) {
	stmts := []Statement{
		Statement{
			Effect:  allow,
			Service: "EC2",
			Actions: []string{"StartInstance"},
		},
		Statement{
			Effect:  allow,
			Service: "S3",
			Actions: []string{"CreateBucket"},
		},
	}

	pruned, err := PruneStatements(stmts)
	if err != nil {
		t.Fatalf("Pruning failed: %v", err)
	}

	if len(pruned) != 2 {
		t.Fatalf("Expected 2 statements, got %d", len(pruned))
	}
}

func TestMergeTwoStatements(t *testing.T) {
	stmts := []Statement{
		Statement{
			Effect:  allow,
			Service: "EC2",
			Actions: []string{"StartInstance"},
		},
		Statement{
			Effect:  allow,
			Service: "EC2",
			Actions: []string{"CreateInstance"},
		},
	}

	pruned, err := PruneStatements(stmts)
	if err != nil {
		t.Fatalf("Pruning failed: %v", err)
	}

	if len(pruned) != 1 {
		t.Fatalf("Expected 1 statements, got %d", len(pruned))
	}
	if len(pruned[0].Actions) != 2 {
		t.Fatalf("Expected to have two actions, got %d", len(pruned[0].Actions))
	}
}
