package sim

import "handofgod/internal/domain"


// Diese Datei ist der ORCHESTRATOR: die "Welt". Sie startet pro Gruppe eine
// Goroutine, treibt die Zeit über eine Tick-Uhr voran, sammelt die Wünsche ein
// und PAUSIERT, sobald der Spieler am Zug ist (Tick-gesteuert mit Pause –
// genau das von dir gewählte Runden-Modell).
//
// WAS HIER ENTSTEHEN SOLL (Skizze)
//
//   type World struct {
//       groups []*domain.Group
//       // pro Gruppe ein commands-Channel; ein gemeinsamer updates-Channel
//       // (Fan-in); eine Tick-Uhr (time.Ticker); ein Channel "Spieler am Zug"
//       // und ein Resume-Signal ...
//   }
//
//   func NewWorld(groups []*domain.Group) *World { ... }
//
//   // Run startet alle Gruppen-Goroutines und die Tick-Schleife. Nach einer
//   // festen Zahl Ticks (= "Runde vorbei") pausiert die Welt und meldet über
//   // einen Channel: "Spieler ist am Zug".
//   func (w *World) Run() { ... }
//
//   // Snapshot liefert eine sichere KOPIE des aktuellen Zustands für die UI,
//   // damit die UI nicht direkt auf laufende Goroutine-Daten zugreift
//   // (Data-Race vermeiden!).
//   func (w *World) Snapshot() []domain.Group { ... }
//
//   // Resume hebt die Pause auf -> die Zeit läuft weiter, neue Runde beginnt.
//   func (w *World) Resume() { ... }
//
// WICHTIGSTES CONCURRENCY-THEMA HIER: keine geteilten Daten ohne Schutz.
// Go-Leitsatz: "Don't communicate by sharing memory; share memory by
// communicating." -> Zustand fließt über Channels, nicht über gemeinsame
// Variablen. Prüfe deinen Code später mit `go run -race .`.
//
// GO-KONZEPTE HIER
//   - time.Ticker als Spiel-Uhr
//   - Fan-in: viele Gruppen-Channels -> ein Welt-Channel
//   - sync.WaitGroup, um sauber auf das Ende aller Goroutines zu warten
//   - Data-Races erkennen und vermeiden (der -race-Detektor)

type World struct {
	groups []*domain.Group
	commands []chan Command
	updates chan WishUpdate
}

func NewWorld(groups []*domain.Group) *World {
	w := &World{
		groups: groups,
		updates: make(chan WishUpdate),
	}
	for range groups {
		w.commands = append(w.commands, make(chan Command))
	}
	return w
}

func (w *World) Start() {
	for i, g := range w.groups {
		go runGroup(g, w.commands[i], w.updates)
	}
}

func (w *World) Tick() {
	for _, c := range w.commands {
		c <- CmdTick
	}
	for range w.groups {
		<-w.updates
	}
}

func (w *World) Stop() {
	for _, c := range w.commands {
		c <- CmdStop
	}
}

func (w *World) Snapshot() []domain.Group {
	var snapshot []domain.Group
	for _, g := range w.groups {
		snapshot = append(snapshot, *g)
	}
	return snapshot
}