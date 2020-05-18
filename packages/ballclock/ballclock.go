package ballclock

import (
	"encoding/json"
	"fmt"
)

type track struct {
	capacity int8
	track    []int8
	next     *track
}

var (
	balls			int8
	time			int8
	mainTrack 	 	track
	minTrack 	 	track
	fiveMinTrack 	track
	hourTrack 	 	track
	initialOrder 	[]int8
)


func initialize(amount int8) {
	balls = amount
	mainTrack = track{balls, []int8{}, &minTrack}
	hourTrack = track{11, []int8{}, &mainTrack}
	fiveMinTrack = track{11, []int8{}, &hourTrack}
	minTrack = track{4, []int8{}, &fiveMinTrack}

	for i := int8(1); i <= balls; i++ {
		mainTrack.track = append(mainTrack.track, i)
	}
	initialOrder = mainTrack.track
}

func RunAmount(ballCount int8) string {
	initialize(ballCount)
	cycle := int32(0)
outer:
	for {
		ball := moveOneBall()
		cycle++
		moveToTrack(ball, &minTrack)
		if int8(len(mainTrack.track)) == balls {
			for x, y := range mainTrack.track {
				if y != initialOrder[x] {
					continue outer
				}
			}
			break outer
		}
	}
	return fmt.Sprintf("%d balls cycle after %d days.", balls, cycle/(60*24))
}

func RunTime(ballCount int8, minutes uint16) string {
	initialize(ballCount)
	for i := uint16(0); i < minutes; i++ {
		ball := moveOneBall()
		moveToTrack(ball, &minTrack)
	}
	return printResult()
}

func moveOneBall() int8 {
	ball := mainTrack.track[0]
	mainTrack.track = mainTrack.track[1:]
	return ball
}

func moveToTrack(ball int8, t *track) {
	if int8(len(t.track)) < t.capacity {
		t.track = append(t.track, ball)
		return
	}
	for i := t.capacity - 1; i >= 0; i-- {
		mainTrack.track = append(mainTrack.track, t.track[i])
	}
	t.track = t.track[:0]
	moveToTrack(ball, t.next)
}

func printResult() string {
	main, _ := json.Marshal(mainTrack.track)
	hour, _ := json.Marshal(hourTrack.track)
	fiveMin, _ := json.Marshal(fiveMinTrack.track)
	min, _ := json.Marshal(minTrack.track)
	return fmt.Sprintf("Main:%s,Min:%s,FiveMin:%s,Hour:%s", main, min, fiveMin, hour)
}
