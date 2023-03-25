package main

import (
	"fmt"
	"math"
)

//
//type AutoGenerated struct {
//	ID       int    `jsonning:"id"`
//	Code     string `jsonning:"Code"`
//	Ccy      string `jsonning:"Ccy"`
//	CcyNmRU  string `jsonning:"CcyNm_RU"`
//	CcyNmUZ  string `jsonning:"CcyNm_UZ"`
//	CcyNmUZC string `jsonning:"CcyNm_UZC"`
//	CcyNmEN  string `jsonning:"CcyNm_EN"`
//	Nominal  string `jsonning:"Nominal"`
//	Rate     string `jsonning:"Rate"`
//	Diff     string `jsonning:"Diff"`
//	Date     string `jsonning:"Date"`
//}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panicking! disc<0")

		}
	}()
	fmt.Println("a*x^2+b*x+c")
	var (
		a, b, c int
		x1, x2  float64
	)
	fmt.Print("a ni kiriting: ")
	fmt.Scan(&a)
	fmt.Print("b ni kiriting: ")
	fmt.Scan(&b)
	fmt.Print("c ni kiriting: ")
	fmt.Scan(&c)
	if a%2 == 0 && b%2 == 0 && c%2 == 0 {
		a = a / 2
		b = b / 2
		c = c / 2
	}
	d := b*b - 4*a*c
	fmt.Println("Diskriminant = ", d)
	if d > 0 {
		x1 = (-float64(b) + math.Sqrt(float64(d))) / 2 * float64(a)
		x2 = (-float64(b) - math.Sqrt(float64(d))) / 2 * float64(a)
		fmt.Printf(`x1-%v,
x2-%v`, x1, x2)

	} else if d == 0 {
		x1 = (-float64(b) + math.Sqrt(float64(d))) / 2 * float64(a)
		fmt.Println("x1", x1)
	} else {
		panic(fmt.Sprintf("%v", d))

	}
}

//
//data, err := os.ReadFile("jsonning.jsonning")
//if err != nil {
//	panic(err)
//}
//var Js []AutoGenerated
//err = jsonning.Unmarshal(data, &Js)
//if err != nil {
//	fmt.Println("Error unmarshaling JSON:", err)
//	return
//}
//var val string
//fmt.Print("Valyutani kiriting: ")
//fmt.Scan(&val)
//for _, Js := range Js {
//	if val == Js.Ccy {
//		fmt.Println(Js)
//	}
//}
//
//var b = [...]int64{1, 2, 3, 4, 5, 5}
//fmt.Println(b)