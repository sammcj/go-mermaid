package validator_test

import (
	"testing"

	"github.com/sammcj/mermaid-check/ast"
	"github.com/sammcj/mermaid-check/validator"
)

func TestValidDirection(t *testing.T) {
	rule := &validator.ValidDirection{}

	if rule.Name() != "valid-direction" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "valid-direction")
	}

	tests := []struct {
		name      string
		direction string
		wantError bool
	}{
		{"TB valid", "TB", false},
		{"TD valid", "TD", false},
		{"BT valid", "BT", false},
		{"RL valid", "RL", false},
		{"LR valid", "LR", false},
		{"invalid direction", "XY", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowchart := &ast.Flowchart{
				Type:      "flowchart",
				Direction: tt.direction,
				Pos:       ast.Position{Line: 1, Column: 1},
			}

			errors := rule.Validate(flowchart)
			if tt.wantError && len(errors) == 0 {
				t.Error("expected validation error, got none")
			}
			if !tt.wantError && len(errors) > 0 {
				t.Errorf("unexpected validation error: %v", errors)
			}
		})
	}
}

func TestNoUndefinedNodes(t *testing.T) {
	rule := &validator.NoUndefinedNodes{}

	if rule.Name() != "no-undefined-nodes" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "no-undefined-nodes")
	}

	t.Run("all nodes defined", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node A", Pos: ast.Position{Line: 2}},
				&ast.NodeDef{ID: "B", Label: "Node B", Pos: ast.Position{Line: 3}},
				&ast.Link{From: "A", To: "B", Arrow: "-->", Pos: ast.Position{Line: 4}},
			},
		}

		errors := rule.Validate(flowchart)
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors: %v", errors)
		}
	})

	t.Run("undefined node in link", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node A", Pos: ast.Position{Line: 2}},
				&ast.Link{From: "A", To: "C", Arrow: "-->", Pos: ast.Position{Line: 3}},
			},
		}

		errors := rule.Validate(flowchart)
		// Nodes referenced in links are implicitly defined, so this should pass
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors: %v", errors)
		}
	})
}

func TestNoParenthesesInLabels(t *testing.T) {
	rule := &validator.NoParenthesesInLabels{}

	if rule.Name() != "no-parentheses-in-labels" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "no-parentheses-in-labels")
	}

	tests := []struct {
		name      string
		label     string
		wantError bool
	}{
		{"no parentheses", "Simple Label", false},
		{"with parentheses", "Label (with note)", true},
		{"opening paren only", "Label (incomplete", true},
		{"closing paren only", "Label incomplete)", true},
		{"empty label", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flowchart := &ast.Flowchart{
				Type:      "flowchart",
				Direction: "TD",
				Statements: []ast.Statement{
					&ast.NodeDef{ID: "A", Label: tt.label, Pos: ast.Position{Line: 2}},
				},
			}

			errors := rule.Validate(flowchart)
			if tt.wantError && len(errors) == 0 {
				t.Error("expected validation error, got none")
			}
			if !tt.wantError && len(errors) > 0 {
				t.Errorf("unexpected validation error: %v", errors)
			}
		})
	}

	t.Run("nested in subgraph", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.Subgraph{
					Title: "Test Subgraph",
					Statements: []ast.Statement{
						&ast.NodeDef{ID: "A", Label: "Label (with paren)", Pos: ast.Position{Line: 3}},
					},
					Pos: ast.Position{Line: 2},
				},
			},
		}

		errors := rule.Validate(flowchart)
		if len(errors) == 0 {
			t.Error("expected validation error for node in subgraph")
		}
	})
}

func TestNoDuplicateNodeIDs(t *testing.T) {
	rule := &validator.NoDuplicateNodeIDs{}

	if rule.Name() != "no-duplicate-node-ids" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "no-duplicate-node-ids")
	}

	t.Run("no duplicates", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node A", Pos: ast.Position{Line: 2}},
				&ast.NodeDef{ID: "B", Label: "Node B", Pos: ast.Position{Line: 3}},
				&ast.NodeDef{ID: "C", Label: "Node C", Pos: ast.Position{Line: 4}},
			},
		}

		errors := rule.Validate(flowchart)
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors: %v", errors)
		}
	})

	t.Run("duplicate node IDs", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "First A", Pos: ast.Position{Line: 2}},
				&ast.NodeDef{ID: "B", Label: "Node B", Pos: ast.Position{Line: 3}},
				&ast.NodeDef{ID: "A", Label: "Second A", Pos: ast.Position{Line: 4}},
			},
		}

		errors := rule.Validate(flowchart)
		if len(errors) != 1 {
			t.Errorf("expected 1 validation error, got %d", len(errors))
		}
		if len(errors) > 0 && !contains(errors[0].Message, "duplicate") {
			t.Errorf("expected 'duplicate' in error message, got: %s", errors[0].Message)
		}
	})

	t.Run("duplicates in subgraph", func(t *testing.T) {
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Outer A", Pos: ast.Position{Line: 2}},
				&ast.Subgraph{
					Title: "Test",
					Statements: []ast.Statement{
						&ast.NodeDef{ID: "A", Label: "Inner A", Pos: ast.Position{Line: 4}},
					},
					Pos: ast.Position{Line: 3},
				},
			},
		}

		errors := rule.Validate(flowchart)
		if len(errors) != 1 {
			t.Errorf("expected 1 validation error for duplicate across subgraph, got %d", len(errors))
		}
	})
}

func TestValidator(t *testing.T) {
	t.Run("default rules", func(t *testing.T) {
		v := validator.New(validator.DefaultRules()...)

		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node A", Pos: ast.Position{Line: 2}},
				&ast.Link{From: "A", To: "B", Arrow: "-->", Pos: ast.Position{Line: 3}},
			},
		}

		errors := v.Validate(flowchart)
		// Should pass with default rules
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors with default rules: %v", errors)
		}
	})

	t.Run("strict rules", func(t *testing.T) {
		v := validator.New(validator.StrictRules()...)

		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node (with paren)", Pos: ast.Position{Line: 2}},
			},
		}

		errors := v.Validate(flowchart)
		// Should fail with strict rules due to parentheses in label
		if len(errors) == 0 {
			t.Error("expected validation error with strict rules")
		}
	})
}

func TestValidationErrorString(t *testing.T) {
	err := validator.ValidationError{
		Line:     42,
		Column:   10,
		Message:  "test error message",
		Severity: validator.SeverityError,
	}

	result := err.Error()
	if !contains(result, "line 42") {
		t.Errorf("expected 'line 42' in error string, got: %s", result)
	}
	if !contains(result, "error") {
		t.Errorf("expected 'error' in error string, got: %s", result)
	}
	if !contains(result, "test error message") {
		t.Errorf("expected 'test error message' in error string, got: %s", result)
	}
}

func TestSeverityString(t *testing.T) {
	tests := []struct {
		severity validator.Severity
		want     string
	}{
		{validator.SeverityError, "error"},
		{validator.SeverityWarning, "warning"},
		{validator.SeverityInfo, "info"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.severity.String(); got != tt.want {
				t.Errorf("Severity.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && findSubstring(s, substr)))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestNewGeneric(t *testing.T) {
	v := validator.NewGeneric()
	if v == nil {
		t.Error("NewGeneric() returned nil")
	}
}

func TestNewSequence(t *testing.T) {
	v := validator.NewSequence()
	if v == nil {
		t.Error("NewSequence() returned nil")
	}
}

func TestDefaultRules(t *testing.T) {
	rules := validator.DefaultRules()
	if len(rules) == 0 {
		t.Error("DefaultRules() returned empty rules")
	}
}

func TestStrictRules(t *testing.T) {
	rules := validator.StrictRules()
	if len(rules) == 0 {
		t.Error("StrictRules() returned empty rules")
	}
}

func TestValidateDiagram(t *testing.T) {
	t.Run("flowchart with default validator", func(t *testing.T) {
		v := validator.NewGeneric()
		flowchart := &ast.Flowchart{
			Type:      "flowchart",
			Direction: "TD",
			Statements: []ast.Statement{
				&ast.NodeDef{ID: "A", Label: "Node A", Pos: ast.Position{Line: 2}},
			},
		}

		errors := v.ValidateDiagram(flowchart)
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors: %v", errors)
		}
	})

	t.Run("sequence with sequence validator", func(t *testing.T) {
		v := validator.NewSequence()
		seq := &ast.SequenceDiagram{
			Type: "sequence",
			Statements: []ast.SeqStmt{
				&ast.Participant{ID: "Alice", Pos: ast.Position{Line: 2}},
			},
		}

		errors := v.ValidateDiagram(seq)
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors: %v", errors)
		}
	})

	t.Run("unsupported diagram type", func(t *testing.T) {
		v := validator.NewGeneric()
		generic := ast.NewGenericDiagram("unknown", "unknown diagram", ast.Position{Line: 1})

		errors := v.ValidateDiagram(generic)
		// Should handle gracefully
		if len(errors) > 0 {
			t.Errorf("unexpected validation errors for generic: %v", errors)
		}
	})
}
