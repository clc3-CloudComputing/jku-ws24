package main

import "testing"

func TestGetMinuteUp(t *testing.T) {

	res := getMinute(1, 30)
	if res != 2 {
		t.Errorf("getMinute(1, 30) = %d; wanted 2", res)
	}
}

func TestGetMinuteDown(t *testing.T) {

	res := getMinute(1, 15)
	if res != 1 {
		t.Errorf("getMinute(1, 15) = %d; wanted 1", res)
	}
}
