package main

import "fmt"

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func main() {

	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"dog", "racecar", "car"}
	// strs := []string{"a"}
	strs := []string{"ab", "a"}

	commonStr := ""
	i := 0

	if len(strs) == 0 {
		return
	}

	for {

		if i >= len(strs[0]) {
			break
		}

		tmpCmnStr := commonStr + string(strs[0][i])

		isMatched := true
		for j := 0; j < len(strs); j++ {

			if len(tmpCmnStr) > len(strs[j]) {
				isMatched = false
				break
			}

			// checking if current prefix is there
			if tmpCmnStr != strs[j][:len(tmpCmnStr)] {
				isMatched = false
				break
			}
		}
		if !isMatched {
			break
		}
		commonStr = tmpCmnStr
		i += 1
	}

	if commonStr == "" {
		fmt.Println("No common prefix")
	} else {
		fmt.Println("common Prefix found: ", commonStr)
	}

}
