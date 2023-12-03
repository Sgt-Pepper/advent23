package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color string

type game map[string]int

var games []game

func ReadFile() {
	games = make([]game, 101)
	readFile, err := os.Open("input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		parseGame(fileScanner.Text())
	}
	fmt.Println(games)

}

func parseGame(line string) {

	parts := strings.Split(line, ":")
	start := strings.Fields(parts[0])
	gameNr, _ := strconv.Atoi(start[1])
	games[gameNr] = make(game)

	// evaluate draws
	draws := strings.Split(parts[1], ";")
	for _, d := range draws {
		evalDraw(gameNr, d)
	}

}
func evalDraw(nr int, draw string) {
	balls := strings.Split(draw, ",")
	for _, b := range balls {
		h := strings.Fields(b)
		balls, _ := strconv.Atoi(h[0])
		updateMap(nr, h[1], balls)

	}

}

func updateMap(nr int, color string, balls int) {
	old, ok := games[nr][color]
	// If the key exists
	if !ok || balls > old {
		games[nr][color] = balls

	}

}
