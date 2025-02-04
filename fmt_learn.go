package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//1、Print例子
	const name1, age1 = "Kim", 22
	fmt.Print(name1, " is ", age1, " years old.\n")

	//2、Println例子 不用空格
	const name2, age2 = "Jim", 22
	fmt.Println(name2, "is", age2, "years old.")

	//3、Printf例子
	const name3, age3 = "Vitalik", 28
	fmt.Printf("%s is %d years old.\n", name3, age3)

	//4、Sprint例子
	const name4, age4 = "liming", 30
	s1 := fmt.Sprint(name4, " is ", age4, "years old.\n")
	io.WriteString(os.Stdout, s1)

	//5、Sprintln例子
	const name5, age5 = "gavin", 40
	s2 := fmt.Sprintln(name5, "is", age5, "years old.")
	io.WriteString(os.Stdout, s2)

	//Fscanf例子
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf:%v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)

	var ptr *int
	fmt.Printf("ptr的值为：%x\n", ptr)
}
