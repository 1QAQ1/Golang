package main

import "fmt"

func main() {
	fmt.Println("huhu")

	sad := []int{1, 3, 5, 2, 7}

	sadd := twoSum(sad, 10)
	fmt.Print("huhu = ", sadd[0], sadd[1])

}

func twoSum(nums []int, target int) []int {
	huhu := make(map[int]int)
	ret := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		huhu[nums[i]] = i
	}
	for n := 0; n < len(nums); n++ {
		for k, v := range huhu {
			if k == target-nums[n] && n != v {
				ret[0] = v
				ret[1] = n
			}
		}
	}
	return ret

}

// func twoSum(nums []int, target int) []int {
//     hashTable := map[int]int{}
//     for i, x := range nums {
//         if p, ok := hashTable[target-x]; ok {
//             return []int{p, i}
//         }
//         hashTable[x] = i
//     }
//     return nil
// }
