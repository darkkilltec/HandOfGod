// Package sim enthält die Simulation: die laufende Welt und das nebenläufige
// "Leben" der Gruppen über Goroutines und Channels. Hier wohnt dein
// wichtigstes Lernziel – echtes Go-Concurrency.
package sim

// Diese Datei definiert die NACHRICHTEN, die über Channels zwischen der Welt
// und den Gruppen-Goroutines fließen. Eigene, klar benannte Nachrichtentypen
// machen Nebenläufigkeit verständlich und testbar (statt "nackter" Werte).
//
// WAS HIER ENTSTEHEN SOLL (Skizze)
//
//   // Eine Gruppe meldet ihren neuen Wunsch an die Welt:
//   type WishUpdate struct {
//       GroupName string
//       Wish      domain.Wish
//   }
//
//   // Steuersignale von der Welt AN eine Gruppe:
//   type Command int
//   const (
//       CmdTick Command = iota // "berechne deinen nächsten Wunsch"
//       CmdStop                // "beende dich sauber" (Goroutine stoppen)
//   )
//
// GO-KONZEPTE HIER
//   - Channels als typsichere "Röhren" zwischen Goroutines
//   - die Richtung des Datenflusses bewusst entwerfen (wer sendet, wer empfängt)
//   - warum eigene Nachrichtentypen robuster sind als einzelne lose Werte
