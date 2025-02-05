package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	/*语法练习*/
	//GP1
	fmt.Println("Hello World!")

	//GP2  var 变量名 变量类型 = 赋值
	var (
		name   string = "小明"
		age    int    = 23
		gender bool   = true
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(gender)

	//GP3-1 变量默认赋值  变量 变量名 类型
	var (
		name1   string
		age1    int
		gender1 bool
	)
	fmt.Println(name1)
	fmt.Println(age1)
	fmt.Println(gender1)

	//GP3-2 结构体
	type Person struct {
		name   string
		age    int
		gender bool
	}
	v1 := Person{}
	fmt.Println(v1.name)
	fmt.Println(v1.age)
	fmt.Println(v1.gender)
	/*
		结构体的类型为struct ，并且以type开头。
		结构体不需要用var来修饰
		结构体构造的时候:= 来自动定义变量类型，同时创建基本类型的时候需要用大括号{}，而不是小括号().
	*/

	//GP4-1 常量 显示定义
	const a string = "中国"
	const b string = "英国"
	const c string = "美国"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//GP4-2 常量 隐式定义
	const (
		a1 = "中国"
		b1 = "英国"
		c1 = "美国"
	)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)

	//GP5-1 值和指针  *取值，&取址
	var (
		a2 int = 1
		b2 int = 2
	)
	var ptr1 *int = &a2
	var ptr2 *int = &b2
	var ret1 bool
	var ret2 bool
	ret1 = ptr1 == ptr2
	ret2 = *ptr1 == *ptr2
	fmt.Println([]bool{ret1, ret2})

	//GP5-2 切片
	var (
		a3 int = 1
		b3 int = 2
	)
	var myslice []bool
	var ptr3 *int = &a3
	var ptr4 *int = &b3
	myslice = append(myslice, ptr3 == ptr4)
	myslice = append(myslice, *ptr3 == *ptr4)
	fmt.Println(myslice)

	//GP6 拼接字符串 +, for _,x := range s 遍历数组
	var s []string = []string{"hello", "-", "world"}
	var result string
	for _, x := range s {
		result += x
	}
	fmt.Println(result)

	//GP7 字符个数
	var s2 string = "小明的英文名叫jack"
	var run = []rune(s2)
	fmt.Println(run)
	ans := len(run)
	fmt.Println(ans)
	/*
		1,汉字是采用unicode编码，占三个字节
		2,字符串转化为rune数组
		2,rune是int32的别名（-231~231-1），对比byte（-128～127），可表示的字符更多
		3,len()可以求出rune数组的长度
	*/

	//GP8 回文判断
	var d int = 121121
	var str string = strconv.Itoa(int(d))
	var length = len(str)
	var ret bool = true
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			fmt.Println(!ret)
			break
		}
	}
	fmt.Println(ret)
	/*
		不能直接string(x)
		不同类型之间转换，建议用strconv的转换函数
		整型转化为字符串,字符串的遍历
	*/

	//GP9 格式化字符串 fmt.Sprintf来格式化字符串
	var ddd int = 100
	var str2 string = fmt.Sprintf("%d", ddd)
	fmt.Println(str2)
	/*
		Go 使用 import 关键字来导入包
		Go 可以使用 fmt.Sprintf 来格式化字符串，fmt.Sprintf(格式化样式, 参数列表…),格式化样式如下：
		%v   按值的本来值输出
		%+v  在 %v 基础上，对结构体字段名和值进行展开
		%#v  输出 Go 语言语法格式的值
		%T   输出 Go 语言语法格式的类型和值
		%%   输出 % 本体
		%b   整型以二进制方式显示
		%o   整型以八进制方式显示
		%d   整型以十进制方式显示
		%x   整型以十六进制方式显示
		%X   整型以十六进制、字母大写方式显示
		%U   Unicode 字符
		%f   浮点数
		%p   指针，十六进制方式显示
	*/

	//GP10-1 strconv.ParseInt 返回一个int64值，int是int32。bitSize后面写32也是int64类型的，改不了
	var aa string = "100"
	var bb string = "200"
	num1, _ := strconv.ParseInt(aa, 10, 64)
	num2, _ := strconv.ParseInt(bb, 10, 64)
	cc := num1 + num2
	res1 := strconv.Itoa(int(cc))
	fmt.Println(res1)
	//GP10-2 strconv方法
	var aa1 string = "100"
	var bb1 string = "200"
	num11, _ := strconv.Atoi(aa1)
	num22, _ := strconv.Atoi(bb1)
	res2 := strconv.Itoa(num11 + num22)
	fmt.Println(res2)
	/*
		strconv.ParseInt 是将字符串转换为数字的函数,
		参数1 数字的字符串形式,
		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制,
		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	*/

	//GP11-1 手动排序
	var (
		j int = 3
		q int = 4
		k int = 2
	)
	max11 := max1(j, max(q, k))
	min11 := min1(j, min(q, k))
	fmt.Println(max11 - min11)

	//GP11-2 调用排序
	var k1, k2, k3 int = 3, 4, 2
	kk := []int{k1, k2, k3}
	sort.Ints(kk)
	fmt.Println(kk[2] - kk[0])

	//GP15 逻辑运算
	var a11, b11 bool = true, true
	var ret11 []bool
	ret11 = append(ret11, a11 && b11)
	ret11 = append(ret11, b11 || b11)
	ret11 = append(ret11, !a11)
	ret11 = append(ret11, !b11)
	fmt.Println(ret11)
	/*
		golang中，&& 表示逻辑 AND 运算符。
		golang中，｜｜ 表示逻辑 OR 运算符。
		golang中，！ 表示逻辑 NOT 运算符。
	*/

	//GP16 位运算
	var a22, b22 int = 1, 1
	var slice22 []int
	slice22 = append(slice22, a22&b22)
	slice22 = append(slice22, a22|b22)
	slice22 = append(slice22, a22^b22)
	fmt.Println(slice22)

	//GP17 奇偶
	var a33 int = 17
	ret33 := a33 % 2
	if ret33 == 1 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}

	//GP19 创建切片
	length11 := 10
	capacity11 := 10
	var slice11 []int = make([]int, length11, capacity11)
	for i, _ := range slice11 {
		slice11[i] = i
	}
	fmt.Println(slice11)
	/*
		make() 函数来创建切片:var slice1 []type = make([]type, len) 指定容量，
		其中 capacity 为可选参数：make([]T, length, capacity)
	*/

	//GP20 切片复制，length和capacity要一致
	var src []int = []int{1, 2, 3, 4}
	var des []int = []int{5, 6}
	des = make([]int, len(src), cap(src)) //修改des大小与src一致
	copy(des, src)
	fmt.Println(des)
	/*
		copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice
		可指向同一底层数组，允许元素区间重叠。
	*/

	//GP21 删除切片里的元素
	var sk []int = []int{1, 2, 3, 4, 5, 6}
	var index int = 3
	sk = append(sk[:index], sk[index+1:]...)
	fmt.Println(sk)

	//GP22-1 切片最大、最小值
	var sss []int = []int{1, 2, 3, 4, 5, 6, 7, 7, 6}
	var retm []int = []int{sss[0], sss[0]}
	for _, xm := range sss {
		retm[0] = min1(retm[0], xm)
		retm[1] = max1(retm[1], xm)
	}
	fmt.Println(retm)

	//GP22-2 sort方法
	sort.Ints(sss)
	fmt.Println(sss[0], sss[len(sss)-1])

	//GP23 反转切片
	sss = []int{1, 2, 3, 4, 5, 6, 7, 7, 6}
	for i := 0; i < len(sss)/2; i++ {
		sss[i], sss[len(sss)-1-i] = sss[len(sss)-1-i], sss[i]
	}
	fmt.Println(sss)

	//GP30
	s1 := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6}
	var res []int
	intMap := make(map[int]int)
	for _, x1 := range s1 {
		intMap[x1]++
	}
	fmt.Println(intMap)
	for y, _ := range intMap {
		if intMap[y] == 1 {
			res = append(res, y)
		}
	}
	sort.Ints(res)
	fmt.Println(res)

	//GP38
	score := []int{1, 2, 8, 4, 2}
	target := 7
	t := canPass(score, target)
	fmt.Println(t)

	//GP46
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		panic("three")
	}()
	defer func() {
		panic("two")
	}()
	panic("one")
}

func canPass(score []int, target int) bool {
	// write code here
	for _, v := range score {
		if v > target {
			return true
		}
	}
	return false
}

func max1(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min1(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
