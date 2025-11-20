package ast

import (
	"testing"
)

func TestStateDiagram_GetType(t *testing.T) {
	tests := []struct {
		name string
		typ  string
	}{
		{"state", "state"},
		{"stateDiagram-v2", "stateDiagram-v2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &StateDiagram{Type: tt.typ}
			if got := sd.GetType(); got != tt.typ {
				t.Errorf("GetType() = %v, want %v", got, tt.typ)
			}
		})
	}
}

func TestStateDiagram_GetPosition(t *testing.T) {
	pos := Position{Line: 1, Column: 1}
	sd := &StateDiagram{Pos: pos}
	if got := sd.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestStateDiagram_GetSource(t *testing.T) {
	source := "stateDiagram-v2\n    [*] --> State1"
	sd := &StateDiagram{Source: source}
	if got := sd.GetSource(); got != source {
		t.Errorf("GetSource() = %v, want %v", got, source)
	}
}

func TestState_GetPosition(t *testing.T) {
	pos := Position{Line: 2, Column: 5}
	s := &State{Pos: pos}
	if got := s.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestTransition_GetPosition(t *testing.T) {
	pos := Position{Line: 3, Column: 5}
	tr := &Transition{Pos: pos}
	if got := tr.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestStartState_GetPosition(t *testing.T) {
	pos := Position{Line: 4, Column: 5}
	ss := &StartState{Pos: pos}
	if got := ss.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestEndState_GetPosition(t *testing.T) {
	pos := Position{Line: 5, Column: 5}
	es := &EndState{Pos: pos}
	if got := es.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestFork_GetPosition(t *testing.T) {
	pos := Position{Line: 6, Column: 5}
	f := &Fork{Pos: pos}
	if got := f.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestJoin_GetPosition(t *testing.T) {
	pos := Position{Line: 7, Column: 5}
	j := &Join{Pos: pos}
	if got := j.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestChoice_GetPosition(t *testing.T) {
	pos := Position{Line: 8, Column: 5}
	c := &Choice{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestStateNote_GetPosition(t *testing.T) {
	pos := Position{Line: 9, Column: 5}
	n := &StateNote{Pos: pos}
	if got := n.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestStateComment_GetPosition(t *testing.T) {
	pos := Position{Line: 10, Column: 5}
	c := &StateComment{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

// Compile-time interface compliance checks
var (
	_ StateStmt = (*State)(nil)
	_ StateStmt = (*Transition)(nil)
	_ StateStmt = (*StartState)(nil)
	_ StateStmt = (*EndState)(nil)
	_ StateStmt = (*Fork)(nil)
	_ StateStmt = (*Join)(nil)
	_ StateStmt = (*Choice)(nil)
	_ StateStmt = (*StateNote)(nil)
	_ StateStmt = (*StateComment)(nil)
)
