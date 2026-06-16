package domain

import "math/rand"

// Diese Datei beschreibt, WAS sich eine Gruppe wünscht (Wish) und wie DRINGEND
// (Urgency). Beide Typen gehören eng zur Group und liegen daher im selben
// Paket – ein Paket darf aus mehreren Dateien bestehen.
//
// WAS HIER ENTSTEHEN SOLL
//
//   type WishKind int   // Art des Wunsches als Aufzählung (enum)
//
//   const (
//       WishRain    WishKind = iota // iota zählt automatisch hoch: 0, 1, 2, ...
//       WishWar                     // 1
//       WishTemple                  // 2
//       WishHarvest                 // 3
//       // ... beliebig erweiterbar
//   )
//
//   type Wish struct {
//       Kind    WishKind // welcher Wunsch
//       Urgency int      // Dringlichkeit 0..100 -> Höhe des ASCII-Balkens
//   }
//
// IDEE: eine Methode, die WishKind als Text liefert (für die Anzeige):
//   func (k WishKind) String() string { ... }   // "Regen", "Krieg", "Tempel"...
//
// GO-KONZEPTE HIER
//   - benannte Konstanten + iota für saubere Aufzählungen
//   - eigener Integer-Typ (type WishKind int) für mehr Typsicherheit
//   - das fmt.Stringer-Interface für lesbare Ausgabe von eigenen Typen
type WishKind int

const (
	WishRain WishKind = iota
	WishWar
	WishTemple
	WishHarvest
	WishRiches
	wishKindCount //Counter for wishes needs to be at the end always, new entries go before this entry
)

type Wish struct {
	Kind    WishKind
	Urgency int
}

func (k WishKind) String() string {
	return []string{"Regen", "Krieg", "Tempel", "Ernte", "Reichtum"}[k]
}

// helper function to get a random wish with the maximum number of wishes
func RandomWishKind() WishKind {
	return WishKind(rand.Intn(int(wishKindCount)))
}