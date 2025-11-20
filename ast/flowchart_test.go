package ast

import (
	"testing"
)

func TestFlowchart_GetType(t *testing.T) {
	tests := []struct {
		name string
		typ  string
	}{
		{"flowchart", "flowchart"},
		{"graph", "graph"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := &Flowchart{Type: tt.typ}
			if got := fc.GetType(); got != tt.typ {
				t.Errorf("GetType() = %v, want %v", got, tt.typ)
			}
		})
	}
}

func TestFlowchart_GetPosition(t *testing.T) {
	pos := Position{Line: 1, Column: 1}
	fc := &Flowchart{Pos: pos}
	if got := fc.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestFlowchart_GetSource(t *testing.T) {
	source := "flowchart TD\n    A-->B"
	fc := &Flowchart{Source: source}
	if got := fc.GetSource(); got != source {
		t.Errorf("GetSource() = %v, want %v", got, source)
	}
}

func TestNodeDef_GetPosition(t *testing.T) {
	pos := Position{Line: 2, Column: 5}
	n := &NodeDef{Pos: pos}
	if got := n.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestLink_GetPosition(t *testing.T) {
	pos := Position{Line: 3, Column: 5}
	l := &Link{Pos: pos}
	if got := l.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestSubgraph_GetPosition(t *testing.T) {
	pos := Position{Line: 4, Column: 5}
	s := &Subgraph{Pos: pos}
	if got := s.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestClassDef_GetPosition(t *testing.T) {
	pos := Position{Line: 5, Column: 5}
	c := &ClassDef{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestClassAssignment_GetPosition(t *testing.T) {
	pos := Position{Line: 6, Column: 5}
	ca := &ClassAssignment{Pos: pos}
	if got := ca.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestComment_GetPosition(t *testing.T) {
	pos := Position{Line: 7, Column: 5}
	c := &Comment{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

// Compile-time interface compliance checks
var (
	_ Statement = (*NodeDef)(nil)
	_ Statement = (*Link)(nil)
	_ Statement = (*Subgraph)(nil)
	_ Statement = (*ClassDef)(nil)
	_ Statement = (*ClassAssignment)(nil)
	_ Statement = (*Comment)(nil)
)
