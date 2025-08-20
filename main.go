package main

import (
	"blog-aggregator/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	s := &state{
		config: &cfg,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	input := os.Args
	if len(input) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	cmd := command{
		name: input[1],
		args: input[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
