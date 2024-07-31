package pkg

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Call() {
	scanner := bufio.NewScanner(os.Stdin)
	command := os.Args[1]
	var args []string
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./myXargs <command>")
		os.Exit(1)
	}
	for scanner.Scan() {
		input := scanner.Text()
		args = append(args, input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
	cmd := exec.Command(command, args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "executing command:", err)
		os.Exit(1)
	}
}
