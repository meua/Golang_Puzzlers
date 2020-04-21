package main

import "testing"

func TestFailLog(t *testing.T) {
	//t.FailNow()
	t.Logf("navy is man, age is % d", 23)
}
