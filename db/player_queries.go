package db

import (
	"game-db/game"
	"game-db/player"
)

// GetPlayer sucht einen Spieler in der Datenbank anhand eines Namens.
// Liefert einen Zeiger auf den Spieler zurück, wenn er gefunden wird,
// oder nil, wenn er nicht gefunden wird.
func (db *GameDb) GetPlayer(name string) *player.Player {
	// Hinweis: Durchsuchen Sie die Liste der Spieler in der Datenbank und
	// vergleichen Sie die Namen der Spieler mit dem gesuchten Namen.
	// Wenn ein Spieler gefunden wird, liefern Sie einen Zeiger auf diesen Spieler zurück.

	for _, p := range db.Players {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// GetPlayersByGame sucht Spieler in der Datenbank, die ein bestimmtes Spiel gespielt haben.
// Erwartet den Titel des Spiels und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayersByGame(title string, min_played int) []*player.Player {
	// Hinweis: Durchsuchen Sie die Liste der Spieler in der Datenbank und überprüfen Sie für jeden Spieler,
	// ob er das gesuchte Spiel mit mindestens der angegebenen Anzahl gespielter Stunden gespielt hat.
	// Sammeln Sie die passenden Spieler in einer Liste, die Sie zurückgeben.

	players := []*player.Player{}

	for _, p := range db.Players {
		if p.HasPlayedMore(&game.Game{Title: title}, min_played) {
			players = append(players, p)
		}
	}
	return players
}
