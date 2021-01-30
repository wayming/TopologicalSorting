package topologicalsorting

import (
	"sort"
)

// Sort maps keyed by string, and values are array of strings
func Sort(topoMap map[string][]string) (ret []string) {
	var items []string
	for item := range topoMap {
		items = append(items, item)
	}
	sort.Strings(items)

	var doneItem = make(map[string]bool)

	var visitAll func(remainingItems []string)
	visitAll = func(remainingItems []string) {
		for _, k := range remainingItems {
			if doneItem[k] {
				continue
			}

			if topoMap[k] == nil || len(topoMap[k]) == 0 {
				ret = append(ret, k)
			} else {
				visitAll(topoMap[k])
				ret = append(ret, k)
			}

			doneItem[k] = true
		}
	}

	visitAll(items)
	return ret
}
