package codeDiv

import "fmt"

type person struct {
	name    string
	age     int
	gender  string
	favFood []string
}

func Day3() {
	/********** 05. pointer *************************************
	 * 앰퍼샌드(&) 		: 주소를 사용
	 * 애스터리스크(*) 	: 주소에 담긴 값을 사용
	 ************************************************************/
	fmt.Println("\n -- 05. pointer")

	a := 2
	b := &a
	a = 10
	fmt.Println(a, *b) // a의 주소를 공유했기때문에, *b는 2가아닌 10을 가져온다 (b의 주소에 담긴 값 즉, a의 주소에 담긴값을 가져온다.)

	*b = 20
	fmt.Println(a, *b) // b의 주소와 a의 주소는 연결된 상태이므로, b의 주소에 담긴값은  a의 주소에 담긴 값과 동일하다.

	/********** 06. array and slices ***************************
	 * slices 란 ? array 에서 size를 명시해주지않으면, 가변 array로 사용가능
	 * slices 는 append 를 통해 값을 추가가능하며, append는 java와 는 다르게 slices를 return 한다.
	 * ex) slices = append(slices, value)
	 ************************************************************/
	fmt.Println("\n -- 06. array and slices")

	art := []string{"dali", "van", "picaso"}
	fmt.Println(append(art, "artist")) // append 는 기존 slice 값에 return 만을 할뿐이다
	fmt.Println(art)

	art = append(art, "beenzino") // sliece에 덮어 씌워줘야, 실질적인 append가 이루어진다.
	fmt.Println(art)

	/********** 07. map ****************************************
	 * mapName := map[keyType]valType{key:val}
	 * get 해오는 방법 - mapName[key]
	************************************************************/
	fmt.Println("\n -- 07. map")

	//  map 선언과 동시에 초기화
	beenzino := map[string]string{"name": "beenzino", "job": "rapper", "age": "34"}
	fmt.Println(beenzino)

	fmt.Println("-------------------------------------------------------------------")

	//  map 선언후 생성
	dprlive := map[string]string{}
	dprlive["name"] = "dprilive"
	dprlive["age"] = "28"
	dprlive["song01"] = "jasmine"
	dprlive["song02"] = "laputa"
	fmt.Println(dprlive)

	// key 로 value 가져오기
	fmt.Println(dprlive["name"])
	fmt.Println(dprlive["album"]) // key가 없어도 error를 발생하지는 않는다.
	fmt.Println("-------------------------------------------------------------------")

	// range 도 사용가능하다
	for key, value := range dprlive {
		fmt.Println(key, value)
	}

	/********** 08. struct ****************************************
	 * 구조체
	 * for Custom Data Type Expression
	 * 필드들의 집합체, 필드들의 컨테이너
	 * 필드 데이타만을 가지며, 메서드를 갖지 않는다 (스트럭쳐와 메서드의 분리)
	 *** go 는 class obejct 가 없다. 오직 struct 로 대신한다
	 ***
	************************************************************/
	fmt.Println("\n -- 08. struct")

	favFood := []string{"corn", "beef", "coke"}

	// 선언과 동시에 초기화
	// kendrick := person{name: "kendrick lamar", age: 33, gender: "M", favFood: favFood}

	// 선언 후 초기화
	kendrick := person{}
	kendrick.name = "kendrick lamar"
	kendrick.age = 33
	kendrick.gender = "M"
	kendrick.favFood = favFood

	fmt.Println(kendrick)

}
