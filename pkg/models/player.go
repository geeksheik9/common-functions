package models

type Player struct {
	User
	Decks []Deck `json:"decks"`
}
