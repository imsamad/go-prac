package main
import ("fmt")

// tmp:= 1 // not allowed

var tmp int = 1

// var vs :=
// var can be use anywhere locally or gloally where := not is scope
// using var variable can be init after declaration wherease := does not allow it

func main() {
	fmt.Println(tmp)
	// unlike rust, can be decalared/defined anywhere in any scope
	const (
		A int = 1
		B = 3.14
		C = "Hi!"
	  )
		fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)

  var (
	a int
	b int = 1
	c string = "hello"
  )

 fmt.Println(a)
 fmt.Println(b)
 fmt.Println(c)
}