package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Sgt-Pepper/advent23/day01"
)

func main() {

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	fmt.Println("the grand total is: ", day01.EvalDocument(day01.Puzzle))
	//fmt.Println("the grand total is: ", day01.EvalDocument(day01.Teststring))

}
