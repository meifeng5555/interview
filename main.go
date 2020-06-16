package main

import (
	"fmt"
	"sort"
	"time"
)

type Depart struct {
	Name string
	Val int
}
type ByVal []Depart
func (p ByVal) Len() int           { return len(p) }
func (p ByVal) Less(i, j int) bool { return p[i].Val < p[j].Val }
func (p ByVal) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main()  {
	fmt.Println("interview")

	// 1.
	nums := make([]Depart, 0, 4)
	nums = append(nums, Depart{"A",10}, Depart{"B",7}, Depart{"C",5}, Depart{"D",4})
	// fmt.Println(DepartOptimal(nums, len(nums), 120))
	// fmt.Println(DepartOptimal2([]int{10,7,5,4}, 120))
	// fmt.Println(ValidateInviteCode("000000000000k0k0"))
	t := time.Now()
	fmt.Println(CombineMoney(3, 5))
	fmt.Println(CombineMoney2(15,100))
	fmt.Println(time.Since(t))
}

// 暴力美学
// nums map[string]int 初始数组
// n int 迭代次数
func DepartOptimal(nums []Depart, len int, n int) []Depart {
	if n <= 0 {
		return nums
	}
	sort.Sort(ByVal(nums))
	for i := 0; i < len-1; i++ {
		nums[i].Val += 1
	}
	nums[len-1].Val -= len - 1
	return DepartOptimal(nums, len, n-1)
}

// 动态规划
func DepartOptimal2(nums []int, n int) []int {
	dp, ln := make([][]int, 0, n), len(nums)
	dpSIdx := -1
	GetMaxIdx := func(nums []int) int {
		max, idx := 0, -1
		for i, val := range nums {
			if val > max {
				max = val
				idx = i
			}
		}
		return idx
	}
	Equals := func(a, b []int, l int) bool {
		for i := 0; i < l; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < n; i++ {
		maxIdx := GetMaxIdx(nums)
		if maxIdx == -1 {
			return nums
		}
		// 部门优化
		for in, _:= range nums {
			if in == maxIdx {
				nums[in] -= ln-1
			} else {
				nums[in] += 1
			}
		}
		ans := make([]int, ln); copy(ans, nums)
		dp = append(dp, ans)
		// 动态规划开始
		if i >= ln && Equals(dp[i], dp[i-ln], ln) {
			dpSIdx = i-ln
			break
		}
	}

	if dpSIdx != -1 {
		return dp[(n-1-dpSIdx)%ln+dpSIdx]
	}

	return nums
}

func ValidateInviteCode(s string) string {
	if len(s) != 16 {
		return "error"
	}

	charIntMap, num := make(map[byte]int), 1
	for i := byte('a'); i <= byte('z') ; i++ {
		if num >= 10 { num = 1 }
		charIntMap[i] = num
		num++
	}

	// 字符串长度为偶数的情况下
	// 逆向奇数位 = 正向偶数位
	// 逆向偶数位 = 正向奇数位
	// 否则，反转字符串，在操作
	var total int
	for i, _:= range s {
		var newVal int
		if s[i] >= 48 && s[i] <= 57 {
			newVal = int(s[i]-48)
		} else {
			newVal = charIntMap[s[i]]
		}
		if (i+1)%2 == 0 {
			total += newVal
		} else {
			newVal *= 2
			if newVal >= 10 { newVal -= 9 }
			total += newVal
		}
	}

	ret := "ok"
	if total%10 != 0 {
		ret = "error"
	}
	return ret
}

// 动态规划
func CombineMoney(n, m int) int {
	dp := make([][]int, n+1)
	coins := []int{1,2,5,10}
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1
	fmt.Println("coin, i, j, val")
	for _, coin := range coins {
		for i := 1; i <= n; i++ {
			for j := coin; j <= m; j++ {
				dp[i][j] += dp[i-1][j-coin]
				fmt.Printf("%d, %d, %d, %d\n",coin,i,j,dp[i][j])
			}
		}
	}
	return dp[n][m]
}
func CombineMoney2(n, m int) int {
	cnt := 0
	for i:=0;i<=n;i++ {
		for j:=0;j<=n-i;j++ {
			for y:=0;y<=n-i-j;y++ {
				for k:=0;k<=n-i-j-y;k++ {
					if i+2*j+5*y+10*k == m && i+j+y+k == n {cnt++}
				}
			}
		}
	}
	return cnt
}