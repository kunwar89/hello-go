package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var out []int
var g_var int

func int_and_add(s string) {
	num, err := strconv.Atoi(s)
	if err == nil && num > 0 {
		g_var = g_var + num
		fmt.Printf("+%v +", g_var)
	}

}

func xint() {

	elreader := bufio.NewReader(os.Stdin)
	lim, _ := elreader.ReadString('\n')
	lim_slice := strings.Fields(lim)
	fmt.Println(lim_slice[0])
	limit, _ := strconv.Atoi(lim_slice[0])
	g_var = 0
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	xint_slice := strings.Fields(text)
	//slice_len := len(xint_slice)
	recur_for_ia(limit, int_and_add, xint_slice)
	out = append(out, g_var)
	fmt.Printf("\n %v \n +", out)

}

func recur_for_ia(count int, fp func(string), args []string) int {

	if count == 0 {
		return 1
	} else {
		fp(args[count-1])
		recur_for_ia(count-1, fp, args)
		return 2
	}
}

func recur_for_cases(count int) int {

	if count == 0 {
		return 1
	} else {
		xint()
		recur_for_cases(count - 1)
		return 2
	}
}

func main() {

	elreader := bufio.NewReader(os.Stdin)
	lim, _ := elreader.ReadString('\n')
	lim_slice := strings.Fields(lim)
	fmt.Println(lim_slice[0])
	limit, _ := strconv.Atoi(lim_slice[0])
	recur_for_cases(limit)

}