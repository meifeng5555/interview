1. 部门优化

   两种解，1.暴力递归，2.DP找规律，两种方法都有个缺陷，在没找到规律前，每次都需要找最大值，这里想不到更好的解法

   ```go
   type Depart struct {
   	Name string
   	Val int
   }
   type ByVal []Depart
   func (p ByVal) Len() int           { return len(p) }
   func (p ByVal) Less(i, j int) bool { return p[i].Val < p[j].Val }
   func (p ByVal) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

   // 1 暴力递归
   // 暴力美学，每次递归最新数组，一路往下
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

   // 2.动态规划
   // 在暴力递归后，传入不同参数，发现数据会趋近于规律化
   // 因为每次迭代，总和不变，当数组值间的差值越来越小后，就会开始循环一组数（组的数量跟数组的大小一致）
   // 当检查到一组循环值出现时，首次出现的位置标记 dpSIdx，终止迭代，因为此时可以算出第N轮的迭代值(N-1 >= dpSIdx)，则第N轮迭代值等于(ln数组长度) dp[(n-1-dpSIdx)%ln+dpSIdx]
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
   ```

   



2. 邀请码检测：

   ​	题目要求对逆序奇数，逆序偶数位分别计算

   ​	若字符串长度为偶数，

   ​	则 逆向奇数位 = 正向偶数位, 逆向偶数位 = 正向奇数位

   ​	若字符串长度为奇数，则需按数组中位数反转字符串，在计算。

   ​	唯一要处理的点就是ascii码的'0'-'9'

   ```go
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
   ```

   

   3. 游戏币组合

      还是两种解法：暴力循环，和动态规划

      看到题目的时候，就想到了动态规划，但是，第一遍没找出规律，遂放弃，直接用循环解决，循环就是控制每个币的数量，迭代出所有情况，然后判断出符合条件的情况。

      动态规划：后来在纸上画了图，模拟了下n=3,m=5问题的拆解情况，发现也是跟暴力循环的思想一致

      设置 dp矩阵, n+1,m+1维，为什么+1，便于写代码不容易糊，要不然 就是dp(2,4)对应3,5了。

      dp(3,5)和前面元素的关系，也是控制数量减少，面值下降

      dp(3,5) 

      ​	= dp(2,4) + dp(2,3)

      ​	= dp(1,3) + dp(1,2) + dp(1,2) + dp(1,1)

      ​	= dp(0,2) + dp(0,1) + dp(0,1) + dp(0,0) + dp(0,1) + dp(0,0) + dp(0,0)

      ​	= 3 （即 1,2,2|2,1,2|2,2|1）三种等于1种，要去重

      此时公式已经推到完毕即 dp(n,m) = dp(n-1, m-1) + dp(n-1, m-2) + dp(n-1,m-5) + dp(n-1,m-10)

      此时还剩下一个问题，重复如何去掉，第一时间想到了记录硬币，然后去重，写着写着发现很麻烦，并且判断硬币数组是否相同还要排序，不利于效率。

      ```go
      // 动态规划
      func CombineMoney(n, m int) int {
      	dp := make([][]int, n+1)
      	coins := []int{1,2,5,10}
      	for i := 0; i < n+1; i++ {
      		dp[i] = make([]int, m+1)
      	}
      	dp[0][0] = 1
      	for _, coin := range coins {
      		for j := 1; j <= n; j++ {
      			for k := coin; k <= m; k++ {
      				dp[j][k] += dp[j-1][k-coin]
      			}
      		}
      	}
      	return dp[n][m]
      }

      // 暴力美学
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
      ```

      
   
   4.有趣的两位数
   
   由公式 ab * cd = ba * dc 可以推出 ac = bd
   
   因为是两位数，所以范围可以缩小到1-9的乘积，枚举出99乘法表种所有的 ac bd 集合，然后循环组装成ab cd
   
   ```go
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
   ```
   
   

6.单链表处理

   不新建额外的空间的情况下，应该就是将链表拆分成两部分，然后后半部分链表反转，按照中位值，将前半部分和反转后的后半部分链表重新关联。

```go
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
```



7. 设计unionid
   1. 市面成熟的方案：uuid，snowflake，或者mysql分步长设置自增id
   2. 基本思想就是保持每次获取到变化的值，这样才能保证唯一性，不谈多机器的情况下，请求ip(16进制)+时间戳+内存计数，基本能满足，且不消耗额外网络资源，多机器就是加机器的唯一标识，如mac地址等等