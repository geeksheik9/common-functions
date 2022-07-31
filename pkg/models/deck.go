package models

type Deck struct {
	ID       int64         `json:"id"`
	Deck     []Card        `json:"deck"`
	Monsters []MonsterCard `json:"monsters"`
	Spells   []SpellCard   `json:"spells"`
	Traps    []TrapCard    `json:"traps"`
}
