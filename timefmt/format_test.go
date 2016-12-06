package timefmt

import (
	"testing"
)

type TestCase struct {
	Input  string
	Expect string
}

var testCases []TestCase = []TestCase{
	TestCase{
		Input:  "YYYY",
		Expect: "2006",
	},
	TestCase{
		Input:  "YYYYMMDD",
		Expect: "20060102",
	},
	TestCase{
		Input:  "YYYY/MM/DD",
		Expect: "2006/01/02",
	},
	TestCase{
		Input:  "YYYY/MM/DD hh:mm",
		Expect: "2006/01/02 15:04",
	},
}

func TestFormat(t *testing.T) {
	for _, testCase := range testCases {
		actual := Format(testCase.Input)
		if actual != testCase.Expect {
			t.Errorf("got %v\nwant %v", actual, testCase.Expect)
		}
	}
}
