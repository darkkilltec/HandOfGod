# Hand of God

Ein kleines, rundenbasiertes Terminal-Spiel zum **Go lernen**. Du spielst
**Gott** über mehrere Menschengruppen. Jede Gruppe lebt als eigene Goroutine
und entwickelt pro Runde **einen dominanten Wunsch** mit einer Dringlichkeit.
Eine Tick-Uhr treibt die Welt voran, **pausiert aber, sobald du am Zug bist**.

Du siehst ein **ASCII-Balkendiagramm** der Wünsche, wählst eine **Aktion** und
eine **Zielgruppe**, danach läuft die Zeit weiter und die Wünsche ändern sich.

## Spielablauf (Ziel)

```
Runde 7
========================================
Talbewohner   Regen  |##############        | 62
Bergvolk      Krieg  |#####                 | 21
Küstenvolk    Tempel |####################  | 88

Aktionen:  1) Regen schicken   2) Seuche   3) Segen   4) Beben
Gruppen:   1) Talbewohner      2) Bergvolk 3) Küstenvolk

> Aktion: 1   Gruppe: 3
```

## Projektstruktur

```
HandOfGod/
├── go.mod                          Modul-Definition
├── main.go                         Einstiegspunkt, Game-Loop ("verdrahtet" alles)
└── internal/
    ├── domain/                     reine Datentypen, keine Logik
    │   ├── group.go                Group (Menschengruppe)
    │   └── wish.go                 Wish + Wunsch-Arten (enum via iota)
    ├── action/
    │   └── action.go               Gott-Aktionen und ihre Wirkung
    ├── sim/                        die nebenläufige Simulation
    │   ├── messages.go             Channel-Nachrichtentypen
    │   ├── group_runtime.go        Goroutine: "Leben" einer Gruppe
    │   └── world.go                Orchestrator: Ticks, Pause/Resume, Aggregation
    └── ui/
        ├── render.go               ASCII-Balkendiagramm + Menü
        └── input.go                Spielereingabe lesen & prüfen
```

> Alle `.go`-Dateien sind aktuell **nur mit Kommentaren** gefüllt (kein Code).
> Sie beschreiben, was die Datei tut und welche Go-Konzepte dort geübt werden.
> Wir füllen sie Schritt für Schritt im geführten Tutorial.

## Voraussetzung: Go installieren

Go ist auf diesem Rechner noch nicht installiert. Unter Fedora/Nobara:

```bash
sudo dnf install golang
# danach prüfen:
go version
```

(Alternativ die offizielle Version von https://go.dev/dl/ installieren.)

## Starten (sobald Code drin ist)

```bash
go run .            # Spiel starten
go run -race .      # mit Data-Race-Detektor (wichtig bei Goroutines!)
go test ./...       # Tests ausführen
```
