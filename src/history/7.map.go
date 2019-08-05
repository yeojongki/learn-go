package main

import "fmt"

// any inner type can be a key excluding `slice, map, function`
func main() {
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	var m2 map[string]bool = map[string]bool{
		"k1": true,
		"k2": false,
	}

	// cant assign value
	m3 := make(map[string]int) // m2 === map[]

	// empty map
	var m4 map[string]int // m3 === nil

	fmt.Println(m1, m2, m3, m4)
	fmt.Println(m3 == nil) // false
	fmt.Println(m4 == nil) // true

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// get map value
	v1, exist := m1["key1"]
	fmt.Println(v1, exist)

	// get with wrong key
	vWrong, exist := m1["wrong"]
	fmt.Println(vWrong, exist) // zero value

	// delete value
	delete(m1, "key1")
	v1, exist = m1["key1"]
	fmt.Println(v1, exist)

	// get length
	fmt.Println(len(m1))
}
