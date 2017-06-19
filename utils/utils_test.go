package utils_test

import "testing"
import "coloso-queue/utils"

type compareStruct struct {
	A string
	B int
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
