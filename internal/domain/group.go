// Package domain enthält die zentralen Datentypen ("das Vokabular") des Spiels.
//
// Hier liegt KEINE Spiellogik und KEINE Nebenläufigkeit – nur reine Datentypen.
// Dadurch bleibt die Domäne einfach, testbar und unabhängig vom Rest. Andere
// Pakete (action, sim, ui) bauen auf diesen Typen auf.
package domain

// Diese Datei beschreibt eine Menschengruppe (Group).
//
// WAS HIER ENTSTEHEN SOLL
//
//   type Group struct {
//       Name    string // Anzeigename, z.B. "Talbewohner"
//       Faith   int    // Glaube an dich (0..100) – optional fürs Balancing
//       Current Wish   // der aktuelle dominante Wunsch dieser Runde
//       // bei Bedarf weitere Felder: ID, Bevölkerungszahl, Zufriedenheit ...
//   }
//
// IDEEN FÜR METHODEN (kommen später im Tutorial)
//   - func (g Group) String() string  // hübsche Textausgabe (fmt.Stringer)
//
// GO-KONZEPTE HIER
//   - struct: eigene zusammengesetzte Datentypen
//   - exported vs unexported: GROSSbuchstabe am Anfang = von anderen Paketen
//     sichtbar, kleinbuchstabe = nur innerhalb von domain sichtbar
//   - Werte- vs Zeiger-Empfänger bei Methoden (value vs pointer receiver)
type Group struct {
	Name    string
	Faith   int
	CurrentWish Wish
}

func (g *Group) ChangeFaith(delta int) {
	g.Faith += delta
	if g.Faith > 100 {
		g.Faith = 100
	}
	if g.Faith < 0 {
		g.Faith = 0
	}
}