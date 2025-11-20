package ast

import (
	"testing"
)

func TestClassDiagram_GetType(t *testing.T) {
	cd := &ClassDiagram{Type: "class"}
	if got := cd.GetType(); got != "class" {
		t.Errorf("GetType() = %v, want %v", got, "class")
	}
}

func TestClassDiagram_GetPosition(t *testing.T) {
	pos := Position{Line: 1, Column: 1}
	cd := &ClassDiagram{Pos: pos}
	if got := cd.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestClassDiagram_GetSource(t *testing.T) {
	source := "classDiagram\n    Class01 <|-- Class02"
	cd := &ClassDiagram{Source: source}
	if got := cd.GetSource(); got != source {
		t.Errorf("GetSource() = %v, want %v", got, source)
	}
}

func TestClass_GetPosition(t *testing.T) {
	pos := Position{Line: 2, Column: 5}
	c := &Class{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestRelationship_GetPosition(t *testing.T) {
	pos := Position{Line: 3, Column: 5}
	r := &Relationship{Pos: pos}
	if got := r.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestClassNote_GetPosition(t *testing.T) {
	pos := Position{Line: 4, Column: 5}
	n := &ClassNote{Pos: pos}
	if got := n.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestClassComment_GetPosition(t *testing.T) {
	pos := Position{Line: 5, Column: 5}
	c := &ClassComment{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

// Compile-time interface compliance checks
var (
	_ ClassStmt = (*Class)(nil)
	_ ClassStmt = (*Relationship)(nil)
	_ ClassStmt = (*ClassNote)(nil)
	_ ClassStmt = (*ClassComment)(nil)
)
