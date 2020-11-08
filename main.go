package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	s := string(stdin)

	switch os.Args[1] {
	case "=", "==":
		if s == os.Args[2] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	case "!=":
		if s != os.Args[2] {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

	case "-eq", "-ne", "-gt", "-ge", "-lt", "-le":
		num1, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		switch os.Args[1] {
		case "-eq":
			if num1 == num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "-ne":
			if num1 != num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "-gt":
			if num1 > num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "-ge":
			if num1 >= num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "-lt":
			if num1 < num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		case "-le":
			if num1 <= num2 {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
	}
}
