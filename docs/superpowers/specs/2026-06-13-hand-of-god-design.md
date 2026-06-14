# Hand of God — Design & Tutorial-Fahrplan

**Datum:** 2026-06-13
**Zweck:** Erstes Go-Projekt zum Vertiefen der Sprache. Lernfokus: Goroutines,
Channels, `select`, Pakete/Sichtbarkeit, structs/Methoden, Fehlerbehandlung.

## Spielkonzept

Der Spieler ist **Gott** über mehrere Menschengruppen. Jede Gruppe hat pro
Runde **genau einen dominanten Wunsch** mit einer Dringlichkeit (0–100). Der
Spieler sieht ein **ASCII-Balkendiagramm** der Dringlichkeiten, wählt eine
**Aktion** und eine **Zielgruppe**; danach ändert sich die Welt und eine neue
Runde beginnt.

## Architekturentscheidungen

- **Layout A — mehrere kleine Pakete** unter `internal/` (Domäne, Aktion,
  Simulation, UI). Lehrt Paketgrenzen, Imports und exported/unexported.
- **Goroutines: echte Nebenläufigkeit pro Gruppe.** Jede Gruppe läuft als
  eigene Goroutine und kommuniziert ausschließlich über Channels mit der Welt.
- **Runden-Modell: Tick-gesteuert, pausiert bei Eingabe.** Eine Tick-Uhr
  treibt die Welt; ist der Spieler am Zug, pausiert die Welt (Signal über
  Channel), nach der Aktion läuft sie weiter (Resume).
- **Wunsch-Modell: ein dominanter Wunsch pro Runde** je Gruppe, mit Dringlichkeit.

## Komponenten & Verantwortlichkeiten

| Paket / Datei | Verantwortung | Hängt ab von |
|---|---|---|
| `domain/group.go`, `wish.go` | reine Datentypen (Group, Wish, WishKind) | — |
| `action/action.go` | Aktionskatalog + Wirkung auf eine Gruppe | domain |
| `sim/messages.go` | Channel-Nachrichtentypen (WishUpdate, Command) | domain |
| `sim/group_runtime.go` | Goroutine: Leben einer Gruppe (Wunsch pro Tick) | domain, sim |
| `sim/world.go` | Orchestrator: Ticks, Pause/Resume, Fan-in, Snapshot | domain, sim |
| `ui/render.go` | ASCII-Balkendiagramm + Aktionsmenü | domain, action |
| `ui/input.go` | Spielereingabe lesen & validieren | — |
| `main.go` | alles verdrahten + Game-Loop | alle |

## Datenfluss

1. `main` legt Start-Gruppen an und startet `sim.World`.
2. `World.Run` startet pro Gruppe eine `runGroup`-Goroutine und eine Tick-Uhr.
3. Jeder Tick: Welt schickt `CmdTick` an die Gruppen; jede Gruppe würfelt ihren
   neuen Wunsch und sendet eine `WishUpdate` zurück (Fan-in in einen Channel).
4. Nach N Ticks signalisiert die Welt „Spieler am Zug" und pausiert.
5. `main` holt `World.Snapshot()`, zeigt Chart + Menü, liest die Wahl,
   wendet die `action.Action` auf die Gruppe an, ruft `World.Resume()`.
6. Schleife wiederholt sich (oder endet bei „quit").

## Wichtige Concurrency-Regel

Kein direkter, ungeschützter Zugriff der UI auf laufende Goroutine-Daten. Die
UI bekommt nur **Snapshots** (Kopien). Leitsatz: *„Don't communicate by sharing
memory; share memory by communicating."* Verifikation mit `go run -race .`.

## Tutorial-Fahrplan (Reihenfolge der Implementierung)

1. **domain** — `Group`, `Wish`, `WishKind` (+ `String()`-Methoden). Erstes
   `go build` einer Datei ohne Nebenläufigkeit.
2. **ui/render** — Balkendiagramm aus einer fest verdrahteten Beispielliste.
   Sofortiges visuelles Feedback, noch ohne Logik.
3. **action** — Aktionskatalog + `Apply`. Funktionswerte vs. Interface.
4. **sim/messages + group_runtime** — erste Goroutine, erste Channels.
5. **sim/world** — Tick-Uhr, Fan-in, Pause/Resume, Snapshot.
6. **ui/input** — robustes Einlesen mit Fehlerbehandlung.
7. **main** — alles verdrahten, Game-Loop, mit `-race` testen.

## Bewusst weggelassen (YAGNI)

Speichern/Laden, mehrere Bedürfnis-Dimensionen, Hintergrund-Welt-Events,
Netzwerk, GUI. Können später als Erweiterungsübungen dazukommen.
