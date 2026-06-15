package main

import "testing"

func TestFullFlow(t *testing.T) {
	got, err := run("test_files/good_basic_case")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "1 3 N\n5 1 E" {
		t.Errorf("got %q, expected %q", got, "1 3 N\n5 1 E")
	}
}
