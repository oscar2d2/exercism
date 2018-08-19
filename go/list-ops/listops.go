package listops

type IntList []int
type unaryFunc func(int) int
type binFunc func(int, int) int
type predFunc func(int) bool

func (l IntList) Foldl(f binFunc, initial int) int {
	ans := initial

	if l.Length() > 0 {
		ans = l[1:].Foldl(f, f(ans, l[0]))
	}

	return ans
}

func (l IntList) Foldr(f binFunc, initial int) int {
	ans := initial
	n := l.Length()
	if n > 0 {
		ans = l[:n-1].Foldr(f, f(l[n-1], ans))
	}

	return ans
}

func (l IntList) Filter(f predFunc) IntList {
	n := l.Length()
	ans := make(IntList, n)

	ind := 0
	for i := 0; i < n; i++ {
		if f(l[i]) {
			ans[ind] = l[i]
			ind++
		}
	}

	return ans[:ind]
}

func (l IntList) Length() int {
	if l == nil || len(l) == 0 {
		return 0
	}

	return 1 + IntList(l[1:]).Length()
}

func (l IntList) Map(f unaryFunc) IntList {
	n := l.Length()
	ans := make(IntList, n)
	for i, elem := range l {
		ans[i] = f(elem)
	}
	return ans
}

func (l IntList) Reverse() IntList {
	n := l.Length()
	ans := make(IntList, n)
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		ans[i] = l[j]
	}
	return ans
}

func (l IntList) Append(that IntList) IntList {
	oriLength := l.Length()
	n := oriLength + that.Length()
	ans := make(IntList, n)
	for i, elem := range l {
		ans[i] = elem
	}
	offset := oriLength
	for i, elem := range that {
		ans[offset+i] = elem
	}
	return ans
}

func (l IntList) Concat(listOfList []IntList) IntList {
	n := l.Length()
	ans := make(IntList, n)

	for i := 0; i < n; i++ {
		ans[i] = l[i]
	}

	if len(listOfList) > 0 {
		ans = ans.Append(listOfList[0])
	}

	if len(listOfList) > 1 {
		ans = ans.Concat(listOfList[1:])
	}

	return ans
}
