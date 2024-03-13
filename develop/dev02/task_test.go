package main

import "testing"

type TestCase struct {
	input    string
	expected string
}

func TestUnpack(t *testing.T) {

	testCases := []TestCase{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", "некорректная строка"},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
		{"q45", "некорректная строка"},
	}

	for _, tc := range testCases {
		res, err := Unpack(tc.input)
		if err != nil {
			if err.Error() == tc.expected {
				continue
			}
			t.Errorf("error: %s", err)
		}
		if res != tc.expected {
			t.Errorf("expected: %s, got: %s", tc.expected, res)
		}
	}

}
