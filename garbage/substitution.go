package garbage

import (
	"sort"
	"strings"
)

type substitution struct {
	Origin      string
	Destination string
}

type orderedSubstitution []substitution

func (os orderedSubstitution) Len() int {
	return len(os)
}

func (os orderedSubstitution) Swap(i, j int) {
	os[i], os[j] = os[j], os[i]
}

// less is more
func (os orderedSubstitution) Less(i, j int) bool {
	return len(os[i].Origin) > len(os[j].Origin)
}

func (os *orderedSubstitution) Order() {
	sort.Sort(os)
}

func (os orderedSubstitution) toReplacerArray() []string {
	res := []string{}
	for _, elem := range os {
		res = append(res, elem.Origin, elem.Destination)
	}
	return res
}

func mapToOrderedSubstitution(m map[string]string) orderedSubstitution {
	var os orderedSubstitution

	for key, value := range m {
		os = append(os, substitution{strings.ToUpper(key), value})
	}

	os.Order()

	return os
}
