package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}

	acc := initial

	for _, val := range(s) {
		acc = fn(acc, val)
	}

	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}

	acc := initial

	for i := s.Length() - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}

	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := IntList{}

	for _, val := range(s) {
		if fn(val) {
			res = res.Append([]int{val})
		}
	}

	return res
}

func (s IntList) Length() int {
	length := 0

	for range(s) {
		length++
	}

	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	res := IntList{}

	for _, val := range(s) {
		res = res.Append(IntList{fn(val)})
	}

	return res
}

func (s IntList) Reverse() IntList {
	res := IntList{}

	for i := s.Length() - 1 ; i >= 0; i-- {
		res = res.Append(IntList{s[i]})
	}

	return res
}

func (s IntList) Append(lst IntList) IntList {
	appended := make(IntList, s.Length()+lst.Length())

	copy(appended, s)
	copy(appended[s.Length():], lst)

	return appended
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range(lists) {
		s = s.Append(list)
	}

	return s
}
