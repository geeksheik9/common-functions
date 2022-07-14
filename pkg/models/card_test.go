package models

import "testing"

func Test_IsValid_Monster_Valid(t *testing.T) {
	card := MonsterCard{MonsterType: "Normal"}
	if err := card.MonsterType.IsValid(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func Test_IsValid_Monster_InValid(t *testing.T) {
	card := MonsterCard{MonsterType: "Invalid"}
	if err := card.MonsterType.IsValid(); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func Test_IsValid_Spell_Valid(t *testing.T) {
	card := SpellCard{SpellType: "Normal"}
	if err := card.SpellType.IsValid(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func Test_IsValid_Spell_Invalid(t *testing.T) {
	card := SpellCard{SpellType: "Invalid"}
	if err := card.SpellType.IsValid(); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

func Test_IsValid_Trap_Valid(t *testing.T) {
	card := TrapCard{TrapType: "Normal"}
	if err := card.TrapType.IsValid(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func Test_IsValid_Trap_Invalid(t *testing.T) {
	card := TrapCard{TrapType: "Invalid"}
	if err := card.TrapType.IsValid(); err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}
