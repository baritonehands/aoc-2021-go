package main

import (
	_ "embed"
	"github.com/baritonehands/aoc-2021-go/day16/packet"
	"strings"
)

//go:embed input.txt
var input string

var hexBin = map[byte]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func main() {
	bin := strings.Builder{}
	for _, hex := range input {
		bin.WriteString(hexBin[byte(hex)])
	}
	data := []byte(bin.String())

	//fmt.Printf("%v\n", bin.String())
	p1, _ := packet.ParseN(data, -1)
	sum := 0
	for _, p := range p1 {
		sum += p.VersionSum()
	}
	println("part1", sum)
	//fmt.Printf("%v = %d, %v\n", p1, sum, rest)

	println("part2", p1[0].Value())
}
