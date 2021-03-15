package codeDiv

import "fmt"

func canIDrink(age int) (int, bool) {

	if koreanAge := age + 1; koreanAge < 18 {
		return age, false
	} else {
		return age, true
	}
}

func canIDrink_switch(age int) bool {
	switch koreanAge := age + 1; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return false
}

func Day2() {
	/********** 04. if & switch **********************************
	 * go의 특별한 점은 if 문에서 변수를 선언 할수 있다는 것이다.
	 * 이는 if 문만을 위한 변수이며, 명시적인 장점을 가진다
	 ************************************************************/
	fmt.Println("\n --  04.  if")

	age, res := canIDrink(16)
	fmt.Println("미국나이", age, "세 는 한국에서는 술을 마실 수", res)

	fmt.Println(canIDrink_switch(17))

}
