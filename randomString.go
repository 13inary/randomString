package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var leng int
var scope string
var src []string
var number = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var small = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z"}
var big = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z"}
var char1 = []string{
	"!", "@", "#", "$", "%", "^", "&", "*"}
var char2 = []string{
	"(", ")", "-", "=", "+", "~", "[", "]", "{", "}",
	"|", ";", ":", ",", ".", "?", "<", ">", "_"}

// " ' \ / `
var all = [][]string{
	number, small, big, char1, char2}

func getRandomPassword() (randomPassword string) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= leng; i++ {
		tmp := src[rand.Intn(len(src))]
		randomPassword += tmp
	}
	return
}

// setSrc
func setSrc() {
	for k, v := range scope {
		if v == '1' {
			src = append(src, all[k]...)
		}
	}
}

func main() {
	flag.IntVar(&leng, "l", 18, "length of password")
	flag.StringVar(&scope, "s", "11111", "char2 char1 big small number")
	flag.Parse()
	setSrc()
	fmt.Println(getRandomPassword())
}
