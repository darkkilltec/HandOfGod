// Package main ist der Einstiegspunkt des Spiels "Hand of God".
//
// AUFGABE DIESER DATEI
//   Das Spiel "verdrahten" (engl. composition root): Gruppen anlegen, die
//   Welt (sim.World) starten und die Haupt-Spielschleife (Game-Loop) ausführen.
//   Diese Datei enthält bewusst WENIG Logik – sie ruft nur die anderen Pakete
//   auf. So bleibt jeder Teil für sich testbar.
//
// WAS HIER SPÄTER ENTSTEHEN SOLL
//   func main() {
//       1. Eine Liste von Start-Gruppen anlegen ([]*domain.Group).
//       2. Eine sim.World mit diesen Gruppen erzeugen und starten.
//          World.Run() startet pro Gruppe eine eigene Goroutine.
//       3. Game-Loop (for-Schleife):
//            a) Warten auf das Signal "Spieler ist am Zug" (Channel der Welt).
//            b) Aktuellen Zustand als Snapshot holen und mit ui.RenderChart
//               anzeigen (ASCII-Balkendiagramm) sowie ui.RenderMenu (Aktionen).
//            c) Mit ui.ReadChoice die Wahl des Spielers einlesen.
//            d) Gewählte action.Action auf die Zielgruppe anwenden.
//            e) Der Welt "weiter" signalisieren (World.Resume) -> neue Runde.
//          Die Schleife endet z.B. nach N Runden oder bei Eingabe "quit".
//   }
//
// GO-KONZEPTE, DIE DU HIER ÜBST
//   - package main und func main als Programmstart
//   - Pakete importieren (handofgod/internal/...)
//   - Channels, um auf Ereignisse zu warten (<-ch)
//   - die for-Schleife als Game-Loop
package main
import (
	"handofgod/internal/domain"
	"handofgod/internal/ui"
)

func main() {
	groups := []domain.Group{
		{Name: "Talbewohner", Faith: 50, CurrentWish: domain.Wish{Kind: domain.WishRain, Urgency: 62}},
		{Name: "Bergvolk", Faith: 30, CurrentWish: domain.Wish{Kind: domain.WishWar, Urgency: 21}},
		{Name: "Küstenvolk", Faith: 70, CurrentWish: domain.Wish{Kind: domain.WishTemple, Urgency: 88}},
	}	
	ui.RenderChart(groups)
}