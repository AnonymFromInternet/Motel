package main

import "testing"

func TestPrepareAppDataBeforeRun(t *testing.T) {
	err := prepareAppDataBeforeRunGetDBConnectionPool()
	if err != nil {
		t.Errorf("[func TestPrepareAppDataBeforeRun] - cannot run function for tests")
	}
}
