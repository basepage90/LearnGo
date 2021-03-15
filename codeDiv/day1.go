package codeDiv

import (
	"fmt"
	"strings"

	something "github.com/woojebiz/learngo/codeDiv/something"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// naekd return : reutrn 변수에 대한 명시가 없어도, 함수 변수로 선언한 값을 선언했던 type 으로  return 해준다
func substrAndUpper(indata string) (substr string, uppercase string) {
	defer fmt.Println("working done : this is defer") // defer 는 return 직후 실해된다. callback 역활을 하는듯?

	substr = indata[0:10]
	uppercase = strings.ToUpper(substr)
	return
}

func superAdd(numbers ...int) int {

	// range 표현
	for index, number := range numbers {
		fmt.Println("index :", index, "number :", number)
	}

	// 기본표현
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index :", i, "number :", numbers[i])
	}

	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}

func Day1() {

	fmt.Println("Hello World!!!")

	/********** 00. Package ******************
	*
	* import 된 package의 func 의 첫알파벳의 대소문자의 차이
	* : 대문자로 시작하면 자바의 public /  소문자라면 자바의 defualt (같은 패키지내) 로 비교 할 수 있다.
	*****************************************/
	fmt.Println("\n-- 00. Package")

	something.SayHello()
	// something.sayBye()	// 소문자이므로 같은 패키지에서만 실행가능

	/********** 01. const & var  *************
	 *
	 * const : 상수 , 선언후 변경불가
	 * var   : 변수 , 선언후 변경가능
	 *****************************************/
	fmt.Println("\n-- 00. const & var")

	/* cosnt*/
	const test string = "origin"
	// test = "trans"

	/* var */
	var name string = "riven"
	fmt.Println(name + " is rabbit :)")

	name = "zed"
	fmt.Println(name + " is not rabbit :(")

	/* var simple expression :  auto type match !!!
	 * 단, 축약형은 func 내부에서만 사용가능하며, 선언 시점의 타입으로 정해진다.
	 */
	newName := "yasuo"
	fmt.Println(newName)

	newName = "talon"
	fmt.Println(newName)

	/********** 02.func ***********************
	 * go의 특징으로 return 값을 여러개사용가능
	 * underscore(_) : ignored value
	 * naked return  : 반환 값을 명시해주지않아도 go가 처리해 준다
	 * defer 		 : return 후의 시점에 실행된다
	 ******************************************/
	fmt.Println("\n-- 02. func")

	fmt.Println(multiply(3, 5))

	fmt.Println(lenAndUpper("steven"))

	// multi return
	len, upperName := lenAndUpper("steven")
	fmt.Println(len, upperName)

	// return ignore
	len2, _ := lenAndUpper("steven")
	fmt.Println(len2)

	// ...string
	repeatMe("I", "understand", "this", "logic")

	/* naked return */
	fmt.Println(substrAndUpper("Look at me now"))

	/********** 03. Loop ***********************
	 * only for
	 ******************************************/
	fmt.Println("\n-- 03. Loop : ")

	fmt.Println(superAdd(1, 2, 3, 4, 5))
}
