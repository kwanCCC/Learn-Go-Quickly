package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	five()
	six()
	seven()
	slicePractice()
	praticeString()
}
func six() {
	str := "dsjkdshdjsdh....js"
	bytes := []byte(str)
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str, len(bytes), utf8.RuneCount(bytes))
}

func five() {
	fmt.Println("===Question Five===")
	str := "A"
	for index := 0; index < 100; index++ {
		fmt.Printf("%s\n", str)
		str = str + "A"
	}
}
func seven() {
	fmt.Println("===Question Seven===")
	s := "dsjkdshdjsdh....js"
	r := []rune(s)
	after := []byte(s)
	copy(r[4:4+3], []rune("abc"))
	copy(after[4:7], []byte("abc"))
	fmt.Printf("Before: %s\n", s)
	fmt.Printf("After : %s\n", string(r))
	fmt.Printf("After : %s\n", string(after))

}

func slicePractice() {
	a := [...]int{1, 2, 3}
	b := [...]int{4, 5, 6}
	c := a
	e := make([]int, 10)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	a = [...]int{7, 8, 9}
	fmt.Println(`after redeclare a`)
	fmt.Println(a)
	fmt.Println(`c 仍然指向的是 a 的指向的数组`)
	fmt.Println(c)
	f := a[:]
	fmt.Println(`生成 指向 a 的 slice f , len 与 cap 与 a 一样`)
	fmt.Println("注意 \n 这里的 a 已经被 46 行重新赋值")
	fmt.Println(f)
	fmt.Println("看我 \n 再给换回来")
	f = c[:]
	fmt.Println(f)
	g := f[1:2]
	fmt.Println(`再来一个 从slice 中 slice 出来`)
	fmt.Println(g)

	h := map[int]string{1: `1`,}
	fmt.Println(h)

	fmt.Printf("\\%T\n", a)
	fmt.Printf("\\%T\n", g)
	ints := append(f, 1, 2, 3)
	fmt.Println(ints)

}

func praticeString() {
	a := `123456789`
	b := []rune(a)
	fmt.Println(`origin string ` + a)
	fmt.Print(`change to rune `)
	fmt.Println(b)
	c := []byte(a)
	fmt.Println(c)
}
