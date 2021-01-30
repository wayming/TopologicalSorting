package topologicalsorting

import (
	"reflect"
	"testing"
)

func TestMultipleKeysSort(t *testing.T) {

	topoMapSort :=
		Sort(map[string][]string{
			"aa": {
				"dd", "cc",
			},
			"dd": {
				"gg",
			},
			"gg": {
				"hh", "ii",
			},
			"cc": {
				"jj", "ii",
			},
		})
	expect := []string{"hh", "ii", "gg", "dd", "jj", "cc", "aa"}
	if !reflect.DeepEqual(topoMapSort, expect) {
		t.Fatal("Expected:", expect, "Actual:", topoMapSort)
	}
}
