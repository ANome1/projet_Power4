package power4

type Players struct {
	Player1       string
	Player2       string
	Player1_Score int
	Player2_Score int
	Difficulty    string
}

type Game struct {
	Turn string
	Grid [][]string
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
	}
	return &Game{
		Turn: "Player1",
		Grid: grid,
	}
}

func PlacePawn(game *Game, col int, player *Players) bool {
	for i := len(game.Grid) - 1; i >= 0; i-- {
		if game.Grid[i][col] == "" {
			if game.Turn == "Player1" {
				game.Grid[i][col] = "X"
				game.Turn = "Player2"
			} else {
				game.Grid[i][col] = "O"
				game.Turn = "Player1"
			}
			return true
		}
	}
	return false
}
