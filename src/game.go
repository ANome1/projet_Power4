package power4

type Players struct {
	Player1       string
	Player2       string
	Player1_Score int
	Player2_Score int
	Difficulty    string
}

type Game struct {
	Turn      string
	Grid      [][]string
	Players   Players
	TurnCount int
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
	case "gravity":
		row = 6
		col = 7
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

// place un jeton dans la colonne spécifiée
func (g *Game) PlaceToken(col int, color string) bool {
	// Vérifier que la colonne est valide
	if col < 0 || col >= len(g.Grid[0]) {
		return false
	}

	// la première case vide en partant du bas
	for row := len(g.Grid) - 1; row >= 0; row-- {
		if g.Grid[row][col] == "" {
			g.Grid[row][col] = color
			return true
		}
	}

	return false
}

func (g *Game) GetCurrentPlayerColor() string {
	if g.Turn == g.Players.Player1 {
		return "red"
	}
	return "yellow"
}

func (g *Game) SwitchTurn() {
	if g.Turn == g.Players.Player1 {
		g.TurnCount++
		g.Turn = g.Players.Player2
	} else {
		g.Turn = g.Players.Player1
	}
}

func (g *Game) WinCond() string {
	rows := len(g.Grid)
	cols := len(g.Grid[0])

	// Vérifier les lignes
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

	// Vérifier les colonnes
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

	// Vérifier les diagonales descendantes
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

	// Vérifier les diagonales montantes
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
func (g *Game) ReverseGravity() {
	// Example: reverse gravity every 5 turns (assuming you have a turn counter)
	// You need to add a TurnCount field to Game struct if you want this logic.
	// For now, just leave this function empty or implement your own logic.
	if g.TurnCount%5 == 0 {
		for col := 0; col < len(g.Grid[0]); col++ {
			// Collect all tokens in the column
			tokens := []string{}
			for row := 0; row < len(g.Grid); row++ {
				if g.Grid[row][col] != "" {
					tokens = append(tokens, g.Grid[row][col])
					g.Grid[row][col] = ""
				}
			}
			// Place tokens at the top of the column
			for i := 0; i < len(tokens); i++ {
				g.Grid[i][col] = tokens[i]
			}
		}
	}
}
