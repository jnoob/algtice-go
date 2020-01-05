package leetcode

import (
	"fmt"
	"strings"
)

func racecarCmds(target int) string {
	memt := make(map[int]string)
	return commands(target, memt)
}

func racecar(target int) int {
	memt := make(map[int]string)
	return len(commands(target, memt))
}



func commands(target int, memt map[int]string) string {
	if target == 0 {
		return ""
	}

	var cmds string
	cmds, contains := memt[target]
	if contains {
		fmt.Printf("=====> %v\n\t from mem -> %v\n", target, cmds)
		return cmds
	}

	countA, distance, prev := 0, 0, 0
	for distance < target {
		countA ++
		prev = distance
		distance = 1 << countA - 1
	}

	if distance == target {
		cmds = strings.Repeat("A", countA)
	} else {
		overCmds := "AR" + commands(distance - target, memt)
		behindCmds := simplyfyRR(commands(target - prev, memt))
		if len(overCmds) > len(behindCmds) {
			cmds = strings.Repeat("A", countA - 1) + behindCmds
			discard :=  strings.Repeat("A", countA - 1) + overCmds
			fmt.Printf("\t discard -> %v\n", discard)
		} else {
			cmds = strings.Repeat("A", countA - 1) + overCmds
			discard := strings.Repeat("A", countA - 1) + behindCmds
			fmt.Printf("\t discard -> %v\n", discard)
		}
	}
	memt[target] = cmds
	fmt.Printf("=====> %v\n\t eval -> %v\n", target, cmds)
	return cmds
}

func simplyfyRR(cmds string) string {
	RCount, lastRIndex := 0, -1
	for j := len(cmds) - 1; j >= 0 ; j --{
		if cmds[j] == 'R' {
			RCount++
			if lastRIndex < 0 {
				lastRIndex = j
			}
		}
	}
	// is odd, last Rotation can be simplified
	if RCount & 0b1 > 0 {
		tmp := cmds[lastRIndex:] + "R" + cmds[0:lastRIndex]
		fmt.Printf("\t %v simplified to %v\n", "RR" + cmds, tmp)
		return tmp
	} else {
		return "RR" + cmds
	}
}