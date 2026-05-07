package db

import "game-db/game"

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	var games []*game.Game
	player := db.GetPlayer(name)
	if player == nil {
		return games
	}

	for _, title := range player.PlayedGames(min_played) {
		for _, g := range db.Games {
			if g.Title == title {
				games = append(games, g)
				break
			}
		}
	}
	return games
}

// GetGamesByGenre sucht Spiele in der Datenbank anhand ihres Genres.
func (db *GameDb) GetGamesByGenre(genre string) []*game.Game {
	var games []*game.Game
	for _, g := range db.Games {
		if g.Genre == genre {
			games = append(games, g)
		}
	}
	return games
}
