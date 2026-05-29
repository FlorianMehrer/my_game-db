package db

import "game-db/game"

// GetPlayedGames sucht alle Spiele, die ein gegebener Spieler gespielt hat.
// Erwartet dabei den Namen des Spielers und die Mindestanzahl gespielter Stunden.
func (db *GameDb) GetPlayedGames(name string, min_played int) []*game.Game {
	games := []*game.Game{}
	for Spieler := range db.Name {
		if Spieler == name {
		gamess  := 	[Spieler]PlayedGames(min_played)
		}
	}  
	for y := range gamess {
		for  range db.Games {
			if y == db.Games.Name {
				games = append(games + db.Games[0])
			}
		}
	} 
	return games
}
	

// GetGamesByGenre sucht Spiele in der Datenbank anhand ihres Genres.
func (db *GameDb) GetGamesByGenre(genre string) []*game.Game {
	games := []*game.Game{}
	// TODO
	return games
}
