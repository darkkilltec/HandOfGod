package sim

import (
	"handofgod/internal/domain"
)
// Diese Datei enthält das "Leben" einer einzelnen Gruppe: eine Funktion, die
// als eigene Goroutine läuft (genau eine pro Gruppe). Das ist das Herzstück
// deines Concurrency-Lernziels.
//
// WAS HIER ENTSTEHEN SOLL (Skizze)
//
//   // runGroup läuft als Goroutine (Start in world.go via: go runGroup(...)).
//   // Sie wartet auf Befehle der Welt und schickt neue Wünsche zurück.
//   func runGroup(
//       g        *domain.Group,
//       commands <-chan Command,    // NUR-empfangen: Befehle von der Welt
//       updates  chan<- WishUpdate, // NUR-senden:   Wünsche an die Welt
//   ) {
//       for cmd := range commands {     // läuft, bis der Channel geschlossen wird
//           switch cmd {
//           case CmdTick:
//               // neuen dominanten Wunsch bestimmen (z.B. zufällig, abhängig
//               // vom aktuellen Zustand) und per updates an die Welt senden.
//           case CmdStop:
//               return // Goroutine sauber beenden
//           }
//       }
//   }
//
// WIE ÄNDERT SICH EIN WUNSCH? Zum Beispiel zufällig (math/rand) und abhängig
// vom aktuellen Zustand der Gruppe (z.B. niedriger Glaube -> dringlicherer
// Wunsch). Die genaue Regel entwirfst du im Tutorial.
//
// GO-KONZEPTE HIER
//   - Goroutines (nebenläufige Funktionen)
//   - gerichtete Channels (<-chan / chan<-) für klare, sichere Rollen
//   - select / switch, um auf Befehle zu reagieren
//   - eine Goroutine sauber beenden (Channel schließen, statt sie zu "killen")
//   - math/rand für Zufall
func runGroup(g *domain.Group, commands <-chan Command, updates chan<- WishUpdate) {
	for cmd := range commands {
		switch cmd {
		case CmdTick:
			// neuen dominanten Wunsch bestimmen (z.B. zufällig, abhängig
			// vom aktuellen Zustand) und per updates an die Welt senden.
		case CmdStop:
			return // Goroutine sauber beenden
		}
	}
}