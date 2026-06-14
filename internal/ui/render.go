// Package ui kümmert sich um die Darstellung im Terminal und um das Einlesen
// der Spielereingaben. Es kennt die Domäne (domain, action), aber NICHT die
// Nebenläufigkeit – die UI rechnet nichts aus, sie zeigt nur an.
package ui

// Diese Datei zeichnet die Ausgabe: das ASCII-Balkendiagramm der Wünsche und
// das Aktionsmenü.
//
// WAS HIER ENTSTEHEN SOLL (Skizze)
//
//   // RenderChart zeichnet pro Gruppe einen Balken, dessen Länge der
//   // Dringlichkeit (Urgency) ihres Wunsches entspricht. Beispiel:
//   //
//   //   Talbewohner   Regen  |##############        | 62
//   //   Bergvolk      Krieg  |#####                 | 21
//   //
//   func RenderChart(groups []domain.Group) { ... }
//
//   // RenderMenu listet die wählbaren Aktionen und Gruppen nummeriert auf,
//   // damit der Spieler per Zahl auswählen kann.
//   func RenderMenu(actions []action.Action, groups []domain.Group) { ... }
//
// TIPP: Balkenlänge = Urgency auf eine feste Breite skalieren (z.B. 0..100 ->
// 0..40 Zeichen). strings.Repeat("#", n) baut den Balken.
//
// GO-KONZEPTE HIER
//   - fmt.Printf mit Formatierung (feste Spaltenbreite: %-12s, %3d ...)
//   - strings.Repeat zum Bauen der Balken
//   - saubere Trennung von Darstellung und Logik
