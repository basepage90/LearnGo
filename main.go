package main

import (
	"fmt"

	"github.com/woojebiz/learngo/accounts"
)

// day1 "github.com/woojebiz/learngo/codeDiv"
// day2 "github.com/woojebiz/learngo/codeDiv"
// day3 "github.com/woojebiz/learngo/codeDiv"

func main() {

	/*******************************************************************************************/
	/***********************************       Theory        ***********************************/
	/*******************************************************************************************/
	// day1.Day1()
	// day2.Day2()
	// day3.Day3()

	/*******************************************************************************************/
	/*********************************** mini project : bank ***********************************/
	/*******************************************************************************************/

	/******* bank 01 ***************************************************************************
	 * struct와 variable 의 public/private 한 사용
	 * func으로 constructor 를 대신하기
	 *******************************************************************************************/
	account := accounts.NewAccount("kanye", "US") //  func를 통해서 construct 역활을 대신하였다.
	// account.owner = "dean"
	// account.balance = 1000
	account.Address = "FR" // public 변수인 Address 는 변경이 가능하다

	fmt.Println(account)

	/******* bank 02 ***************************************************************************
	 * method
	 * func 에 receiver를 붙여주면, struct 의 method 가 된다
	 *** value receiver   : struct를 복사한다 ex) func (t Test) MethodName () {}
	 *** pointer receiver : struct를 참조한다 ex) func (t *Test) MethodName () {}
	 *******************************************************************************************/
	account.Deposit(10)

	fmt.Println(account)
	fmt.Println(account.Balance())

	/******* bank 03 ***************************************************************************
	 * exception
	 * exception 시 error 처리를 별도로 해줘야하며, exception 시에도 이후 동작을 멈추지않는다.
	 *******************************************************************************************/
	err := account.Withdraw(20)

	if err != nil { // nil 은 null 로 생각 할 수 있다.
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}
