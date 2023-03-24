package main

import "testing"

func TestPrepareAppDataBeforeRun(t *testing.T) {
	err := prepareAppDataBeforeRun()
	if err != nil {
		t.Errorf("[func TestPrepareAppDataBeforeRun] - cannot run function for tests")
	}
}
