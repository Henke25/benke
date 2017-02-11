package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	cleverbot "github.com/ugjka/cleverbot-go"
)

func main() {
	cb, err := cleverbot.New()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Press CTRL-C to exit.")

	// Start chat loop.
	for scanner.Scan() {
		response, err := ask(cb, scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Cleverbot: " + response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}

func ask(cb *cleverbot.Session, input string) (response string, err error) {
	return cb.Ask(input)
}
