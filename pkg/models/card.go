package models

import "errors"

// Card is the basic implementation of a card
// swagger:model Card
type Card struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
	ArtworkPath string `json:"artworkPath"`
}

// MonsterType is the type of a monster, extends the type of a card
// swagger:model MonsterType
type MonsterCard struct {
	Card
	Attack      int         `json:"attack"`
	Defense     int         `json:"defense"`
	MonsterType MonsterType `json:"monsterType"`
}

type MonsterType string

const (
	EFFECT_TYPE         MonsterType = "Effect"
	NORMAL_MONSTER_TYPE MonsterType = "Normal"
	FUSION_TYPE         MonsterType = "Fusion"
	XYZ_TYPE            MonsterType = "XYZ"
	SYNCHRO_TYPE        MonsterType = "Synchro"
	RITUAL_MONSTER_TYPE MonsterType = "Ritual"
)

func (mt MonsterType) IsValid() error {
	switch mt {
	case EFFECT_TYPE, NORMAL_MONSTER_TYPE, FUSION_TYPE, XYZ_TYPE, SYNCHRO_TYPE, RITUAL_MONSTER_TYPE:
		return nil
	}
	return errors.New("bad input: Invalid Monster Type")
}

type SpellCard struct {
	Card
	SpellType SpellType `json:"magicType"`
}

type SpellType string

const (
	FIELD_TYPE        SpellType = "Field"
	NORMAL_SPELL_TYPE SpellType = "Normal"
	QUICK_TYPE        SpellType = "Quick"
	CONTINUOUS_TYPE   SpellType = "Continuous"
	EQUIP_TYPE        SpellType = "Equip"
	RITUAL_SPELL_TYPE SpellType = "Ritual"
)

func (st SpellType) IsValid() error {
	switch st {
	case FIELD_TYPE, NORMAL_SPELL_TYPE, QUICK_TYPE, CONTINUOUS_TYPE, EQUIP_TYPE, RITUAL_SPELL_TYPE:
		return nil
	}
	return errors.New("bad input: Invalid Spell Type")
}

type TrapCard struct {
	Card
	TrapType TrapType `json:"trapType"`
}

type TrapType string

const (
	NORMAL_TRAP_TYPE     TrapType = "Normal"
	CONTINUOUS_TRAP_TYPE TrapType = "Continuous"
	COUNTER_TRAP_TYPE    TrapType = "Counter"
)

func (tt TrapType) IsValid() error {
	switch tt {
	case NORMAL_TRAP_TYPE, CONTINUOUS_TRAP_TYPE, COUNTER_TRAP_TYPE:
		return nil
	}
	return errors.New("bad input: Invalid Trap Type")
}
