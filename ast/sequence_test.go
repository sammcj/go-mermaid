package ast

import (
	"testing"
)

func TestSequenceDiagram_GetType(t *testing.T) {
	sd := &SequenceDiagram{Type: "sequence"}
	if got := sd.GetType(); got != "sequence" {
		t.Errorf("GetType() = %v, want %v", got, "sequence")
	}
}

func TestSequenceDiagram_GetPosition(t *testing.T) {
	pos := Position{Line: 1, Column: 1}
	sd := &SequenceDiagram{Pos: pos}
	if got := sd.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestSequenceDiagram_GetSource(t *testing.T) {
	source := "sequenceDiagram\n    Alice->>Bob: Hello"
	sd := &SequenceDiagram{Source: source}
	if got := sd.GetSource(); got != source {
		t.Errorf("GetSource() = %v, want %v", got, source)
	}
}

func TestParticipant_GetPosition(t *testing.T) {
	pos := Position{Line: 2, Column: 5}
	p := &Participant{Pos: pos}
	if got := p.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestMessage_GetPosition(t *testing.T) {
	pos := Position{Line: 3, Column: 5}
	m := &Message{Pos: pos}
	if got := m.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestActivation_GetPosition(t *testing.T) {
	pos := Position{Line: 4, Column: 5}
	a := &Activation{Pos: pos}
	if got := a.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestLoop_GetPosition(t *testing.T) {
	pos := Position{Line: 5, Column: 5}
	l := &Loop{Pos: pos}
	if got := l.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestAlt_GetPosition(t *testing.T) {
	pos := Position{Line: 6, Column: 5}
	a := &Alt{Pos: pos}
	if got := a.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestOpt_GetPosition(t *testing.T) {
	pos := Position{Line: 7, Column: 5}
	o := &Opt{Pos: pos}
	if got := o.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestPar_GetPosition(t *testing.T) {
	pos := Position{Line: 8, Column: 5}
	p := &Par{Pos: pos}
	if got := p.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestCritical_GetPosition(t *testing.T) {
	pos := Position{Line: 9, Column: 5}
	c := &Critical{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestBreak_GetPosition(t *testing.T) {
	pos := Position{Line: 10, Column: 5}
	b := &Break{Pos: pos}
	if got := b.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestNote_GetPosition(t *testing.T) {
	pos := Position{Line: 11, Column: 5}
	n := &Note{Pos: pos}
	if got := n.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestBox_GetPosition(t *testing.T) {
	pos := Position{Line: 12, Column: 5}
	b := &Box{Pos: pos}
	if got := b.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestAutonumber_GetPosition(t *testing.T) {
	pos := Position{Line: 13, Column: 5}
	a := &Autonumber{Pos: pos}
	if got := a.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

func TestSeqComment_GetPosition(t *testing.T) {
	pos := Position{Line: 14, Column: 5}
	c := &SeqComment{Pos: pos}
	if got := c.GetPosition(); got != pos {
		t.Errorf("GetPosition() = %v, want %v", got, pos)
	}
}

// Compile-time interface compliance checks
var (
	_ SeqStmt = (*Participant)(nil)
	_ SeqStmt = (*Message)(nil)
	_ SeqStmt = (*Activation)(nil)
	_ SeqStmt = (*Loop)(nil)
	_ SeqStmt = (*Alt)(nil)
	_ SeqStmt = (*Opt)(nil)
	_ SeqStmt = (*Par)(nil)
	_ SeqStmt = (*Critical)(nil)
	_ SeqStmt = (*Break)(nil)
	_ SeqStmt = (*Note)(nil)
	_ SeqStmt = (*Box)(nil)
	_ SeqStmt = (*Autonumber)(nil)
	_ SeqStmt = (*SeqComment)(nil)
)

func TestSequenceDiagram_Fields(t *testing.T) {
	sd := &SequenceDiagram{
		Type:   "sequence",
		Source: "sequenceDiagram\n    Alice->>Bob: Hello",
		Statements: []SeqStmt{
			&Participant{ID: "Alice", Pos: Position{Line: 2, Column: 1}},
		},
		Pos: Position{Line: 1, Column: 1},
	}

	if sd.Type != "sequence" {
		t.Errorf("Type = %v, want %v", sd.Type, "sequence")
	}
	if len(sd.Statements) != 1 {
		t.Errorf("len(Statements) = %d, want 1", len(sd.Statements))
	}
}
