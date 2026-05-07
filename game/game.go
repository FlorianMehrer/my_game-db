package game

// Repräsentiert ein Spiel in einer Spiele-Datenbank.
type Game struct {
	Title string // Der Titel des Spiels
	Genre string // Das Genre des Spiels
}

// New erstellt ein neues Spiel mit dem gegebenen Titel.
func New(title string, genre string) *Game {
	// Hinweis: Erstellen Sie ein neues Game-Objekt mit dem übergebenen Titel und Genre.
	//          Liefern Sie einen Pointer auf das neue Game-Objekt zurück.

	// TODO
	return nil
}

// HasGenre prüft, ob das Spiel ein bestimmtes Genre hat.
func (g *Game) HasGenre(genre string) bool {
	// Hinweis: Vergleichen Sie den Genre-String des Spiels mit dem übergebenen Genre-String.

	// TODO
	return false
}
