package ballclock

import (
	"testing"
)

type byAmountTest struct {
	ballsAmount   int
	expected string
}
type byTimeTest struct {
	cycles   int
	expected string
}

var timeTests = []byTimeTest{
	{325, `Main:[11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9],Min:[],FiveMin:[22,13,25,3,7],Hour:[6,12,17,4,15]`},
	{300, `Main:[18,26,5,11,22,8,19,30,2,13,20,29,10,24,25,1,28,21,16,3,9,27,14,23,7],Min:[],FiveMin:[],Hour:[6,12,17,4,15]`},
	{400, `Main:[22,13,25,3,8,1,11,5,20,16,21,30,28,10,2,27],Min:[],FiveMin:[18,14,24,7,26,19,29,9],Hour:[6,12,17,4,15,23]`},
}
var amountTests = []byAmountTest{
	{30, `30 balls cycle after 15 days.`},
	{45, `45 balls cycle after 378 days.`},
	{123, `123 balls cycle after 108855 days.`},
}

func TestClock_RunTime(t *testing.T) {
	for _, y := range timeTests {
		actual := RunTime(int8(30),uint16(y.cycles))
		if actual != y.expected {
			t.Errorf("RunTime(%d): expected %s, actual %s", y.cycles, y.expected, actual)
		}
	}
}
func TestClock_RunAmount(t *testing.T) {
	for _, y := range amountTests {
		actual := RunAmount(int8(y.ballsAmount))
		if actual != y.expected {
			t.Errorf("RunTime(%d): expected %s, actual %s", y.ballsAmount, y.expected, actual)
		}
	}
}