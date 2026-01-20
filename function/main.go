// package main /// main program,program bnane ke liye use hota hai

// import "fmt" // print krne ke liye

// func simpleFunction(){
// 	fmt.Println("simple function")
// }

// func add(a,b int) int {
// 	return  a + b

// }

// func multiply(a float32,b int) (result float32){
// 	result = a * float32(b) 
// 	return
// }

// func main(){  // execution yhi se start hota hai
//    fmt.Println("we are learning go lang")
//    simpleFunction()

//    ans:= add(3,4)
//    fmt.Println("add of two number is :",ans)


//    store:=multiply(4.65,3)
//    fmt.Println("multiply of two numbers",store)


// }

//types of function 

// 1 vardiac function

package main

import "fmt"

func sums(nums...int) int{
	total:= 0

	for _,num:=range nums{
		total = total + num
	}

	return total
}

func main(){
	nums:=[]int{3,7,5,4}
	result:= sums(nums...)
	fmt.Println(result)

}

// 2 

