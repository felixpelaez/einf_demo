package util

import "testing"

func TestGetFourRandomChars(t *testing.T) {

	randomchars := GetRandomChars(4)

	if len(randomchars) != 4 {
		t.Errorf("number ofrandom chars not correct, got: %d, want 4", len(randomchars))

	}
}

func TestGetTenRandomChars(t *testing.T) {

	t.Log("get 10 random chars")
	randomchars := GetRandomChars(10)

	if len(randomchars) != 10 {
		t.Errorf("number of random chars not correct, got: %d, want 10", len(randomchars))

	}
}

func TestTwoConsecutiveRequests(t *testing.T) {

	randomchars := GetRandomChars(4)

	randomchars1 := GetRandomChars(4)

	if randomchars == randomchars1 {
		t.Errorf("Two consecutive request does give same chars %v", randomchars)
	}

}
