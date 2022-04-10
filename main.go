package main

import (
	"fmt"
)

func main() {
	query := "ดี"
	// query := os.Args[1]

	dict := map[string]string{
		"good":    "ดี",
		"great":   "เยี่ยม",
		"perfect": "ยอดเยี่ยม",
		"nil":     " ",
	}
	// delete(dict, "good") remove ข้อมูล ใน dict

	thai := make(map[string]string, len(dict))
	for key, value := range dict {
		thai[value] = key
	}

	if value, ok := thai[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	if value, ok := dict[query]; ok { // ใส่ 2 ค่า เพิ่ม ok bool เพื่อเช็คความถูกต้อง
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)
}
