package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/woojebiz/learngo/accounts"
	"github.com/woojebiz/learngo/mydict"
)

// day1 "github.com/woojebiz/learngo/codeDiv"
// day2 "github.com/woojebiz/learngo/codeDiv"
// day3 "github.com/woojebiz/learngo/codeDiv"

var errorRequestFailed = errors.New("Request failed")
var errorNoneResp = errors.New("None Response")

func main() {

	var results = make(map[string]string)

	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.amazon.com",
		"https://www.stackoverflow.com",
		"http://hiphople.co.kr",
		"https://www.soundcloud.com",
		"https://httpstat.us/404",
		"https://httpstat.us/400",
	}

	for _, url := range urls {

		result := "OK"
		err, resp := hitURL(url)

		if err != nil {
			result = "FAILED :" + err.Error() + "RESPONSE :" + resp
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

}

func hitURL(url string) (error, string) {
	fmt.Println("Checking :", url)

	resp, err := http.Get(url)

	if resp == nil {
		return errorNoneResp, ""
	} else if err != nil || resp.StatusCode >= 400 {
		return errorRequestFailed, resp.Status
	}
	return nil, ""
}

func main_asis() {

	/*******************************************************************************************/
	/***********************************       Theory        ***********************************/
	/*******************************************************************************************/
	// day1.Day1()
	// day2.Day2()
	// day3.Day3()

	/*******************************************************************************************/
	/*********************************** mini project : bank ***********************************/
	/*******************************************************************************************/

	fmt.Println("*********************************** mini project : bank ***********************************")

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

	/*******************************************************************************************/
	/********************************* mini project : dictionary *******************************/
	/*******************************************************************************************/

	fmt.Println("********************************* mini project : dictionary *******************************")

	dictionary := mydict.Dictionary{"first": "firstWord"}

	word := "secound"
	def := "secoundWord"

	fmt.Println("-- Add")

	err1 := dictionary.Add(word, def)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println("-- Search")
	// definition, err2 := dictionary.Search("first")
	definition, err2 := dictionary.Search("secound")
	// definition, err2 := dictionary.Search("third")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}

	fmt.Println("-- Update")
	err3 := dictionary.Update("secound", "New SecoundWord")

	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println("-- Search")
	definition, err2 = dictionary.Search("secound")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}

	fmt.Println("-- Delete")

	err4 := dictionary.Delete("secound")

	if err4 != nil {
		fmt.Println(err4)
	}

	fmt.Println("-- Search")
	definition, err2 = dictionary.Search("secoucnd")

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(definition)
	}

}
