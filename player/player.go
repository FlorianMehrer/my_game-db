package player

import "game-db/game"

// Repräsentiert einen Spieler in einer Spiele-Datenbank.
type Player struct {
	Name        string
	playedGames map[string]int
}

// New erstellt einen neuen Spieler mit dem gegebenen Namen und Geburtsjahr.
func New(name string) *Player {
	// Hinweis: Erstellen Sie ein neues Player-Objekt mit dem übergebenen Namen und einem leeren Map für gespielte Spiele.
	//          Liefern Sie einen Pointer auf das neue Player-Objekt zurück.

	// begin:solution
	return &Player{Name: name, playedGames: map[string]int{}}
	// end:solution
}

// PlayGame fügt ein Spiel zu den gespielten Spielen des Spielers hinzu und erhöht die Anzahl der gespielten Stunden.
func (p *Player) PlayGame(g *game.Game, hours int) {
	// Hinweis: Fügen Sie das Spiel mit der Anzahl der gespielten Stunden zur Map `playedGames` des Spielers hinzu.

	// begin:solution
	p.playedGames[g.Title] = hours
	// end:solution
}

// HasPlayed prüft, ob der Spieler ein bestimmtes Spiel gespielt hat.
func (p *Player) HasPlayed(g *game.Game) bool {
	// Hinweis: Überprüfen Sie, ob das Spiel in der Map `playedGames` des Spielers vorhanden ist.

	// begin:solution
	_, played := p.playedGames[g.Title]
	return played
	// end:solution
}

// HasPlayedMore prüft, ob der Spieler ein bestimmtes Spiel mindestens `hours` Stunden gespielt hat.
func (p *Player) HasPlayedMore(g *game.Game, hours int) bool {
	// Hinweis: Überprüfen Sie, ob die gespielten Stunden für das Spiel in der Map `playedGames`
	// des Spielers größer oder gleich `hours` sind.

	// begin:solution
	playedHours := p.playedGames[g.Title]
	return playedHours >= hours
	// end:solution
}

// PlayedGames liefert eine Liste mit den Spielen, die der
// Spieler mehr als `hours` Stunden gespielt hat.
func (p *Player) PlayedGames(hours int) []string {
	// Hinweis: Durchsuchen Sie die Map `playedGames` des Spielers und sammeln Sie die Titel der Spiele,
	// die der Spieler mindestens `hours` Stunden gespielt hat, in einer Liste.

	// begin:solution
	games := []string{}
	for title, playedHours := range p.playedGames {
		if playedHours >= hours {
			games = append(games, title)
		}
	}
	return games
	// end:solution
}
