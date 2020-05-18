package main

import (
	"./packages/ballclock"
	"fmt"
	"os"
	"strconv"
	"time"
)

func throwError() error{
	return fmt.Errorf("Wrong input")
}

func main() {
	start := time.Now()

	argumentsLength := len(os.Args)
	if argumentsLength < 2 || argumentsLength > 3 {
		panic(throwError())
	}

	ballcount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(throwError())
	}

	var cycles int
	if argumentsLength == 3 {
		cycles, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(throwError())
		}
	}

	if ballcount < 27 || ballcount > 127 || cycles < 0 {
		panic(throwError())
	}

	if argumentsLength == 3 {
		run := ballclock.RunTime(int8(ballcount),uint16(cycles))
		fmt.Println(run)
	} else {
		run := ballclock.RunAmount(int8(ballcount))
		fmt.Println(run)
	}
	end := time.Since(start)
	fmt.Printf("Completed in %d milliseconds (%1.3f seconds)", end/time.Millisecond, end.Seconds())
}
