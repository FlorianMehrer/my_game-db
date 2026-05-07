package db

import (
	"game-db/player"
)

// GetPlayer sucht einen Spieler in der Datenbank anhand eines Namens.
// Liefert einen Zeiger auf den Spieler zurück, wenn er gefunden wird,
// oder nil, wenn er nicht gefunden wird.
func (db *GameDb) GetPlayer(name string) *player.Player {
	// Hinweis: Durchsuchen Sie die Liste der Spieler in der Datenbank und
	// vergleichen Sie die Namen der Spieler mit dem gesuchten Namen.
	// Wenn ein Spieler gefunden wird, liefern Sie einen Zeiger auf diesen Spieler zurück.

	// TODO
	return nil
}

// GetPlayersByGame sucht Spieler in der Datenbank, die ein bestimmtes Spiel gespielt haben.
// Erwartet den Titel des Spiels und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayersByGame(title string, min_played int) []*player.Player {
	// Hinweis: Durchsuchen Sie die Liste der Spieler in der Datenbank und überprüfen Sie für jeden Spieler,
	// ob er das gesuchte Spiel mit mindestens der angegebenen Anzahl gespielter Stunden gespielt hat.
	// Sammeln Sie die passenden Spieler in einer Liste, die Sie zurückgeben.

	players := []*player.Player{}
	// TODO
	return players
}
