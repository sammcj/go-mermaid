package validator_test

import (
	"testing"

	"github.com/sammcj/mermaid-check/ast"
	"github.com/sammcj/mermaid-check/validator"
)

func TestNoDuplicateStates(t *testing.T) {
	tests := []struct {
		name       string
		diagram    *ast.StateDiagram
		wantErrors int
	}{
		{
			name: "no duplicates",
			diagram: &ast.StateDiagram{
				Type: "state",
				Statements: []ast.StateStmt{
					&ast.State{ID: "Still", Pos: ast.Position{Line: 2, Column: 1}},
					&ast.State{ID: "Moving", Pos: ast.Position{Line: 3, Column: 1}},
				},
			},
			wantErrors: 0,
		},
		{
			name: "with duplicates",
			diagram: &ast.StateDiagram{
				Type: "state",
				Statements: []ast.StateStmt{
					&ast.State{ID: "Still", Pos: ast.Position{Line: 2, Column: 1}},
					&ast.State{ID: "Still", Pos: ast.Position{Line: 3, Column: 1}},
				},
			},
			wantErrors: 1,
		},
	}

	rule := &validator.NoDuplicateStates{}

	if rule.Name() != "no-duplicate-states" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "no-duplicate-states")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := rule.ValidateState(tt.diagram)
			if len(errors) != tt.wantErrors {
				t.Errorf("ValidateState() errors = %d, want %d", len(errors), tt.wantErrors)
			}
		})
	}
}

func TestValidStateReferences(t *testing.T) {
	tests := []struct {
		name       string
		diagram    *ast.StateDiagram
		wantErrors int
	}{
		{
			name: "valid references",
			diagram: &ast.StateDiagram{
				Type: "state",
				Statements: []ast.StateStmt{
					&ast.State{ID: "Still", Pos: ast.Position{Line: 2, Column: 1}},
					&ast.State{ID: "Moving", Pos: ast.Position{Line: 3, Column: 1}},
					&ast.Transition{From: "Still", To: "Moving", Pos: ast.Position{Line: 4, Column: 1}},
				},
			},
			wantErrors: 0,
		},
		{
			name: "fork and join states",
			diagram: &ast.StateDiagram{
				Type: "state",
				Statements: []ast.StateStmt{
					&ast.Fork{ID: "fork1", Pos: ast.Position{Line: 2, Column: 1}},
					&ast.Join{ID: "join1", Pos: ast.Position{Line: 3, Column: 1}},
					&ast.Choice{ID: "choice1", Pos: ast.Position{Line: 4, Column: 1}},
				},
			},
			wantErrors: 0,
		},
		{
			name: "implicit states from transitions",
			diagram: &ast.StateDiagram{
				Type: "state",
				Statements: []ast.StateStmt{
					&ast.Transition{From: "A", To: "B", Pos: ast.Position{Line: 2, Column: 1}},
				},
			},
			wantErrors: 0,
		},
	}

	rule := &validator.ValidStateReferences{}

	if rule.Name() != "valid-state-references" {
		t.Errorf("Name() = %q, want %q", rule.Name(), "valid-state-references")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := rule.ValidateState(tt.diagram)
			if len(errors) != tt.wantErrors {
				t.Errorf("ValidateState() errors = %d, want %d", len(errors), tt.wantErrors)
			}
		})
	}
}

func TestStateDefaultRules(t *testing.T) {
	rules := validator.StateDefaultRules()
	if len(rules) != 2 {
		t.Errorf("StateDefaultRules() returned %d rules, want 2", len(rules))
	}
}

func TestStateStrictRules(t *testing.T) {
	rules := validator.StateStrictRules()
	if len(rules) != 2 {
		t.Errorf("StateStrictRules() returned %d rules, want 2", len(rules))
	}
}

func TestNewState(t *testing.T) {
	rule := &validator.NoDuplicateStates{}
	v := validator.NewState(rule)
	if v == nil {
		t.Error("NewState() returned nil")
	}
}
