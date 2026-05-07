package db

import "game-db/game"

// RecommendGames erwartet einen Spielernamen und generiert Spiele-Empfehlungen.
//
// Die Funktion sucht nach Spielen, die der Spieler oft gespielt hat.
// Für diese Spiele werden Spiele mit gleichem Genre gesucht,
// die von anderen Spielern häufig gespielt wurden.
func (db *GameDb) RecommendGames(playerName string) []*game.Game {
	// Hinweis: Implementieren Sie die Empfehlungslogik in mehreren Schritten:
	// 1. Finde den Spieler in der Datenbank.
	// 2. Sammle alle Spiele, die der Spieler oft gespielt hat (mindestens `MinHoursForRecommendation` Stunden).
	// 3. Bestimme die Genres der gespielten Spiele.
	// 4. Suche nach Spielen mit gleichem Genre in der Spiele-Datenbank.
	// 5. Baue das Ergebnis auf:
	//    - Zähle, wie oft diese Spiele von anderen Spielern gespielt wurden.
	//    - Füge Spiele hinzu, die von mindestens `MinPlayersForRecommendation`
	//      anderen Spielern für jeweils mindestens `MinHoursForRecommendation`
	//      gespielt wurden.

	recommendedGames := []*game.Game{}
	// TODO
	return recommendedGames
}
