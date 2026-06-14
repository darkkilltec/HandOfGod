package ui

// Diese Datei liest die Eingaben des Spielers von der Tastatur (stdin) und
// prüft sie auf Gültigkeit.
//
// WAS HIER ENTSTEHEN SOLL (Skizze)
//
//   // ReadChoice fragt den Spieler nach Aktion + Zielgruppe und gibt die
//   // gewählten Indizes zurück – oder einen Fehler bei ungültiger Eingabe.
//   func ReadChoice(maxAction, maxGroup int) (actionIdx, groupIdx int, err error) { ... }
//
// ROBUSTHEIT: ungültige Eingaben (Buchstabe statt Zahl, Zahl zu groß) sollen
// das Spiel NICHT abstürzen lassen, sondern eine freundliche Meldung erzeugen
// und erneut fragen.
//
// GO-KONZEPTE HIER
//   - bufio.Scanner zum zeilenweisen Lesen von os.Stdin
//   - strconv.Atoi: Text -> Zahl, inklusive Fehlerbehandlung
//   - Gos Fehler-Idiom: Funktionen geben (Wert, error) zurück; der Aufrufer
//     prüft `if err != nil { ... }`
//   - Eingabe-Validierung in einer Schleife
