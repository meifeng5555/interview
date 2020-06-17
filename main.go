package main

import (
	"fmt"
	"sort"
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
	// t := time.Now()
	fmt.Println(CombineMoney(3, 5))
	// fmt.Println(CombineMoney2(10,50))
	// fmt.Println(time.Since(t))
	// fmt.Println(FunnyTwoNums())
	// root := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}}
	// root = SingleList(root)
	// for root != nil {
	// 	fmt.Println(root)
	// 	root = root.next
	// }
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
				if coin == 2 && i==1 && j==1 {
					fmt.Printf("%d, %d, %d, %d\n", coin, i-1, j-coin, dp[i-1][j-coin])
				}
			}
		}
	}
	return dp[n][m]
}
func CombineMoney2(n, m int) int {
	cnt := 0
	fmt.Println(1,2,5,10)
	for i:=0;i<=n;i++ {
		for j:=0;j<=n-i;j++ {
			for y:=0;y<=n-i-j;y++ {
				for k:=0;k<=n-i-j-y;k++ {
					if i+2*j+5*y+10*k == m && i+j+y+k == n {cnt++;fmt.Println(i,j,y,k)}
				}
			}
		}
	}
	return cnt
}

// ab * cd = ba * dc => ac = bd
// 满足99乘法表的组合的ac，bd相互组合即可
func FunnyTwoNums() [][2]int {
	funnyMap := make(map[int][][2]int)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			funnyMap[i*j] = append(funnyMap[i*j], [2]int{i,j})
		}
	}
	funnyArr := make([][2]int, 0, 100)
	for _, val := range funnyMap {
		l := len(val)
		if l >= 2 {
			for i := 0; i < l-1; i++ {
				for j := i+1; j < l; j++ {
					a := val[i][0]
					c := val[i][1]
					b := val[j][0]
					d := val[j][1]
					funnyArr = append(funnyArr, [2]int{a*10+b, c*10+d})
				}
			}
		}
	}
	return funnyArr
}

type Node struct {
	data int
	next *Node
}
// 分治法，将两边拆分为两段，反转后部分的链表，两个链表同步遍历
func SingleList(root *Node) *Node {
	len := 0
	head := &Node{0, root}
	for root != nil{
		root = root.next; len++
	}
	midNode := head.next
	mid, i := len/2, 1
	for i <= mid {
		midNode = midNode.next
		i++
	}
	midNode = reverseList(midNode)
	newHead := &Node{-1, nil}
	dummy := &Node{-2, newHead}
	root = head.next
	for j := 1; j <= mid; j++ {
		newHead.next = root
		root = root.next
		newHead.next.next = midNode
		midNode = midNode.next
		newHead = newHead.next.next
	}
	return dummy.next.next
}
func reverseList(root *Node) *Node {
	if root.next == nil {
		return root
	}

	head := &Node{-1, root}
	prev := root
	pCur := root.next
	for pCur != nil {
		prev.next = pCur.next
		pCur.next = head.next
		head.next = pCur
		pCur = prev.next
	}
	return head.next
}