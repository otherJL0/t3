package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	player  string
	winner  string
	board   [][]string
	counter int
	x       int
	y       int
}

func initialModel() model {
	return model{
		player: "X",
		board:  [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
	}
}

func (m model) Init() tea.Cmd {
	// No I/O at this time
	return nil
}

func (m model) DetermineWinner() string {
	rows := []string{"", "", ""}
	cols := []string{"", "", ""}
	lrDiag := fmt.Sprintf("%s%s%s", m.board[0][0], m.board[1][1], m.board[2][2])
	rlDiag := fmt.Sprintf("%s%s%s", m.board[2][0], m.board[1][1], m.board[2][0])
	for icol, row := range m.board {
		for irow, square := range row {
			rows[irow] += square
			cols[icol] += square
		}
	}

	for i := 0; i < 3; i++ {
		if rows[i] == "XXX" || cols[i] == "XXX" {
			return "X"
		}

		if rows[i] == "OOO" || cols[i] == "OOO" {
			return "O"
		}
	}
	if lrDiag == "XXX" || rlDiag == "XXX" {
		return "X"
	}

	if lrDiag == "OOO" || rlDiag == "OOO" {
		return "O"
	}
	return ""
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Handle key press
		switch msg.String() {

		// Replay game
		case "r":
			if m.winner != "" || m.counter == 9 {
				m = initialModel()
			}

		// Quit with the following keys
		case "ctrl+c", "q":
			return m, tea.Quit

		// Movement Up
		case "up", "k":
			if m.y > 0 {
				m.y--
			}

		// Movement Down
		case "down", "j":
			if m.y < 2 {
				m.y++
			}

		// Movement Left
		case "left", "h":
			if m.x > 0 {
				m.x--
			}

		// Movement Right
		case "right", "l":
			if m.x < 2 {
				m.x++
			}

		case "enter", " ":
			if m.board[m.y][m.x] != " " {
				return m, nil
			}

			if m.winner == "" {
				m.board[m.y][m.x] = m.player
				m.counter++
			}

			m.winner = m.DetermineWinner()

			if m.player == "X" {
				m.player = "O"
			} else {
				m.player = "X"
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Tic-Tac-Toe\n\n"

	for iy, row := range m.board {
		s += "+---+---+---+\n|"
		for ix, col := range row {
			cursor := " "
			if m.x == ix && m.y == iy {
				cursor = "_"
			}
			if col == " " {
				s += fmt.Sprintf("%s%s%s|", cursor, cursor, cursor)
			} else {
				s += fmt.Sprintf("%s%s%s|", cursor, col, cursor)
			}
		}
		s += "\n"
	}
	s += "+---+---+---+\n"
	footer := ""
	if m.winner != "" {
		footer = fmt.Sprintf("Winner is %s! (r) Replay", m.winner)
	}

	if m.counter == 9 {
		footer = "No winner. (r) Replay"
	}
	s += footer
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error! %v", err)
		os.Exit(1)
	}
}
