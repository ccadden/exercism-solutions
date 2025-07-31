package stringset

import "strings"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]bool

func New() Set {
	return map[string]bool{}
}

func NewFromSlice(l []string) Set {
	set := New()

	for _, elem := range(l) {
		set.Add(elem)
	}

	return set
}

func (s Set) String() string {
	var b strings.Builder
	b.WriteString("{")
	i := 0

	for key := range(s) {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString("\"")
		b.WriteString(key)
		b.WriteString("\"")
		i++
	}
	b.WriteString("}")

	return b.String()
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]

	return ok
}

func (s Set) Add(elem string) {
	if s.Has(elem) {
		return
	}

	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for key := range(s1) {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for key := range(s1) {
		if s2.Has(key) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}

	return Subset(s1, s2)
}

func Intersection(s1, s2 Set) Set {
	result := New()

	for key := range(s1) {
		if s2.Has(key) {
			result.Add(key)
		}
	}

	return result
}

func Difference(s1, s2 Set) Set {
	result := New()

	for key := range(s1) {
		if !s2.Has(key) {
			result.Add(key)
		}
	}

	return result
}

func Union(s1, s2 Set) Set {
	result := New()

	for key := range(s1) {
		result.Add(key)
	}

	for key := range(s2) {
		result.Add(key)
	}

	return result
}
