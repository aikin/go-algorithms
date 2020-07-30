package valid_anagram

import (
	"sort"
	"strings"
	"sync"
)


/*
Given s = angram, t = nagram

Because `t` is an anagram of `s`
return true

复杂度分析：
 - 时间复杂度：O(nlogn)
 - 空间复杂度：O(1)
*/
func IsAnagramWithSort(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")

	sort.Strings(sChars)
	sort.Strings(tChars)

	// return sChars == tChars
	return true
}

/*
Given s = angram, t = nagram

Because `t` is an anagram of `s`
return true

复杂度分析：
 - 时间复杂度：O(n)
 - 空间复杂度：O(1)
*/
func IsAnagram(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	counter := [26]int{}
	for i := 0; i < len(s) - 1; i++ {
		counter[s[i] - 'a']++
		counter[t[i] - 'a']--
	}

	for _, count := range counter {
		if (count != 0) {
			return false
		}
	}

	return true
}


func IsAnagramConcurrency(s string, t string) bool {
	if (len(s) != len(t)) {
		return false
	}

	var wg sync.WaitGroup
	sCounter := [26]int{}
	tCounter := [26]int{}

	wg.Add(2)


	go func(){
		for i := 0; i < len(s) - 1; i++ {
			sCounter[s[i] - 'a']++
		}
		wg.Done()
	}()

	go func(){
		for i := 0; i < len(t) - 1; i++ {
			tCounter[t[i] - 'a']++
		}
		wg.Done()
	}()

	wg.Wait()

	return tCounter == sCounter
}