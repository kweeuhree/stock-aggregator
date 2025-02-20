package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func getUserInput() string {
	printUserPrompt()

	stock := readUserInput(os.Stdin)
	if stock == "q" {
		fmt.Println("Bye")
		os.Exit(1)
	}

	if stock == "" {
		fmt.Println("no stock ticker acquired. exiting.")
		os.Exit(1)
	}

	return stock
}

func printUserPrompt() {
	fmt.Println("--------------------------------------")
	fmt.Println("Welcome to the Stock Price Aggregator!")
	fmt.Println("--------------------------------------")
	fmt.Println("This tool fetches real-time stock data from multiple sources (MarketStack, FMPCloud, and FinHub)")
	fmt.Println("and calculates the average open, close, high, and low prices for your chosen ticker.")
	fmt.Println("--------------------------------------")
	fmt.Println("Enter a stock ticker (e.g., AAPL, MSFT, GOOGL). Enter q to quit.")
	fmt.Print("-> ")
}

func readUserInput(in io.Reader) string {
	scanner := bufio.NewScanner(in)

	// Check to see if the user wants to quit
	if strings.EqualFold(scanner.Text(), "q") {
		return ""
	} else {
		if scanner.Scan() {
			return strings.TrimSpace(scanner.Text())
		}
	}

	return ""
}
