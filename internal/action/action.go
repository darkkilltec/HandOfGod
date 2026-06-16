// Package action definiert die Aktionen, die du als Gott ausführen kannst, und
// wie sie sich auf eine Gruppe auswirken.
package action

import (
	"handofgod/internal/domain"
)

// Diese Datei beschreibt das Aktions-System. Eine Aktion verändert den Zustand
// einer Gruppe und beeinflusst damit, wie ihr nächster Wunsch ausfällt.
//
// WAS HIER ENTSTEHEN SOLL
//
//	type Action struct {
//	    Name   string                  // "Regen schicken", "Seuche", ...
//	    Effect func(g *domain.Group)   // wie wirkt die Aktion auf die Gruppe?
//	}
//
//	// Catalog liefert den festen Katalog aller verfügbaren Aktionen:
//	func Catalog() []Action { ... }
//
//	// Apply wendet eine Aktion auf eine Gruppe an:
//	func Apply(a Action, g *domain.Group) { ... }
//
// DESIGNFRAGE FÜRS TUTORIAL
//
//	Aktionen als Funktionswerte (func im struct, wie oben) ODER als Interface
//	mit einer Methode Apply(*domain.Group)? Beides ist in Go üblich – wir
//	vergleichen die zwei Varianten und ihre Vor-/Nachteile.
//
// GO-KONZEPTE HIER
//   - Funktionen als Werte (first-class functions, Closures)
//   - Slices ([]Action) als Liste
//   - Zeiger (*domain.Group), damit die Aktion das ORIGINAL verändert
//   - Import eines anderen internen Pakets: handofgod/internal/domain
type Action struct {
	Name   string
	Effect func(g *domain.Group)
}

func Catalog() []Action {
	return []Action{
		{ Name: "Regen schicken", Effect: func(g *domain.Group) { 
			g.CurrentWish.Urgency = 0 
			g.Faith += 10 
		}},
		{ Name: "Segen", Effect: func(g *domain.Group){
			g.Faith += 15
		}},
		{ Name: "Seuche", Effect: func(g *domain.Group){
			g.Faith -= 20
		}},
	}
}

func Apply(a Action, g *domain.Group) {
	a.Effect(g)
}
