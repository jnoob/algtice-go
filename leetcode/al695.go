package leetcode

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	maxArea, maxRegion := 0, 1
	regionsCount, regionMap := map[int]int{}, map[int]int{}
	for i, arr := range grid {
		hasNextLevel := i + 1 < len(grid)
		for k := 0; k< len(arr); {
			if arr[k] == 0 {
				k++
				continue
			} else {
				regionNum, regionCount, joinRegion := 0, 0, -1
				if arr[k] > 1 {
					if _, exist := regionMap[arr[k]]; exist {
						regionNum = regionMap[arr[k]]
						joinRegion = arr[k]
					} else {
						regionNum = arr[k]
					}
					regionCount = regionsCount[regionNum]
				} else {
					maxRegion++
					regionNum = maxRegion
				}
				start := k
				for m:=k;m<len(arr);m++ {
					if arr[m] == 1 {
						arr[m] = regionNum
						regionCount++
					} else if arr[m] == regionNum || arr[m] == joinRegion {

					} else if arr[m] > 1 {
						joinRegion = arr[m]
						regionCount += regionsCount[joinRegion]
						regionMap[joinRegion] = regionNum
					} else { // arr[m] <= 0
						k = m + 1
						break
					}
					if hasNextLevel && grid[i+1][m] == 1 {
						grid[i+1][m] = regionNum
						regionCount++
					}
				}
				if k == start {
					k = len(arr)
				}
				regionsCount[regionNum] = regionCount
			}
		}
	}
	for i,v:=range regionsCount {
		if v > maxArea {
			maxArea = v
			fmt.Printf("max-i:%v\n", i)
		}
	}
	for _, arr := range grid {
		fmt.Printf("%v\n", arr)
	}
	fmt.Printf("%v\n", regionsCount)
	return maxArea
}