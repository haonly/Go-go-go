// https://programmers.co.kr/learn/courses/30/lessons/42576

package Solution42576

func Solution(participant []string, completion [] string) string{
	dict := map[string]int{}

	for _, pName := range participant{
		dict[pName] += 1
	}

	for _, cName := range completion{
		dict[cName] -= 1
	}

	for _, v := range participant{
		if dict[v] != 0 {
			return v
		}
	}

	return ""
}
