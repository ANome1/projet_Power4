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

// WinCond vérifie si un joueur a gagné
// Retourne la couleur du gagnant ("red", "yellow") ou "" si pas de gagnant
func (g *Game) WinCond() string {
	rows := len(g.Grid)
	cols := len(g.Grid[0])

	// Vérifier les lignes horizontales
	for row := 0; row < rows; row++ {
		for col := 0; col <= cols-4; col++ {
			if g.Grid[row][col] != "" &&
				g.Grid[row][col] == g.Grid[row][col+1] &&
				g.Grid[row][col] == g.Grid[row][col+2] &&
				g.Grid[row][col] == g.Grid[row][col+3] {
				return g.Grid[row][col]
			}
		}
	}

	// Vérifier les colonnes verticales
	for col := 0; col < cols; col++ {
		for row := 0; row <= rows-4; row++ {
			if g.Grid[row][col] != "" &&
				g.Grid[row][col] == g.Grid[row+1][col] &&
				g.Grid[row][col] == g.Grid[row+2][col] &&
				g.Grid[row][col] == g.Grid[row+3][col] {
				return g.Grid[row][col]
			}
		}
	}

	// Vérifier les diagonales descendantes (\)
	for row := 0; row <= rows-4; row++ {
		for col := 0; col <= cols-4; col++ {
			if g.Grid[row][col] != "" &&
				g.Grid[row][col] == g.Grid[row+1][col+1] &&
				g.Grid[row][col] == g.Grid[row+2][col+2] &&
				g.Grid[row][col] == g.Grid[row+3][col+3] {
				return g.Grid[row][col]
			}
		}
	}

	// Vérifier les diagonales montantes (/)
	for row := 3; row < rows; row++ {
		for col := 0; col <= cols-4; col++ {
			if g.Grid[row][col] != "" &&
				g.Grid[row][col] == g.Grid[row-1][col+1] &&
				g.Grid[row][col] == g.Grid[row-2][col+2] &&
				g.Grid[row][col] == g.Grid[row-3][col+3] {
				return g.Grid[row][col]
			}
		}
	}

	return ""
}
