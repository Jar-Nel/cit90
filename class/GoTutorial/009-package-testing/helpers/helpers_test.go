package helpers

import "testing"

func TestBar(t *testing.T) {
	want:=421
	if got:=Bar(); got != want{
		t.Errorf("Bar() error  want: %v, got %v",want,got)
	}
}