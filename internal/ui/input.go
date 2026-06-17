package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
func ReadInt(prompt string, min, max int) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		n, err := strconv.Atoi(line)
		if err != nil || n < min || n > max {
			fmt.Printf("Bitte eine Zahl zwischen %d und %d eingeben: ", min, max)
			continue
		}
		return n
	}
}