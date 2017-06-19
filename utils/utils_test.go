package utils_test

import "testing"
import "coloso-queue/utils"

type compareStruct struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func TestCompareStructsEquals(t *testing.T) {
	var structA, structB compareStruct
	var equals bool
	var keys = []string{"A", "B"}

	structA = compareStruct{A: "A", B: 1}
	structB = compareStruct{A: "A", B: 1}

	equals = utils.CompareStructs(structA, structB, keys)

	if !equals {
		t.Fatalf("Bad compare determination, expected: true, got: false")
	}
}

func TestCompareStructsDifferents(t *testing.T) {
	var structA, structB compareStruct
	var equals bool
	var keys = []string{"A", "B"}

	structA = compareStruct{A: "A", B: 1}
	structB = compareStruct{A: "B", B: 1}

	equals = utils.CompareStructs(structA, structB, keys)

	if equals {
		t.Fatalf("Bad compare determination, expected: false, got: true")
	}
}

func TestStructToString(t *testing.T) {
	var structA compareStruct
	var expectedString, resultString string

	expectedString = `{"a":"A","b":1}`
	structA = compareStruct{A: "A", B: 1}

	resultString = utils.StructToString(structA)

	if resultString != expectedString {
		t.Fatalf("Strings doesn't match\nexpected: %s\ngot: %s", expectedString, resultString)
	}
}
