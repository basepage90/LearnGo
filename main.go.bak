package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/woojebiz/learngo/accounts"
	"github.com/woojebiz/learngo/mydict"
)

// day1 "github.com/woojebiz/learngo/codeDiv"
// day2 "github.com/woojebiz/learngo/codeDiv"
// day3 "github.com/woojebiz/learngo/codeDiv"

var errorRequestFailed = errors.New("Request failed")
var errorNoneResp = errors.New("None Response")

type reqResult struct {
	url    string
	status string
}

func main() {

	startTime := time.Now()

	results := make(map[string]string)
	c := make(chan reqResult)

	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"https://www.amazon.com",
		"https://www.stackoverflow.com",
		"http://hiphople.co.kr",
		"https://www.soundcloud.com",
		"https://httpstat.us/404",
		"https://httpstat.us/400",
		"https://instagram.com",
		"https://www.twitch.tv",
	}

	// goroutine : go에서 사용하는 동시성을위한 동작으로 Thread 와는 다르다
	// 			   OS가아닌 Go runtime에서 cooperative scheduling하므로로, overhead나  context swtiching이 Thread에 비해서 적다.
	// goroutine 은 main이 종료되기전까지만 살아있을 수 있다.
	for _, url := range urls {
		go hitURL(url, c)
	}

	// channel : 일종의 pipe
	// result <- c  : channel 값을 다 받아오기전에는 main이 종료 되지않는다.
	// channel 의 갯수가  초과해서는 안된다
	for range urls {
		result := <-c
		results[result.url] = result.status
	}

	// map 출력
	for url, status := range results {
		fmt.Println(url, status)
	}

	elapsed := time.Since(startTime)
	fmt.Println("elapsed", elapsed)

}

// channel send only
// input variable의 chan 에 direction을 지정해주면, 채널 전달 방향을 지정할 수 있다.
//  ex : func pingpong(ping <-chan stirng, pong chan<- string) {
//			msg := <- ping
//			pong <- msg
//       }
func hitURL(url string, c chan reqResult) {
	fmt.Println("Checking URL : ", url)
	resp, err := http.Get(url)
	status := "OK"
	if resp == nil {
		status = "FAILED : " + "None Response"
	} else if err != nil || resp.StatusCode >= 400 {
		status = "FAILED : " + resp.Status
	}

	c <- reqResult{url: url, status: status}
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
