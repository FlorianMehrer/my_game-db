package db

import "game-db/game"

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	// Hinweis: Suchen Sie den Spieler mit dem gegebenen Namen in der Datenbank.
	//          Wenn der Spieler gefunden wird, rufen Sie die Methode `PlayedGames` des Spielers auf,
	//          um die Titel der gespielten Spiele zu erhalten. Vergleichen Sie diese Titel mit den Spielen in der Datenbank
	//          und sammeln Sie die passenden Spiele in einer Liste, die Sie zurückgeben.

	games := []*game.Game{}
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
	// Hinweis: Durchsuchen Sie die Liste der Spiele in der Datenbank und sammeln Sie die Spiele,
	// die das gesuchte Genre haben, in einer Liste, die Sie zurückgeben.

	games := []*game.Game{}
	for _, g := range db.Games {
		if g.Genre == genre {
			games = append(games, g)
		}
	}
	return games
}
