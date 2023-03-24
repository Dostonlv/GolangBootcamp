package main

import "fmt"

// NECHTA WOOF SOZINI YASAB BILISH BO'YICHA MASALA

func main() {
	s := "ooooooofw"
	var j []string
	for _, i := range s {
		j = append(j, string(i))
	}
	var w, o, f int
	for i := 0; i < len(j); i++ {
		if j[i] == "w" {
			w++
		} else if j[i] == "o" {
			o++
		} else if j[i] == "f" {
			f++
		}
	}

	// o%2!=0-> o-1 -> o/2 ->f<w -> f

	if (w%2 == 0 && w*2 == o && f*2 == o && f%2 == 0) || (w%2 != 0 && w*2 == o && f*2 == o && f%2 != 0) {
		fmt.Println(w)
	} else if w == 0 || o == 0 || f == 0 {
		fmt.Println(0)
	} else if o%2 != 0 {
		o = o - 1
		if f < w {
			if o >= f {
				fmt.Println(f)
			}
		} else {
			if o >= w {
				fmt.Println(w)
			}
		}
	}
}
