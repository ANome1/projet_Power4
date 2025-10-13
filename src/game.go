package power4

type Players struct {
	Player1       string
	Player2       string
	Player1_Score int
	Player2_Score int
	Difficulty    string
}

type Game struct {
	Turn    string
	Grid    [][]string
	Players Players
}

func NewGame(player *Players) *Game {
	row := 0
	col := 0
	switch player.Difficulty {
	case "easy":
		row = 6
		col = 7
	case "normal":
		row = 6
		col = 9
	case "hard":
		row = 7
		col = 8
	default:
		row = 6
		col = 7
	}

	grid := make([][]string, row)
	for i := range grid {
		grid[i] = make([]string, col)
		for j := range grid[i] {
			grid[i][j] = ""
		}
	}

	return &Game{
		Turn:    player.Player1,
		Grid:    grid,
		Players: *player,
	}
}

// PlaceToken place un jeton dans la colonne spécifiée
func (g *Game) PlaceToken(col int, color string) bool {
	// Vérifier que la colonne est valide
	if col < 0 || col >= len(g.Grid[0]) {
		return false
	}

	// Trouver la première case vide en partant du bas
	for row := len(g.Grid) - 1; row >= 0; row-- {
		if g.Grid[row][col] == "" {
			g.Grid[row][col] = color
			return true
		}
	}

	// Colonne pleine
	return false
}

// GetCurrentPlayerColor retourne la couleur du joueur actuel
func (g *Game) GetCurrentPlayerColor() string {
	if g.Turn == g.Players.Player1 {
		return "red"
	}
	return "yellow"
}

// SwitchTurn change de joueur
func (g *Game) SwitchTurn() {
	if g.Turn == g.Players.Player1 {
		g.Turn = g.Players.Player2
	} else {
		g.Turn = g.Players.Player1
	}
}
