package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"slices"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("usage: ./clipff <video file> [start time] [end time] [all] [res] <output file>")
	}
	fmt.Println(len(os.Args))

	inputFile := os.Args[1]
	args := []string{"-i", inputFile}
	outputFile := os.Args[len(os.Args)-1]

	for i, arg := range os.Args {
		if i == 0 || i == len(os.Args)-1 {
			continue
		}
		if strings.Contains(arg, ":") && !slices.Contains(args, "-ss") {
			args = append(args, "-ss")
			args = append(args, arg)
		} else if strings.Contains(arg, ":") {
			args = append(args, "-to")
			args = append(args, arg)
		} else if arg == "all" {
			args = append(args, "-map")
			args = append(args, "0:a:3")
			args = append(args, "-map")
			args = append(args, "0:v:0")
		} else if arg == "res" {
			args = append(args, "-vf")
			args = append(args, "scale=-1:1440")
		}
	}
	args = append(args, outputFile)
	fmt.Println(args)

	cmd := exec.Command("ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
