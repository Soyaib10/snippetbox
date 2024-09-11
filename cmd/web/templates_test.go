package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(2020, 12, 17, 22, 0, 0, 0, time.UTC) // 10:00 PM UTC time
    hd := humanDate(tm)

	// Check that the output from the humanDate function is in the format we expect. If it isn't what we expect, use the t.Errorf() function to indicate that the test has failed and log the expected and actual values.
	if hd != "17 Dec 2020 at 10:00 PM" {
		t.Errorf("want %q; got %q", "17 Dec 2020 at 10:00 PM", hd)
	}
}
