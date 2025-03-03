package main

// 중요하게 살펴봐야할건 ⭐️로 표시해 두었으니 검색해서 찾아보면 된다.
// var g int = 10

/* 1.Hello Go World
fmt.Println("Hello Go World Test")
*/

/* 2,3.Text Input/output using fmt and variables

var a int = 10
var msg string = "Hello Variable"

a = 20
msg = "Good Morning"
fmt.Println(msg, a)

var minimumMage int = 10
var workingHour = 20

var income int = minimumMage * workingHour

fmt.Println(minimumMage, workingHour, income)

var a int = 3
var b int
var c = 4
d := 5

fmt.Println(a, b, c, d)
a := 3
var b float64 = 3.5

var c int = int(b)

d := float64(a * c)

var e int64 = 7
f := int64(d) * e

var g int = int(b * 3)
var h int = int(b) * 3
fmt.Println(g, h, f)


var a int16 = 3456
var c int8 = int8(a)

fmt.Println(a)
fmt.Println(c)


var m int = 20

{
	var s int = 50
	fmt.Println(m, s, g)
}

m = s + 20


var a float32 = 1234.523
var b float32 = 3456.123
var c float32 = a * b
var d float32 = c * 3

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)


var a int = 10
var b int = 20
var f float64 = 32799438743.8297
fmt.Print("a:", a, "b:", b)
fmt.Println("a:", a, "b:", b, "f:", f)
fmt.Printf("a: %d b: %d f:%f¥n", a, b, f)


var a = 123
var b = 456
var c = 123456789
fmt.Printf("%5d, %5d\n", a, b)
fmt.Printf("%05d, %05d\n", a, b)
fmt.Printf("%-5d, %-05d\n", a, b)

fmt.Printf("%5d, %5d\n", c, c)
fmt.Printf("%05d, %05d\n", c, c)
fmt.Printf("%-5d, %-05d\n", c, c)


var a = 324.13455
var c = 3.14

fmt.Printf("%08.2f\n", a)
fmt.Printf("%08.2g\n", a)
fmt.Printf("%08.5g\n", a)
fmt.Printf("%f\n", c)


str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

fmt.Print(str)
fmt.Printf("%s", str)
fmt.Printf("%q", str)


var a int
var b int
n, err := fmt.Scan(&a, &b)

if err != nil {
	fmt.Println(n, err)
} else {
	fmt.Println(n, a, b)
}

var a int
var b int

n, err := fmt.Scanf("%d %d\n", &a, &b)
if err != nil {
	fmt.Println(n, err)
} else {
	fmt.Println(n, a, b)
}

var a int
var b int

n, err := fmt.Scanln(&a, &b)
if err != nil {
	fmt.Println(n, err)
} else {
	fmt.Println(n, a, b)
}


stdin := bufio.NewReader(os.Stdin)

var a int
var b int

n, err := fmt.Scanln(&a, &b)

if err != nil {
	fmt.Println(err)
	stdin.ReadString('\n')
} else {
	fmt.Println(n, a, b)
}

n, err = fmt.Scanln(&a, &b)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(n, a, b)
}
*/

/* 4.operator
// Result type of operation
var x int32 = 7
var y int32 = 3

var s float32 = 3.14
var t float32 = 5

fmt.Println("x + y = ", x+y)
fmt.Println("x - y = ", x-y)
fmt.Println("x * y = ", x*y)
fmt.Println("x / y = ", x/y)
fmt.Println("x % y = ", x%y)

fmt.Println("s * t = ", s*t)
fmt.Println("s / t = ", s/t)

// Bit operation
var x1 int8 = 34
var x2 int16 = 34
var x3 uint8 = 34
var x4 uint16 = 34

fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(x1))
fmt.Printf("^%d = %5d, \t %016b\n", x2, ^x2, uint16(x2))
fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
fmt.Printf("^%d = %5d, \t %016b\n", x4, ^x4, ^x4)

// shift operator
var x int8 = 4
var y int8 = 64

fmt.Printf("x:%08b x<<2: %08b x<<2: %d\n", x, x<<2, x<<2)
fmt.Printf("y:%08b y<<2: %08b y<<2: %d\n", y, y<<2, y<<2)

var x int8 = 16
var y int8 = -126
var z int8 = -1
var w uint8 = 128

fmt.Printf("x:%08b x>>2: %08b x>>2: %d\n", x, x>>2, x>>2)
fmt.Printf("y:%08b y>>2: %08b y>>2: %d\n", uint8(y), uint8(y>>2), y>>2)
fmt.Printf("z:%08b z>>2: %08b z>>2: %d\n", uint8(z), uint8(z>>2), z>>2)
fmt.Printf("w:%08b w>>2: %08b w>>2: %d\n", uint8(w), uint8(w>>2), w>>2)

// interger overflow
var x int8 = 127

fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
fmt.Printf("x\t= %4d, %08b\n", x, x)
fmt.Printf("x + 1\t=%4d, %08b\n", x+1, x+1)
fmt.Printf("x + 2\t=%4d, %08b\n", x+2, x+2)
fmt.Printf("x + 3\t=%4d, %08b\n", x+3, x+3)

var y int8 = -128
fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
fmt.Printf("y\t= %4d, %08b\n", y, y)
fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)

// float 비교연산,
var a float64 = 0.1
var b float64 = 0.2
var c float64 = 0.3

fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
fmt.Println(a + b)

// ⭐️ignoring small error and float error
const epsilon = 0.000001

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}

}
var a float64 = 0.1
var b float64 = 0.2
var c float64 = 0.3

fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

a = 0.0000000000004
b = 0.0000000000002
c = 0.0000000000007

fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

//⭐️Use Nextafter to ignore float error and small error
// In GoLang, the math.NextAfter function lets you fine-tune(세부조정하다) the floating-point(부동 소수점) values.
// This function returns the smallest floating point value close to the first argument in the direction given by(~의해 결정되다) the second argument.
// This is useful if you want to make a very small difference in floating-point operations, or move to a slightly(약간) larger/small value in any direction.
func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

//⭐️Use Float objects provided by the math/big package
a, _ := new(big.Float).SetString("0.1")
b, _ := new(big.Float).SetString("0.2")
c, _ := new(big.Float).SetString("0.3")

d := new(big.Float).Add(a, b)
fmt.Println(a, b, c, d)
fmt.Println(c.Cmp(d))

// Multiple substitution(대입) operators
var a int = 10
var b int = 20

a, b = b, a
fmt.Println(a, b)
*/

/* 5.function

func Add(a int, b int) int {
	return a + b
}

//⭐️Definition
// Arguments are copied as parameters.
// Variables declared within parameters and functions cannot be accessed outside the range of variables when the function is terminated.
// Call by value
// Call by reference, use & to access the point address.(php)
// Golang is not Call by refence
c := Add(3, 6)
fmt.Println(c)

// Why Use a Function?
func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}
PrintAvgScore("김일등", 80, 74, 95)
PrintAvgScore("송이등", 88, 92, 53)
PrintAvgScore("박삼등", 78, 73, 78)

// Multi-return function
//⭐️The front bracket represents the parameter, and the back bracket represents the return type.
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
c, success := Divide(9, 3)
fmt.Println(c, success)
d, success := Divide(9, 0)
fmt.Println(d, success)

// Return with variable name
//⭐️When specifying(지정) variable names, you must specify all variable names
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)

//recursive function(재귀함수)
func printNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)
}
printNo(3)
*/

/* 6.constant
//Declare
const C int = 10

var b int = C * 20
C = 20
fmt.Println(&C)

// Use a constant for a value that must not change
const PI1 float64 = 3.141592653589793238
var PI2 float64 = 3.141592653589793238

PI2 = 4
fmt.Printf("circumferential rate(원주율): %f\n", PI1)
fmt.Printf("circumferential rate(원주율): %f\n", PI2)

// Use as code value
func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("Pig")
	} else if animal == Cow {
		fmt.Println("Cow")
	} else if animal == Chicken {
		fmt.Println("Chicken")
	} else {
		fmt.Println("...")
	}
}
const Pig int = 0
const Cow int = 1
const Chicken int = 2

PrintAnimal(Pig)
PrintAnimal(Cow)
PrintAnimal(Chicken)
PrintAnimal(0)
PrintAnimal(1)
PrintAnimal(2)
PrintAnimal(3)

// Easy to use enumerated values with iota
const (
	Red   int = iota //0
	Blue  int = iota //1
	Green int = iota //2
)

const (
	C1 uint = iota + 1 // 1 = 0 + 1
	C2                 // 2 = 1 + 1
	C3                 // 3 = 2 + 1
)

const (
	BitFlag1 uint = 1 << iota // 1 = 1<<0
	BitFlag2                  // 2 = 1<<1
	BitFlag3                  // 4 = 1<<2
	BitFlag4                  // 8 = 1<<2
)

const (
	A int = iota // 0
	B            // 1
)

// Typeless Constant
const PI = 3.14
const FloatPI float64 = 3.14

var a int = PI * 100
var b int = int(FloatPI) * 100 // 형변환 자체가 막혀버림

fmt.Println(a)
fmt.Println(b)

//⭐️ 상수와 리터럴
// 고정값 자체를 리터럴이라고 함.
// 상수는 컴파일 될때는 리터럴이랑 같이 취급
// 상수 표현식 역시 컴파일 타임에 실제 결과값의 리터럴로 변환하기 때문에 상수 표현식 계산에
// CPU자원을 사용하지 않음
// 상수의 메모리 주소값에 접근 할 수 없는 이유 역시 컴파일 타임에
// 리터럴로 전환되어서 실행 파일에 값 형태로 쓰인다. 그래서 동적 할당 메모리 영역을 사용하지 않음
var str string = "Hello World"
var i int = 0
i = 30

*/

/* 7.if
// Basic Usage
light := "red"

if light == "green" {
	fmt.Println("Pass")
} else {
	fmt.Println("Not Pass")
}

temp := 33

if temp > 28 {
	fmt.Println("Aircon")
} else if temp <= 3 {
	fmt.Println("Heater")
} else {
	fmt.Println("Stand by")
}

var age = 22
if age >= 10 && age <= 15 {
	fmt.Println("You are Young")
} else if age > 30 || age < 20 {
	fmt.Println("You are not 20s")
} else {
	fmt.Println("Best age of your life")
}

//⭐️쇼트 서킷
// &&연산은 좌변이 펄스이면 우변을 검사하지 않음. ||연산도 좌변이 트루이면 우변을 검사하지않는데 이걸 쇼트 서킷이라고 부른다.
var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn", cnt)
	cnt++
	return cnt
}
if false && IncreaseAndReturn() < 5 {
	fmt.Println("1 increase")
}

if true || IncreaseAndReturn() < 5 {
	fmt.Println("2 increase")
}

fmt.Println("cnt:", cnt)

// Overlapping(중첩) if
func HasRichFriend() bool {
	return true
}

func GetFriendsCount() int {
	return 3
}
price := 35000

if price > 50000 {
	if HasRichFriend() {
		fmt.Println("HasRichFriend")
	} else {
		fmt.Println("divide up")
	}
} else if price >= 30000 && price <= 50000 {
	if GetFriendsCount() > 3 {
		fmt.Println("GetFriendsCount")
	} else {
		fmt.Println("divide up")
	}
} else {
	fmt.Println("All Price")
}

//⭐️if 초기문; 조건문
func getMyAge() (int, bool) {
	return 33, true
}
if age, ok := getMyAge(); ok && age < 20 {
	fmt.Println("You are Young", age)
} else if ok && age < 30 {
	fmt.Println("Nice age", age)
} else if ok {
	fmt.Println("You are beautiful", age)
} else {
	fmt.Println("Error")
}

// fmt.Println("Your age is", age)
*/

/* 8.switch
// principle of operation
a := 3

switch a {
case 1:
	fmt.Println("a == 1")

case 2:
	fmt.Println("a == 2")

case 3:
	fmt.Println("a == 3")

case 4:
	fmt.Println("a == 4")

default:
	fmt.Println("a > 2")
}

// When to use
var day = 3

if day == 1 {
	fmt.Println("day 1")
} else if day == 2 {
	fmt.Println("day 2")
} else if day == 3 {
	fmt.Println("day 3")
} else if day == 4 {
	fmt.Println("day 4")
} else if day == 5 {
	fmt.Println("day 5")
} else {
	fmt.Println("defalt")
}

// 위와 같이 조건문이 지저분할때 사용
day = 4
switch day {
case 1:
	fmt.Println("day 1")
case 2:
	fmt.Println("day 2")
case 3:
	fmt.Println("day 3")
case 4:
	fmt.Println("day 4")
default:
	fmt.Println("day 1")
}

// 한번에 여러값 비교
day = 1

switch day {
case 1, 2:
	fmt.Println("day 1 or 2")
case 3, 4, 5:
	fmt.Println("day 3 or 4 or 5")
}

//⭐️switch 초기문
func getMyAge() int {
	return 22
}

switch age := getMyAge(); age {
case 10:
	fmt.Println("Teenage")
case 33:
	fmt.Println("Pair 3")
default:
	fmt.Println("My age is", age)
}

switch age := getMyAge(); true {
case age < 10:
	fmt.Println("Child")
case age < 20:
	fmt.Println("Teenager")
case age < 30:
	fmt.Println("20s")
default:
	fmt.Println("My age is", age)
}
		// const열거값과 switch
type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return Blue
}
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))

// ⭐️break와 fallthrough 키워드
// go는 브레이크를 사용하던 사용하지 않던 케이스 하나를 실행 후 스위치문을 빠져나감
// fallthrough를 사용하자 case 4까지 실행됨
// 사용하지 않는것을 권장함
a := 3
switch a {
case 1:
	fmt.Println("a == 1")
	break
case 2:
	fmt.Println("a == 2")
case 3:
	fmt.Println("a == 3")
	fallthrough
case 4:
	fmt.Println("a == 4")
default:
	fmt.Println("a > 4")
}
*/

/* 9.for문
// 기본
for i := 0; i < 10; i++ {
	fmt.Print(i, ",")
}
//⭐️초기문, 후처리 생략이 가능하고 조건문만 있는 경우도 있으며 아예 모든 것을 생략하면
// 무한 루프가 된다
// 초기문 생략 -> for;조건문;후처리
// 초기문 생략 -> for;조건문;후처리
// 조건문만 있는 경우 -> for ;조건문
// 조건문만 있는 경우 -> for 조건문
// 무한루프 -> for true {} for {}

// 정수값 순회
for i := range 10 {
	fmt.Println(i)
}

//continue와 break
stdin := bufio.NewReader(os.Stdin)
for {
	fmt.Println("enter")
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("enter the number")

		stdin.ReadString('\n')
		continue
	}
	fmt.Printf("The number entered is %d.\n", number)
	if number%2 == 0 {
		break
	}
}
fmt.Println("The for loop is done.")

// 중첩 for문
for i := 0; i < 3; i++ {
	for j := 0; j < 5; j++ {
		fmt.Print("*")
	}
	fmt.Println()
}
for i := 0; i < 5; i++ {
	for j := 0; j < i+1; j++ {
		fmt.Print("*")
	}
	fmt.Println()
}

dan := 2
b := 1
for {
	for {
		fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
		b++
		if b == 10 {
			break
		}
	}
	b = 1
	dan++
	if dan == 10 {
		break
	}
}
fmt.Println("The for loop is done.")

// 중첩 for문과 break,
a := 1
b := 1
found := false
for ; a <= 9; a++ {
	for b = 1; b <= 9; b++ {
		if a*b == 45 {
			found = true
			break
		}
	}
	if found {
		break
	}
}

fmt.Printf("%d * %d = %d\n", a, b, a*b)
// ⭐️레이블
// 근데 플래그 변수나 레이블은 꼭 필요할때만 사용하고 잘 함수화해서 스파게티 코드가 되지않도록 방지해야한다
OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}

	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

	// 함수화를 이용한 중첩 for문과 플래그변수, 레이블을 사용하지 않는 방법
func find45(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}
a := 1
b := 1

for ; a <= 9; a++ {
	var found bool
	if b, found = find45(a); found {
		break
	}
}
fmt.Printf("%d * %d = %d\n", a, b, a*b)
*/

/* 10.배열
		//기본
			var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

			for i := 0; i < 5; i++ {
				fmt.Println(t[i])
			}

			// 배열 선언시 개수는 항상 상수
			const Y int = 3
			x := 5
			a := [x]int{1, 2, 3, 4, 5}

			b := [Y]int{1, 2, 3}

			var c [10]int
			// 배열 요소 읽고 쓰기
	//⭐️...으로 배열의 개수를 생략할 수 있다(...은 가변인수확장 가변인수 확장에 대해서는 슬라이스 참조)
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	//⭐️range 순회
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		// 선언 대입문을 사용하여 i에는 인덱스 v에는 요소값이 들어간다
		fmt.Println(i, v)
	}

	for _, v := range t {
		// _를 이용하여 인덱스 생략이 가능하다 v에는 요소값이 들어간다
		fmt.Println(v)
	}
		//⭐️배열 복사
a := [5]int{1, 2, 3, 4, 5}
b := [5]int{500, 400, 300, 200, 100}

for i, v := range a {
	fmt.Printf("a[%d] = %d\n", i, v)
}

fmt.Println()
for i, v := range b {
	fmt.Printf("b[%d] = %d\n", i, v)
}

//대입 연산자를 사용해서 배열을 복사 타입이 일치하지 않으면 에러 발생
b = a
fmt.Println()
for i, v := range b {
	fmt.Printf("b[%d] = %d\n", i, v)
}

// 이중 배열
a := [2][5]int{
	{1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10}, // 닫는 중괄호가 마지막에 같이 있지 않으면 쉼표찍어줘야함
}
// 이건 안찍어도 됨
b := [2][5]int{
	{1, 2, 3, 4, 5},
	{6, 7, 8, 9, 10}}
for _, arr := range a {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}

for _, arr := range b {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
*/

/* 11.구조체
//기본 사용법
type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	//초기값 생략
	var house House
	// 모든 필드 초기화 방법
	////var house House = House{"", 0, 0, ""}
	// 모든 필드 초기화 방법 여러줄로 초기화 할 때
	// var house House = House{
	// 	"",
	// 	0,
	// 	0,
	// 	"",
	// } //⭐️여러줄로 초기화 할때는 마지막에 쉼표를 빼먹어선 안된다
	//일부 필드만 초기화 할때는 필드명:필드값으로 초기화

	house.Address = "tokyo"
	house.Size = 28
	house.Price = 9.8
	house.Type = "apartment"

	fmt.Println(house.Address)
	fmt.Println(house.Size)
	fmt.Println(house.Price)
	fmt.Println(house.Type)
}
	// ⭐️구조체를 포함하는 구조체
type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"name", "1", 35}
	vip := VIPUser{
		User{"VIPName", "2", 35},
		99,
		2500,
	}

	fmt.Printf("User: %s ID: %s age: %d\n", user.Name, user.Id, user.Age)
	fmt.Printf("VIPUser: %s ID: %s age: %dVIPLevel: %d VIPPrice: %d\n ",
		vip.UserInfo.Name,
		vip.UserInfo.Id,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
// 포함된 필드 방식
type User struct {
	Name string
	Id   string
	Age  int
}

type VIPUser struct {
	User     // 필드명을 생략
	VIPLevel int
	Price    int
}

func main() {
	user := User{"name", "1", 35}
	vip := VIPUser{
		User{"VIPName", "2", 35},
		99,
		2500,
	}

	fmt.Printf("User: %s ID: %s age: %d\n", user.Name, user.Id, user.Age)
	fmt.Printf("VIPUser: %s ID: %s age: %dVIPLevel: %d VIPPrice: %d\n ",
		vip.Name,	//⭐️ . 하나로 접근이 가능(포함된 필드)
		vip.Id,
		vip.Age,
		vip.VIPLevel,
		vip.Price,
	)
}
// 필드 중복해결
type User struct {
	Name  string
	Id    string
	Age   int
	Level int
}

type VIPUser struct {
	User
	Price int
	Level int
}

func main() {
	user := User{"name", "nameId", 23, 10}
	vip := VIPUser{
		User{"VIPName", "VIPNameId", 40, 10},
		250,
		3,
	}

	fmt.Printf("User: %s ID: %s age: %d\n", user.Name, user.Id, user.Age)
	fmt.Printf("VIPUser: %s ID: %s age: %d Level: %d VIPLevel: %d\n",
		vip.Name,       // 임베딩된 User의 Name 필드에 접근
		vip.Id,         // 임베딩된 User의 Id 필드에 접근
		vip.Age,        // 임베딩된 User의 Age 필드에 접근
		vip.Level,      // VIPUser의 Level 필드에 접근
		vip.User.Level, // ⭐️User 구조체의 Level 필드에 접근
	)
}
// 구조체 값 복사
type Student struct {
	Age   int
	No    int
	Score float64
}

func PrintStudent(s Student) {
	fmt.Printf("Age:%d No:%d Score:%.2f\n", s.Age, s.No, s.Score)
}

func main() {
	var student = Student{15, 23, 88.2} // 구조체 초기화는 항상 중괄호

	student2 := student

	PrintStudent(student2)
}
//⭐️메모리관리에 관해서 한번 더 보고싶으면 북마크 참조
*/

/* ⭐️12.포인터
//기본사용법
func main() {
	var a int = 500
	var p *int = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}
// 포인터 변수값 비교하기
func main() {
	var a int = 10
	var b int = 20
	var p1 *int = &a
	var p2 *int = &a
	var p3 *int = &b

	fmt.Printf("p1 == p2 : %v\n", p1 == p2)
	fmt.Printf("p2 == p3 : %v\n", p2 == p3)
}

// 포인터의 기본값 nil
// var p *int
// if p!= nil {
// 	// p가 널이 아니라는 얘기는 p가 유효한 메모리 주소를 가리킨다는 뜻
// }


// 포인터를 사용하는 이유
type Data struct {
	value int
	data  [200]int
}

// 매개변수는 값에 의한 전달이므로 arg와 밑에서 선언한 data는 다른변수다
//
//	func ChangeData(arg Data) {
//		arg.value = 999
//		arg.data[100] = 999
//	}
//
// 그래서 포인터 변수를 써줌
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

// 그래서 php에서는 다른 클래스의 값을 변경하려고 할 때 인스턴스를 생성하거나 상속을 통해
// 값을 변경하거나 사용하는데 여긴 클래스도 없고 상속도 없어서 포인터를 사용하나봄
// 근데 인스턴스 생성한다는게 포인터를 이용하겠다는 말이랑 똑같음
// 근데 내장함수로 또 new를 지원하기는 함
func main() {
	// 이런 방법도 있고
	var data Data
	ChangeData(&data)
	// 이런 방식도 있고
	//var p *Data = &Data{}
	// 이런 방식도 있다
	//var p2 = new(Data)

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}
*/

/* 13.문자열
// `<- 이거쓰면 특수문자나 변수같은걸 사용 할 수 없게됨 자바스크립트랑 완전 반대
// 한국말로 백틱,백쿼트라고 하는데 프로그래밍 전반에서는 백틱, 고에서는 백쿼트라고 많이 쓰는듯
//⭐️일본어로도 バッククオート라고 함 오늘 처음알았음
func main() {
	str1 := "Hello\t'World'\n"

	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`
	fmt.Println(str1)
	fmt.Println(str2)
}
// 백쿼트가 효과적일떄는 긴 문자열의 경우 개행이 필요할 때 유용하다
func main() {
	poet1 := "페이커가 쥔 것은 LPL에 대한 다전제 무패 기록이며, LPL이 나아가야 했던 길에서 끝내 얻을 수 없었던 배지입니다. \n우리는 그에게서 영원히 벗어날 수 없습니다.\n그래서 우리는 페이커와 결승전에서 만나기를 갈망했었습니다. \n우리가 마주한 장벽과 정면으로 부딪치기 위해서.\n우리는 이전에 그를 가장 높은 산이며 가장 긴 강으로 불렀지만 산과 바다에도 분명 끝이 있을 것이라 생각했습니다.\n그러나 페이커는 오늘 28세의 나이로 이 경기장에서 완벽한 모습을 보여주었고, \n바다가 끝에 닿으면 하늘이 기슭이 되고, 산의 정상에 오르면 내가 봉우리가 된다.\n는말처럼 그에게는 끝이 없었습니다."

	poet2 := `페이커가 쥔 것은 LPL에 대한 다전제 무패 기록이며, LPL이 나아가야 했던 길에서 끝내 얻을 수 없었던 배지입니다. 우리는 그에게서 영원히 벗어날 수 없습니다.
			그래서 우리는 페이커와 결승전에서 만나기를 갈망했었습니다.
			우리가 마주한 장벽과 정면으로 부딪치기 위해서.
			우리는 이전에 그를 가장 높은 산이며, 가장 긴 강으로 불렀지만, 산과 바다에도 분명 끝이 있을 것이라 생각했습니다.
			“我们曾经称他为最高的山最长的河，以为山海就是尽头.
			그러나 페이커는 오늘 28세의 나이로 이 경기장에서 완벽한 모습을 보여주었고,
			“바다가 끝에 닿으면 하늘이 기슭이 되고, 산의 정상에 오르면 내가 봉우리가 된다.”는 말처럼 그에게는 끝이 없었습니다.`
	fmt.Println(poet1)
	fmt.Println(poet2)
}
// rune 타입으로 한 문자 담기.rune은 int32랑 이름만 다르고 같은 타입임
func main() {
	var char rune = '한'
	fmt.Printf("%T\n", char)
	fmt.Println(char)
	fmt.Printf("%c\n", char)
}
//len()으로 문자열 크기 알아내기(영어 이외에 모두 3바이트씩 먹는다)
func main() {
	str1 := "가나다라마"
	str2 := "abcde"
	str3 := "あいうえお"
	str4 := "アイウエオ"
	str5 := "新勤弟買物"

	fmt.Printf("len(str1) = %d\n", len(str1))
	fmt.Printf("len(str2) = %d\n", len(str2))
	fmt.Printf("len(str3) = %d\n", len(str3))
	fmt.Printf("len(str4) = %d\n", len(str4))
	fmt.Printf("len(str5) = %d\n", len(str5))
}
// []rune 타입 변환으로 글자 수 알아내기
func main() {
	str := "Hello World"

	runes := []rune{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}

	fmt.Println(str)
	fmt.Println(string(runes))

	str2 := "hello 월드"
	runes2 := []rune(str)

	fmt.Println(str2)
	fmt.Println(string(runes2))

	str3 := "hello 월드"
	runes3 := []rune(str3)

	fmt.Printf("len(str) = %d\n", len(str3))
	fmt.Printf("len(runes3) = %d\n", len(runes3))
}
// 문자열 순회(인덱스)
func main() {
	str := "Hello 월드!"

	for i := 0; i < len(str); i++ {
		fmt.Printf(" Type:%T Value:%d StringValue:%c\n", str[i], str[i], str[i])
	}
}
// 문자열 순회(rune)
func main() {
	str := "Hello 월드!"
	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf(" Type:%T Value:%d StringValue:%c\n", arr[i], arr[i], arr[i])
	}
}
// 문자열 순회(range)
func main() {
	str := "Hello 월드!"
	for _, v := range str {
		fmt.Printf(" Type:%T Value:%d StringValue:%c\n", v, v, v)
	}
}
// 문자열 합치기
func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}
// 문자열 비교하기
func main() {
	str1 := "Hello"
	str2 := "Hell"
	str3 := "Hello"

	fmt.Printf("%s == %s : %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s != %s : %v\n", str1, str3, str1 != str3)
}
// 문자열 대소 비교하기
func main() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4)
}
// ⭐️문자열 구조
// (값을 복사하는게 아니라 두 변수가 같은 메모리 주소를 사용하는 방식으로, 정확히는 데이터 포인터와 길이값만 복사해서 변수가 생성되는 것이다.)
// 그래서 아무리 긴 문자열을 복사해서 사용한다 하더라도 메모리 부하에 전혀 문제가 없다
type StringHeader struct {
	Data uintptr
	Len  int
}

func main() {
	str1 := "こんにちは。日本語文字列です。"
	str2 := str1

	// fmt.Printf(str1)
	// fmt.Printf("\n")
	// fmt.Printf(str2)

	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
}
// ⭐️문자열 불변(str, slice가 다른 주소값을 가지고 있으므로 str이 변화하진 않는다)
func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)

	fmt.Printf("str:\t%p\n", unsafe.StringData(str))
	fmt.Printf("slice:\t%p\n", unsafe.SliceData(slice))
}
// ⭐️문자열 합산
// 기존 문자열 메모리 공간을 건드리지 않고
// 새로운 메모리 공간을 만들어서 두 문자열을 합치기 때문에 string합 연산 이후 주소값이 변경됨
// 그래서 문자열 불변 원칙이 준수된다.
func main() {
	var str string = "Hello"
	addr1 := unsafe.StringData(str)
	str += " World"
	addr2 := unsafe.StringData(str)
	str += " Welcome!"
	addr3 := unsafe.StringData(str)
	fmt.Println(str)
	fmt.Printf("addr1:\t%p\n", addr1)
	fmt.Printf("addr2:\t%p\n", addr2)
	fmt.Printf("addr3:\t%p\n", addr3)
}
// ⭐️문자열 합 연산의 메모리 낭비
// 보통의 합연산(합연산을 사용할 때마다 새로운 메모리 공간을할당해서 두 문자열을 더함)
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// 스트링 빌더를 이용
// (strings.Builder는 내부에 슬라이스를 가지고 있기 때문에
// WriteRune메서드를 통해 문자를 더할 때 매번 메모리를 생성하지 않고 기존 메모리 공간에 빈자리가 있으면 그냥 더해버린다)
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
// 불변 원칙을 지키려는 이유->예기지 못한 버그를 방지->합연산이많으면 스트링 빌더를 쓰라는 것이 핵심!

*/

/* 14.패키지
// 패키지 사용하기->
import (
	"fmt"
	"os"
)
// 경로가 있는 패키지 사용하기
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
}
// 겹치는 패키지 문제 별칭으로 풀기
import(
	"text/template"
	htemplate "html/template"
)
//사용하지 않는 패키지 포함하기(_을 이용해서 오류 방지)
import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)
// 패키지 설치->기본패키지는 설치경로에 다있음 외부 패키지의 경우 go.mod에 정의 되어 있음
// public과 private의 구분은 앞이 대문자면 public 소문자면 private
*/

/* 15.숫자 맞추기 게임 만들기
var stdin = bufio.NewReader(os.Stdin)

func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	// 책에는 Seed를 사용하였으나 Seed는 1.20이후에는 권장 되지 않는 방법
	// rand.Seed(time.Now().UnixNano())
	// New와 NewSource를 사용하도록한다

	// --- 책에서의 코드 --- //
	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(100)
	// fmt.Println(n)
	// ----------------- //


	//	왜 이 방식이 더 나은가?
	//	로컬 난수 생성기의 장점:

	//	기본 난수 생성기(rand.Seed)는 전역 상태를 변경하므로 동시성 환경에서 예측 불가능한 동작을 유발할 수 있습니다.
	//	rand.New는 독립된 난수 생성기를 생성하므로 다양한 난수 시퀀스를 병렬적으로 사용할 때 안전합니다.
	//	결정론적 시퀀스 지원:

	//	동일한 source를 재사용하면 항상 동일한 난수 시퀀스를 생성할 수 있습니다.
	//	특정 테스트 환경에서 난수를 재현하기 용이합니다.
	//	Go 1.20 이후의 권장 방식:

	//	rand.Seed는 더 이상 직접 호출할 필요가 없으며, rand.New와 rand.NewSource를 사용하는 것이 더 나은 방법으로 간주됩니다.

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	rn := r.Intn(100)
	cnt := 1

	for {
		fmt.Printf("Enter a number: ")
		n, err := InputValue()
		if err != nil {
			fmt.Println("You must enter a valid number.")
		} else {
			if n > rn {
				fmt.Println("The number you entered is larger.")
			} else if n < rn {
				fmt.Println("The number you entered is smaller.")
			} else {
				fmt.Println("you guessed the number! Attempts made: ", cnt)
				break
			}
			cnt++
		}
	}
}

*/

/* 16.슬라이스
func main() {
	var slice []int

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// 패닉 발생-> 할당되지 않은 메모리 공간에 접근해서 프로그램이 비정상 종료된것
	slice[1] = 10
	fmt.Println(slice)
}

// 초기화 방법
var slice1 = []int{1, 2, 3} // 대괄호 안에 길이를 넣지 않은 것에 주의(php배열 쓰는 방법과 똑같다)
var slice2 = []int{1, 5: 2, 10: 3}

// make를 이용한 초기화
var slice3 = make([]int, 3)

// 접근, 순회는 배열과 똑같다.
// 동적 추가는 append()
func main() {
	var slice = []int{1, 2, 3}
	slice2 := append(slice, 4)
	// 여러 값 추가
	slice3 := append(slice2, 4, 5, 6, 7, 8, 9, 10)

	var slice4 []int
	// 하나씩 추가
	for i := 11; i <= 20; i++ {
		slice4 = append(slice3, i)
	}

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}

// ⭐️⭐️slice동작원리
// 배열과의 동작차이
// 동작차이의 원인
// go언어에서는 모든 값의 대입은 복사로 일어난다.
// 함수에 인수로 전달될 떄나 다른 변수에 대입할 때나 값의 이동은 복사로 일어난다.
// 복사는 타입의 값이 복사됨.
// 포인터는 포인터의 값인 메모리 주소가
// 구조체는 구조체의 모든 필드가
// 배열은 배열의 모든 값이 복사됨
func ChangeArray(array2 [5]int) {
	array2[2] = 200
}

func ChangeSlide(slide2 []int) {
	slide2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slide := []int{1, 2, 3, 4, 5}

	// 값에 의한 복사이므로 원래 array의 값은 그대로 보존(파라미터와 선언한 배열은 완전히 다른 배열)
	ChangeArray(array)

	// 동적배열은 실제배열의 포인터와 실제배열의 길이, 크기를 갖고있는 하나의 구조체.
	// 값에의해 복사되더라도 두 동적배열은 하나의 실제배열의 포인터를 갖고있기때문에
	// 파라미터로 넘겨 실제배열의 주소를 찾아 값을 변경하게 된다
	ChangeSlide(slide)

	fmt.Println("array: ", array) // 1,2,3,4,5
	fmt.Println("slide: ", slide) // 1,2,200,4,5
}
// ⭐️append 사용 시 주의 사항1(slice1, slice2가 하나의 실제 배열을 참조하고 있기 때문에 하나 바꾸면 둘다 바뀐다)
func main() {
	slice1 := make([]int, 3, 5)

	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
// ⭐️append 사용 시 주의 사항2
// append를 통해 현재의 배열보다 긴 요소를 추가 할 경우
// 배열의 길이를 두배로 늘려 새로운 배열을 생성한 뒤
// 새 배열에 값을 복사한다.
// 값에 의한 복사가 일어났으므로
// slice1, slice2는 각자 다른 실제 배열을 참조하고 있다.
func main() {
	slice1 := []int{1, 2, 3}

	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second element")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
}
// 슬라이싱(배열의 일부를 끌어내는 기능 둘다 같은 실제 배열을 참조하고있고 값만 끌어내 쓰는 것)
func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append 500")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))
}
// 유용한 슬라이싱 기능(복제, 추가, 삭제)
// 복제(for)
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	//should use copy(to, from) instead of a loop (S1001)go-staticcheck
	// copy를 쓰는게 더 좋음 이라는 에러문구. 밑에서 설명
	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
// 복제(copy)
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}
// 요소 삭제
func main() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2

	// for문을 이용하는 방법
	// for i := idx + 1; i < len(slice); i++ {
	// 	slice[i-1] = slice[i]
	// }

	// slice = slice[:len(slice)-1]

	// append를 이용하는 방법
	slice = append(slice[:idx], slice[:idx+1]...)

	// ⭐️⭐️⭐️가변인수확장(가변 인수 : 말 그대로 변수의 개수따위가 변할 수 있는 경우에 유용?)
	// * 사용예
	//  슬라이스의 요소들을 개별 인수로 분해해서 함수에 전달, 다른 슬라이스와 병합하는 작업에서 사용
	// 		func 함수이름(param ...T) {
	//			// T는 특정 타입 (int, string 등)
	// 		}

	// * 가변 인수 함수 정의 및 호출
	// 	func sum(nums ...int) int {
	//		total := 0
	//		for _, n := range nums {
	//		total += n
	//		}
	//		return total
	//	}

	//	func main() {
	//		// 개별 인수 전달
	//		fmt.Println(sum(1, 2, 3)) // 출력: 6
	//
	//		// 슬라이스 전달 (가변 인수 확장 사용)
	//		nums := []int{1, 2, 3, 4}
	//		fmt.Println(sum(nums...)) // 출력: 10
	//	}

	// * 슬라이스를 병합할 때 사용
	//	func main() {
	//		slice1 := []int{1, 2, 3}
	//		slice2 := []int{4, 5, 6}
	//
	//		// 슬라이스 병합
	//		merged := append(slice1, slice2...)
	//		fmt.Println(merged) // 출력: [1 2 3 4 5 6]
	//	}

	fmt.Println(slice)
}
// 요소 추가
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	// for문을 이용한 예제
	// slice = append(slice, 0)

	// for i := len(slice) - 2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }

	// slice[idx] = 100

	// append함수로 코드 개선
	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	// 불필요한 메모리 사용이 없도록 코드 개선하기
	// ⭐️인덱스에 :를 붙이는 의미 -> 그 인덱스부터 마지막까지의 인덱스를 의미함
	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[idx] = 100

	fmt.Println(slice)
}
// 슬라이스 정렬
// int
func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}
// 구조체 슬라이스 정렬
type Student struct {
	Name string
	Age  int
}

// []Student의 별칭 타입 Students. 학생이 여러명이므로 구조체들의 집합을 배열로서 관리하려고 하는듯
type Students []Student

// 이런식으로 함수를 정의해서 정렬하는거 같은데 자세한건 메서드와 인터페이스 파트에서 다룸
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"Hwarang", 31}, {"YamaBack", 52}, {"Ryu", 42},
		{"Ken", 38}, {"Sonhana", 18}}

	sort.Sort(Students(s))
	fmt.Println(s)
}
*/

/* 17.메서드
// 메서드 선언(함수와 메서드의 차이는 눈에 숙지 할 것 참조)
type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드 선언방식 (a *account)를 리시버라고 부른다. 어느 구조체에 속해있는지를 나타내줌
// *account는 리시버 타입
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{100}
	withdrawFunc(a, 30)
	a.withdrawMethod(30)
	fmt.Printf("%d \n", a.balance)
}
// 별칭 리시버 타입
type myInt int

func (a myInt) add(b int) int {
	return int(a) + b
}

func main() {
	var a myInt = 10
	fmt.Println(a.add(30))
	var b int = 20
	fmt.Println(myInt(b).add(50))
}
// ⭐️메서드는 왜 필요한가?
// 좋은 프로그래밍이라면 결합도를 낮추고 응집도를 높여야한다.
// 결합도란? 結合度 (けつごうど, Ketsugōdo) 객체간의 의존관계를 나타낸다.
// 응집도란? 凝集度 (ぎょうしゅうど, Gyōshūdo) 모듈내 요소들의 상호 관련성을 나타낸다
// 산탄총 수술문제? 작은 변화에도 산탄총을 맞은 듯 많은 코드 영역을 수정하는 경우
// 메서드는 기능과 데이터를 묶어주는 역할. 즉 응집도를 높이는 역할을 한다.
// OOP : 객체 지향 프로그래밍(이제는 외울떄가 됐다)

// ⭐️⭐️⭐️포인트 메서드와 값타입 메서드
// 리시버를 값 타입과 포인터로 정의 할 수 있다
type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인트 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA *account = &account{100, "Joe", "Park"}

	// 참조하고 있는 실제 구조체의 필드값을 변경하였으므로 밸런스는 70이 됨
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	// 값에 의한 복사가 일어났으므로 실제 구조체의 필드값은 변경이 일어나지 않음 밸런스값은 그대로 70
	// a2의 인스턴스 값은 복사가 일어나 -20이 되었음
	// withdrawValue에서 오류가 나는이유는 값에 의한 복사가 일어나서 기능을 수행했지만
	// 어떠한 리턴도 없어 의미가 없다는 처리라는 것을 나타내고 있다.
	// 챗지피티의 설명 -> Go 코드에서 구조체 필드에 값을 할당했지만, 그 할당이 아무런 효과가 없다는 경고입니다.
	// 이는 구조체의 값을 복사한 상태에서 작업을 수행했거나, 필드 값을 수정하는 방식이 잘못된 경우 발생합니다.
	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	// 그래서 값을 변경하려면 값을 변경한뒤 다시 구조체로 반환을해줘야함
	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	// ⭐️요는 포인터 메서드는 메서드 내부에서 리시버의 값을 변경시킬수 있고 값 타입 메서드는 내부에서 변경안됨
	// 그러므로 포인터 메서드는 인스턴스 중심. 값 타입 메서드는 값 중심이 된다.
}
*/

/* 18.인터페이스

// 선언
// 메서드에는 반드시 메서드명이 필요
// 매개변수 반환이 다르더라도 이름이 같은 메서드는 있을수 없음
// 인터페이스에서는 메서드 구현을 하지않음
// 일반적인 인터페이스의 개념과 같다

	type Stringer interface {
		String() string
	}

	type Student struct {
		Name string
		Age  int
	}

	func (s Student) String() string {
		return fmt.Sprintf("Hello, I'm %s, and I'm %d years old.", s.Name, s.Age)
	}

	func main() {
		student := Student{"Cento", 12}
		var stringer Stringer

		stringer = student

		fmt.Printf("%s\n", stringer.String())
	}

// 인터페이스를 사용하는 이유
// 인터페이스를 이용하면 구체화된 객체가 아닌 인터페이스만으로 메서드를 호출 할 수 있기떄문
import (

	fedex "golang/GoPractice/interfacePractice/fedex"
	post "golang/GoPractice/interfacePractice/post"

)

// 이렇게 메소드를 여러개 나눠서 해야하지만 인터페이스를 사용하면 공통으로 관리가 가능
// func SendBookFedex(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

// func SendBookPost(name string, sender *post.PostSender) {
// 	sender.Send(name)
// }

	type Sender interface {
		Send(parcel string)
	}

// SendBook은 인수가 어떤 인수가 오던지 상관이 없다.
// 어떤 타입이던지 Send 인덱스만 제공하면 어떤 타입이라도 받아들이는 것이 가능

	func SendBook(name string, sender Sender) {
		sender.Send(name)
	}

	func main() {
		// senderFedex := &fedex.FedexSender{}
		// SendBookFedex("Book1", senderFedex)
		// SendBookFedex("book2", senderFedex)

		// senderPost := &post.PostSender{}
		// SendBookPost("book3", senderPost)
		// SendBookPost("book3", senderPost)

		senderFedex := &fedex.FedexSender{}
		SendBook("Book1", senderFedex)
		SendBook("book2", senderFedex)

		senderPost := &post.PostSender{}
		SendBook("book3", senderPost)
		SendBook("book3", senderPost)
	}

// ⭐️추상화 계층, 인터페이스는 추상화 계층이다.
// 추상화(抽象化 ちゅうしょうか、Abstraction) : 내부 동작을 감춰서 서비스를 제공하는 쪽과 사용하는 쪽 모두에게 자유를 주는 방식
// 디커플링 : 추상화 계층을 이용해 의존 관계를 끊는 것

// ⭐️덕 타이핑 : 타입 선언 시 인터페이스 구현 여부를 명시적으로 나타낼 필요 없이 인터페이스에 정의한 메서드 포함여부만으로 결정하는 방식

// 포함된 인터페이스 : 구조체에서 다른 구조체를 포함된 필드로 가질 수 있듯이 인터페이스도 다른 인터페이스를 포함 할 수 있다.

	type Reader interface {
		Read() (n int, err error)
		Close() error
	}

	type Writer interface {
		Write() (n int, err error)
		Close() error
	}

	type ReadWriter interface {
		Reader
		Writer
	}

// 빈 인터페이스 interface{}를 인수로 받기
// interface{} : 메서드를 가지고 있지 않은 빈 인터페이스

	func PrintVal(v interface{}) {
		switch t := v.(type) {
		case int:
			fmt.Printf("v is int %d\n", int(t))
		case float64:
			fmt.Printf("v is float64 %f\n", float64(t))
		case string:
			fmt.Printf("v is string %s\n", string(t))
		default:
			fmt.Printf("Not supported type: %T:%v\n", t, t)
		}
	}

	type Student struct {
		Age int
	}

	func main() {
		PrintVal(10)
		PrintVal(3.14)
		PrintVal("Hello")
		PrintVal(Student{15})
	}
// 인터페이스 기본 값 nil: 인터페이스 기본값은 유효하지 않은 메모리 주소를 나타내는 nil
type Attacker interface {
	Attack()
}

func main() {
	var att Attacker //기본값은nil
	att.Attack() // att가 nil이기떄문에 런타임 에러 발생
}
// ⭐️인터페이스 변환하기(이해 잘 안됨)
// 구체화된 다른 타입으로 타입 변환
type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15}
	PrintAge(s)
}
// 타입 변환시 런타임 에러가 나는 경우
type Stringer interface {
	String() string
}

type Student struct {
}

func (s *Student) String() string {
	return "Student"
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

func ConvertType(stringer Stringer) {
	// run time error *Student타입운 Strubger 인터페이스로 쓰일 수 있지만
	// stringer 값이 *Student타입이 아니기때문에 에러가 발생
	student := stringer.(*Student)
	fmt.Println(student)
}

func main() {
	actor := &Actor{}
	ConvertType(actor)
}
// 다른 인터페이스로 타입 변환하기
type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	// Reader interface변수를 Closer 인터페이스로 타입 변환
	// 런타임 에러 발생
	c := reader.(Closer)
	c.Close()
}

func main() {
	// file 포인터 인스턴스를 ReadFlie() 함수의 인수로 사용합니ㅏ다.
	file := &File{}
	ReadFile(file)
}
// 타입 변환 성공 여부 반환
var a interface
t, ok := a.(ConcreateType)
*/

/* ⭐️19.함수 고급편(이쪽 부분이 잘 이해가 되지 않는다)
// 가변 인수 함수
// 키워드 사용
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums type: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 빈 인터페이스 사용
func Print(args ...interface{}) string {
	for _, arg := range args {
		switch f := arg.(type) {
		case bool:
			val := arg.(bool)
		case float64:
			val := arg.(float64)
		case int:
			val := arg.(int)
		}
	}
	return val
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
// ⭐️⭐️defer 지연 실행 : 함수가 종료되기 직전에 호출되며 호출 순서는 역순으로 호출이 된다.
func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create file")
		return
	}

	defer fmt.Println("It will definitely be called.")
	defer f.Close()
	defer fmt.Println("File closed.")

	fmt.Println("Write hello world in the file.")
	fmt.Fprintln(f, "Hello Worlds")
}
// 함수 타입 변수
func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {

	// var operator func(int, int) int
	// operator = getOperator("*")

	operator := getOperator("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
// 함수 리터럴(익명 함수, 람다)
type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	fn := getOperator("*")

	result := fn(3, 4)
	fmt.Println(result)
}
// 함수 리터럴 내부 상태
func main() {
	i := 0
	f := func() {
		i += 10
	}

	i++

	f()

	fmt.Println(i)	// 값이 11인 이유는 람다는 선언한 시점이 아닌 호출되는 시점에 움직이므로 i++이 한번 움직여서 그렇다
}
// 함수 리터럴 내부 상태 주의점(여기 뭔가 이상)
// 캡쳐 : 함수 리터럴 외부 변수를 내부 상태로 가져오는 것(참조로 가져오므로 조심해야한다.)
func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
// 파일 핸들을 내부 상태로 사용하는 예
type Writer func(string)

func writeHello(writer Writer) {
	writer("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
*/

/* ⭐️20. 자료구조
// 리스트
// 배열과의 차이점은 배열은 연속된 메모리에 데이터를 저장하지만 리스트는 불연속된 메모리에 데이터를 저장함
// 리스트는 각 데이터를 담고 있는 요소들을 포인터로 연결한 자료구조다. 링크드 리스트라고 부르기도함
// 기본 사용법
func main() {
	v := list.New()
	// 리스트 뒤에 요소 추가
	e4 := v.PushBack(4)

	// 리스트 앞에 요소 추가
	e1 := v.PushFront(1)

	// e4 요소 앞에 추가
	v.InsertBefore(3, e4)

	// e1 요소 뒤에 추가
	v.InsertAfter(2, e1)

	// 리스트 순회
	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}

	fmt.Println()

	// 리스트 역순회
	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
}
// 배열 vs 리스트
// ⭐️데이터 추가에 대한 차이는 페이지 418을 참고
// 데이터 지역성 : 데이터가 밀집한 정도를 말함.	데이터 지역성을 참고하여 배열을 사용할 지 리스트를 사용할 지 선택해야한다

// 큐(Queue) : 먼저입력된 값이 먼저 출력(FIFO、first in first out)
// ex) 대기열
// 대기열 작업, 명령큐와 같은 순서가 유지되어야하는 경우에 자주 사용
// 배열, 리스트 둘다 만들 수 있으나 리스트가 효율이 훨씬 좋음
type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = queue.Pop()
	}
}
// 스택 구현하기
// FILO First in last out
// 스택은 요소의 추가와 삭제가 항상 맨 뒤에서 발생하기 때문에 배열로 만들어도 성능에 손해가 없다
// 그래서 큐는 리스트, 스택은 배열로 만드는 것이 일반적이다.
type Stack struct {
	v *list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func main() {
	stack := NewStack()
	for i := 1; i < 5; i++ {
		stack.Push(i)
	}
	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v -> ", val)
		val = stack.Pop()
	}
}
// 링
// 맨뒤의 요소와 맨 앞의 요소가 연결되어있는 자료구조
// 리스트 기반
// 원형으로 연결되어 있어 환형 리스트라고 부르기도 함
// 시작과 끝이 없으며 현재위치만 존재
// 저장할 개수가 고정되고 오래된 요수는 지워도 되는 경우에 적합
// 실행 취소, 고정 크기 버퍼 기능, 리플레이 기능
func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
// 맵
// 키와 값으로 데이터를 저장하는 자료구조
// 딕셔너리, 해시테이블, 해시맵
func main() {
	m := make(map[string]string)
	m["hwalarg"] = "Seoul"
	m["hana"] = "inchoen"
	m["BackDosan"] = "Busan"
	m["Thuonder"] = "Test"

	m["Thuonder"] = "Test2"

	fmt.Printf("hana address is %s.\n", m["hana"])
	fmt.Printf("BackDosan address is %s.\n", m["BackDosan"])
}
// 순회
type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product, 0)
	m[16] = Product{"Ballpen", 500}
	m[46] = Product{"Eraser", 200}
	m[78] = Product{"Ruler", 1000}
	m[345] = Product{"Sharp", 3000}
	m[897] = Product{"Sharp lead", 500}

	for k, v := range m {
		fmt.Println(k, v)	// 정렬되지 않은 값을 출력함. 내부에서 요소를 보관할때 입력한 순서와도 키 값과도 상관없이 데이터를 보관
	}
}
// 맵 요소 삭제와 없는 요소 확인
// delete()함수로 요소를 삭제
// 삭제시 키에 맞는 요소가 없다면 값 타입의 기본값을 반환
func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	fmt.Println(m[3])
	fmt.Println(m[1])
}

// 맵, 배열, 리스트 속도 비교
// 맵 : 속도가 매우 빠르고 삭제, 추가, 읽기에서 요소 개수와 상관없이 속도가 일정, 인덱스를 사용해 접근이 안되고 순서보장이 안됨. 메모리또한 많이 사용
// 배열 : 추가, 삭제에서 요소 개수가 많아질수록 오래 걸림
// 리스트 : 읽기에서 요소 개수가 많아질수록 오래 걸림

// ⭐️맵의 원리(페이지 433참고)
*/

/* 21. 에러 핸들링
// 에러 반환
func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("Failed to create file.", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("Failed to read file.", err)
			return
		}
	}
	fmt.Println("File Content : ", line)
}
// 사용자 에러 반환
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

func main() {
	sqrt, err := Sqrt(-2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Sqrt(-2) = %v\n", sqrt)
}
// 에러 타입
type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "Password length is short."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("Registration successful.")
	}
}
// 에러 랩핑 : 에러를 감싸서 새로운 에러를 만듬
func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
// 패닉 : 프로그램을 정상 진행시키기 어려운 상황을 만났을 때 프로그램 흐름을 중지시키는 기능. panic()
func divide(a, b int) {
	if b == 0 {
		panic("b must not be 0")
	}
	fmt.Printf("%d /%d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}

// 패닉 생성(이하의 모든 패턴이 가능)
panic(42)
panic("unreachable")
panic(fmt.Errorf("This is error num:%d", num))
panic(SomeType(someData))
// 패닉전파와 복구
func f() {
	fmt.Println("Function f() start.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover panic  -", r)
		}
	}()

	g()
	fmt.Println("Function f() is terminated.")
}

func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("The divisor cannot be zero.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("Program continues to run.")
}
*/

/* 22. 고루틴과 동시성 프로그래밍
// ⭐️⭐️⭐️스레드란? 프로세스 안의 세부 작업 단위
// 프로세스? 메모리 공간에 로딩되어 동작하는 프로그램. 한개 이상의 스레드를 가지고 있다.
// 스레드가 하나면 싱글 스레드 프로세스, 여럿이면 멀티 스레드 프로세스라고 한다.
// CPU 코어는 기본적으로 한 번에 한 명령밖에 수행 할 수가 없다.
// CPU 코어가 여러개이면 동시에 명령을 수행 할 수 있고
// 싱글 코어 CPU라고 해도 스레드를 빠르게 전환해 가면서 스레드를 수행하면 사용자 입장에서는
// 동시에 수행하는 것 처럼 보인다.
// 고루틴은 경량 스레드로 함수나 명령을 동시에 실행할 때 사용함.

// 컨텍스트 스위칭 비용
// CPU 코어가 여러 스레드를 전환하면서 수행하면 많은 비용이 발생하고 이것은 컨텍스트 스위칭 비용이라고 한다.
// 스레드를 전환하려면 현재 상태(스레드 컨텍스트)를 보관해야하고
// 다시 현재의 스레드를 진행할떄 스레드 컨텍스트를 복원해야한다.
// 이때 전환 비용이 발생
// 스레드 컨텍스트 : 스레드의 명령 포인터, 스택 메모리등의 정보
// golang에서는 코어마다 OS스레드를 하나만 할당해서 사용하기 때문에
// 이런 비용이 발생하지 않음

// ⭐️⭐️고루틴 사용
func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}

// ⭐️⭐️서브 고루틴이 종료될 때까지 기다리기
// sync패키지의 WaitGroup객체를 사용

var wg sync.WaitGroup

//wg.Add(3)	// 작업 개수 설정
//wg.Done()	// 작업이 완료 될 때마다 호출
//wg.Wait()	// 모든 작업이 완료될 때까지 대기

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("The sum from %d to %d is %d. \n", a, b, sum)
	wg.Done()
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	wg.Wait()

	fmt.Println("All calculations are over")
}
// 고루틴의 동작방법(페이지 468 참조)

// 시스템 콜 호출 시(페이지 469 참조)
// 시스템 콜 : 운영차제가 지원하는 서비스를 호출 할 때(네트워크 기능 등)
// 호출 시 운영체제에서 해당 서비스가 완료될 때까지 대기 해야한다.
// ex) 네트워크로 데이터를 읽을 때 데이터가 들어 올 떄까지 대기상태가 됨

// 동시성 프로그래밍 주의점
// 동일한 메모리 자원에 여러 고루틴이 접근할 때 발생
// ex)화가가 화분을 놓고 그림그리는데 화분 주인이 화분을 치워버리면 화가는 화분을 제대로 못그림
// 두 사람이 동일 자원에 접근했기때문에 발생하는 문제
// 고루틴은 각 코어에서 별도로 동작하지만
// 같은 메모리 공간에 동시에 접근해서 값을 변경시킬 수 있음
type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
// 뮤텍스를 이용한 동시성 문제 해결
// 고루틴에서 값을 변경 할 때 다른 고루틴이 건들지 못하게하면된다.
// 뮤텍스 mutual exclusion의 약자(상호 배제, 자원 접근 권한)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
// 뮤텍스와 데드락
// 뮤텍스의 문제점
// 1. 오직 하나의 고루틴만 제한하기때문에 동시성 프로그래밍으로 얻는 성능 향상을 얻을수 없다.
// 2. 데드락이 발생할수있음(데드락 : 프로그램을 완전히 멈추게 함)-> 어떤 고루틴도 원하는 만큼 뮤텍스를 확보하지 못해서 무한히 대기하게 되는 현상
// 데드락 구현
var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "fork", "spoon")
	go diningProblem("B", spoon, fork, "spoon", "fork")
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("I'm about to eat.\n", name)
		first.Lock()
		fmt.Printf("%s %s acquire.", name, firstName)
		second.Lock()
		fmt.Printf("%s %s acquire.", name, secondName)

		fmt.Printf("I am eating.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
// 또 다른 자원 관리 기법
// 같은 자원에 접근하지 않도록 영역과 역할을 나눠서 사용
// 1. 영역을 나누어 관리(역할은 다음장에서 설명)
type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d job start\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d job done - result: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
*/

/* 23. 채널과 컨텍스트
// 채널이란? : 고루틴끼리 메시지를 전달 할 수 있는 메시지 큐. 큐이므로 선입선출이다(FIFO)
// 사용하기 위해선 먼저 인스턴스를 생성해야한다.(make함수로 생성)
var messages chan string = make(chan string)

// 채널에 데이터 넣기(연산자 <-)
messages <- "This is message"

// 채널에서 데이터 뺴기(채널 인스턴스에 데이터가 없으면 데이터가 들어 올 떄까지 대기함)
var msg string = <- messages
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", n*n)
	wg.Done()
}
// 채널 크기
// 채널을 생성하면 크기가 0인 채널이 생성됨. 채널에서 데이터를 가져가지 않아 데드락이 걸리는 예시
func main() {
	ch := make(chan int)
	ch <- 9
	fmt.Println("Never print")
}
// 버퍼를 가진 채널
// 버퍼? : 내부에 데이터를 보관할 수 있는 메모리 영역(make함수에 크기를 적어주면된다.)
var chan string messages = make(chan string, 2)
// 채널에서 데이터 대기
func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // for문이 끝날때 채널을 닫아준다
	wg.Wait()
}
// select문(채널에 데이터가 오는것을 마냥 기다리지않고 대기하는 동안 다른 작업을 수행할 수 있게해줌 또는 여러 채널을 동시에 대기하고 싶을떄 사용)
func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
// 일정 간격으로 실행
func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
// 채널로 생산자 소비자 패턴 구현하기
// 22장의 역할을 나누는방법이 여기서 이어짐(고루틴의 뮤텍스를 사용하지 않는 방법)
type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup

var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			// Make a body
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		// Make a body
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}

	wg.Done()
}
// 컨텍스트 사용하기(작업 가능 시간, 작업 취소등의 조건을 지시할 수 있는 역할, 새로운 고루틴으로 작업을 시작 할 때 일정시간동안만 작업을 지시하거나 외부에서 작업을 취소할 떄 사용)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
// 작업 시간을 설정한 컨텍스트
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square:%d", n*n)
	}
	wg.Done()
}

*/

/* 24. 제네릭 프로그래밍

// golang은 강타입 언어이기때문에 변수타입에 의해 동작하지 않는 함수가 많다. 이걸 제네릭으로 대응
// 모든 숫자타입을 사용할 수 있는 함수
// 기존 코드는 이제 constraints를 이용해서 사용할 필요가 없다.
// func add[T constraints.Integer | constraints.Float](a, b T) T {

	func add[T int | float64](a, b T) T {
		return a + b
	}

	func main() {
		var a int = 1
		var b int = 2
		fmt.Println(add(a, b))

		var f1 float64 = 3.14
		var f2 float64 = 1.43
		fmt.Println(f1, f2)
	}

// 제네릭함수 : 타입 파라미터를 통해 여러 타입에 대해서 동작하는 함수
// 모든 타입값을 출력하는 제네릭 함수와 인터페이스를 이용한 함수
// 유사하나 면이 많으나 다른 개념이고 다른 효과를 가지고 있기때문에 잘 구분해야함

	func Print[T1 any, T2 any](a T1, b T2) {
		fmt.Println(a, b)
	}

	func Print2(a, b interface{}) {
		fmt.Println(a, b)
	}

	func main() {
		Print(1, 2)
		Print(3.14, 1.43)
		Print("Hello", "World")
		Print(1, "Hello")

		Print2(1, 2)
		Print2(3.14, 1.43)
		Print2("Hello", "World")
		Print2(1, "Hello")

}
// 타입제한
// 타입 제한을 사용할떄 인터페이스를 사용해서 정의하는 것이 일반적이다.

	type Integer interface {
		int8 | int16 | int32 | int64 | int
	}

	func add[T Integer](a, b T) T {
		return a + b
	}

	func main() {
		var a int = 1
		var b int = 2
		fmt.Println(add(a, b))
	}

// 타입 제한 더 알아보기

	type Integer interface {
		// int8 | int16 | int32 | int64 | int
		~int8 | ~int16 | ~int32 | ~int64 | ~int // 별칭타입을 사용하려면 앞에 ~를 붙여준다
	}

	func add[T Integer](a, b T) T {
		return a + b
	}

type MyInt int

	func main() {
		add(1, 2)
		var a MyInt = 3
		var b MyInt = 5
		add(a, b) // MyInt타입이 타입제한에 포함되지 않아 에러가 발생(interface에 ~vlfdy)
	}

// 타입 제한에 메서드 조건 더하기

	type ComparableHasher interface {
		comparable
		Hash() uint32
	}

type MyString string

	func (s MyString) Hash() uint32 {
		h := fnv.New32a()
		h.Write([]byte(s))
		return h.Sum32()
	}

	func Equal[T ComparableHasher](a, b T) bool {
		if a == b {
			return true
		}
		return a.Hash() == b.Hash()
	}

	func main() {
		var str1 MyString = "Hello"
		var str2 MyString = "World"
		fmt.Println(str1, str2)
	}

// 제네릭 함수 예

	func Map[F, T any](s []F, f func(F) T) []T {
		rst := make([]T, len(s))
		for i, v := range s {
			rst[i] = f(v)
		}
		return rst
	}

	func main() {
		doubled := Map([]int{1, 2, 3}, func(v int) int {
			return v * 2
		})

		uppered := Map([]string{"Hello", "world", "abc"}, func(v string) string {
			return strings.ToUpper(v) // 대문자변경
		})

		tostring := Map([]int{1, 2, 3}, func(v int) string {
			return "str" + strconv.Itoa(v) // 문자열로 변경하는 슬라이스
		})

		fmt.Println(doubled)
		fmt.Println(uppered)
		fmt.Println(tostring)
	}

// 제네릭 타입

	type Node[T any] struct {
		val  T
		next *Node[T]
	}

	func NewNode[T any](v T) *Node[T] {
		return &Node[T]{val: v}
	}

	func (n *Node[T]) Push(v T) *Node[T] {
		node := NewNode(v)
		n.next = node
		return node
	}

	func main() {
		node1 := NewNode(1)
		node1.Push(2).Push(3).Push(4)

		for node1 != nil {
			fmt.Print(node1.val, " - ")
			node1 = node1.next
		}
		fmt.Println()

		node2 := NewNode("Hi")
		node2.Push("Hello").Push("How are you")

		for node2 != nil {
			fmt.Print(node2.val, " - ")
			node2 = node2.next
		}

		fmt.Println()
	}

// 인터페이스와 제네릭은 무엇이 다른가?
// 빈 인터페이스를 이용하면 모든 타입값을 가질 수 있으나 그 값을 사용할때 실제 타입값으로 타입 변환을 해야하고
// 넣을 떄 값의 타입과 뺄 때 값의 타입을 정확히 알고 있어야 한다는 문제가 있음
// 제네릭 타입을 사용하는 경우 타입 파라미터에 의해서 필드 타입이 결정되므로 값을 사용할 때 타입 변환이 필요없다.

	type NodeType1 struct {
		val  interface{}
		next *NodeType1
	}

	type NodeType2[T any] struct {
		val  T
		next *NodeType2[T]
	}

	func main() {
		node1 := &NodeType1{val: 1}
		node2 := &NodeType2[int]{val: 2}

		// 밑의 코드는 에러 발생.
		// 첫번째 필드의 타입이 인터페이스이기 때문에 바로 변수값으로 대입 할 수가 없음
		// var v1 int = node1.val
		// 에러를 제거하기위해선 인터페이스의 타입값을 맞춰줄 필요가 있다
		var v1 int = node1.val.(int)

		fmt.Println(v1)
		var v2 int = node2.val
		fmt.Println(v2)
	}

// 성능의 차이

	func main() {
		var v1 int = 3
		var v2 interface{} = &v1
		var v3 int = *(v2.(*int))

		fmt.Printf("v1: %x %T\n", &v1, &v1)
		fmt.Printf("v2: %x %T\n", &v2, &v2)
		fmt.Printf("v3: %x %T\n", &v3, &v3)
	}

// 언제 제네릭 프로그래밍을 사용해야 하는가
// 과사용시 코드 가독성이 떨어짐
// ⭐️성급한 최적화가 프로그래밍에서 모든 죄악의 뿌리이다.
// 동작하는 코드를 먼저 작성하고 그다음에 최적화나 개선을 해도 늦지 않는다는 뜻
// 자료구조가 사용하기 좋은곳이다.
// 객체의 타입이 아닌 객체의 기능이 강조되는 곳은 인터페이스가 더 좋다.

// 제네릭을 사용해 만든 유용한 기본 패키지
// slices

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 55},
		{"Bob", 24},
		{"Gopher", 13},
	}
	n, found := slices.BinarySearchFunc(people, "Bob", func(a Person, b string) int {
		return cmp.Compare(a.Name, b)
	})

	fmt.Println("Bob:", n, found)
}
func main() {
	names := []string{"Alice", "Bob", "Vera"}
	fmt.Println("Equal:", slices.Compare(names, []string{"Alice", "Bob", "Vera"}))
	fmt.Println("V < X:", slices.Compare(names, []string{"Alice", "Bob", "Xera"}))
	fmt.Println("V > C:", slices.Compare(names, []string{"Alice", "Bob", "Cat"}))
	fmt.Println("3 > 2:", slices.Compare(names, []string{"Alice", "Bob"}))
}
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6, 7}
	s3 := []int{8, 9, 10}

	s := slices.Concat(s1, s2, s3)
	fmt.Println(s)
}
// maps
func main() {
	m1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	m2 := map[string]int{
		"one": 10,
	}

	maps.Copy(m2, m1)
	fmt.Println("m2 is:", m2)

	m2["one"] = 100
	fmt.Println("m1 is:", m1)
	fmt.Println("m2 is:", m2)

	m3 := map[string][]int{
		"one": {1, 2, 3},
		"two": {4, 5, 6},
	}

	m4 := map[string][]int{
		"one": {7, 8, 9},
	}

	maps.Copy(m4, m3)
	fmt.Println("m4 is:", m4)

	m4["one"][0] = 100
	fmt.Println("m3 is:", m3)
	fmt.Println("m4 is:", m4)
}
*/

/* 25. 단어 검색 프로그램 만들기
goprojects폴더 참조
*/

/* 26. 테스트와 벤치마크하기(main_test.go참조)

// 테스트

	func square(x int) int {
		//return 81
		return x * x
	}

	func main() {
		fmt.Printf("9 * 9 = %d\n", square(9))
	}

// 벤치마크
// 피보나치 수열을 만드는 예제
func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}

	return rst
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
*/

/* 27. 프로파일링으로 성능 개선하기
// 프로파일링이란? : 프로그램의 성능 지표를 프로그램이 실행 중에 실시간으로 측정 기록하는 것
// 측정 성능지표는 프로그램 실행 시간, 메모리 사용량, 함수 호출 히간과 빈도, 메모리가 생성되는 시점과 빈도
// 이러한 정보를 수집해 성능을 최적화 하는데 사용한다.

// 특정 구간 프로파일링
// go tool pprof cpu.prof 로 pprof툴을 시작
// 시작하면 pprof콘솔 창이 표시 되고 여기에서 명령어를 입력하여 프로파일링을 할 수가 있다.
// top -> 점유율을 높은 순서대로 보여줌
// web -> 그래프를 그릴 수 있는 오픈 소스 프로그램인 Graphviz가 필요, 실행하면 자동으로 인터넷 브라우저가 실행되어 분석 결과가 그래프 형태로 출력(svg)

// 이 함수대로라면 많은 시간이 걸리기 때문에 함수를 개선 해야 한다.
func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
// 개선한 함수
var fibMap [65535]int

func Fib(n int) int {
	f := fibMap[n]
	if f > 0 {
		return f
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		f = 1
	} else {
		f = Fib(n-1) + Fib(n-2)
	}

	fibMap[n] = f
	return f
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	pprof.StartCPUProfile(f)

	defer pprof.StopCPUProfile()
	fmt.Println(Fib(50))

	time.Sleep(10 * time.Second)
}
// 서버에서 프로파일링
// 웹 서버와 같이 계속 실행되는 서버 프로그램에서는 특정 구간만 프로파일링 하기가 어려움
// 웹서버에서 프로파일링 하는 법
// 로컬호스트 8080에 접속하여 http://localhost:8080/debug/pprof/을 입력하면 프로파일링 화면을 볼 수 있음
// 프로 파일링 페이지의 메뉴
// allocs : 프로그램에서 할당된 모든 메모리와 어디에서 할당 되었는지 소스 위치를 통해 알수 있게 해줌
// block : Mutex나 WaitGroup, 채널 등과 같은 멀티 스레드 환경에서 현재 대기하는 객체들을 보여줌
// cmdline : 이 프로그램이 실행 될 떄 어떤 실행 인수로 실행되었는지 보여줌
// goroutine : 현재 실행되고 있는 모든 고루틴을 보여줌
// heap : 현재 메모리가 할당되어 사용중인 객체들을 보여줌
// mutex : 현재 대기중인 mutex의 콜스택을 보여줌
// profile : CPU 프로파일링을 시작하고 그 결과를 파일로 다운로드 받을 수 있습니다.
// 프로파일링하는 시간은 seconds 파라미터로 조정하거나 파라미터가 없으면 30초동안 동작함.
// 프로파일링 결과는 Go 프로파일링 툴을 통해서 분석 할 수 있음
// ex) http://localhost:8080/debug/pprof/profile?seconds=5 5초동안만 동작함
// threadcreate : 새로운 스레드를 생성한 콜스택을 보여줌
// trace : 트레이스 정보 수집을 시작한 뒤 그 결과를 파일로 다운로드 받을 수 있음. 트레이스 정보는 Go프로파일링 툴을 통해서 분석 할 수 있음
import (
	"math/rand"
	"net/http"
	_ "net/http/pprof" // 웹 프로파일링을 실행
	"time"
)

func main() {
	http.HandleFunc("/log", logHandler)
	http.ListenAndServe(":8080", nil)
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
	}()

	select {
	case status := <-ch:
		w.WriteHeader(status)
	case <-time.After(200 * time.Microsecond):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
// Hey로 부하 테스트 하기
// Hey 프로그램을 사용해 웹서버에 부하를 줘서 테스트
// 요청을 반복 전송해 웹서버 성능을 테스트 할 때 사용
// https://github.com/rakyll/hey에서 다운받아
// .\hey.exe http://localhost:8080/log로 실행
// 결과를 통해 웹 서버 요청에 대한 응답 속도를 알 수 있음

// 프로파일링 결과 분석
// CPU성능이 어떤 곳에서 주로 사용되는지 알 수 있고
// 병목점을 발견, 개선해 전체 성능을 향상시킬수 있음
// 8:2법칙 전체 코드 20프로가 전체 성능중 80퍼센트를 차지한다.
// 이 20프로의 코드의 개선이 매우 중요하다.
*/

/* 28. HTTP 웹 서버 만들기
1.기본 net/http 패키지를 사용
2. 핸들러 등록과 웹서버 시작이라는 두단계를 거침

// 핸들러 등록
각 HTTP요청 URL 경로에 대응 할 수 있는 핸들러를 등록
핸들러란? 각 HTTP 요청 URL이 수신됐을 때 그것을 처리하는 함수 또는 객체

// 웹서버 시작
ListenAndServe() 함수를 호출해 웹서버를 시작

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello World") // 웹 핸들러 등록
	})

	http.ListenAndServe(":3003", nil) // 웹 서버 시작
}

// ⭐️⭐️⭐️⭐️⭐️ HTTP 동작 원리
웹 브라우저에 https://goldnabbit.co.kr:3000을 입력 한 뒤 엔터키를 눌렀을 때 발생하는 일
도메인은 https://goldnabbit.co.kr 이 되고 포트 번호는 3000이다.
웹 브라우저는 먼저 도메인 네임 시스템에 도메인에 해당하는 IP주소를 요청함
		웹 브라우저 -> goldnabbit.co.kr -> DNS
		웹 브라우저 <- IP 주소 <- DNS
IP 주소는 목적지(컴퓨터)를 나타낸다면 포트 번호는 수신한 데이터를 놓을 (컴퓨터 내) 창구와 같다.
IP 주소는 컴퓨터 자체
포트 번호는 해당 컴퓨터 내 데이터를 수신할 수 있는 창구
컴퓨터는 0번부터 65535번 포트를 가지고 있다.
포트 번호 없이 https://goldnabbit.co.kr을 입력하면 기본 포트 번호로 요청을 전송함.
HTTP 기본 포트 번호는 80번
HTTPS 기본 포트 번호는 443번
HTTPS:// 는 데이터를 보내는 통신 규약
HTTP 통신 규약에 보인 기능을 추가한 통신규약
HTTP는 하이퍼 텍스트 전송 규약 (HypwrText Transfer Protocol)의 약자로 말 그대로 하이퍼 텍스트를 전송하는 통신 규약(프로토콜)
하이퍼 텍스트란 하이퍼 링크를 포함한 멀티 미디어 텍스트로 문자 그림 이미지등의 멀티미어를 포함하고 다른 문서로 연결되는 링크를 제공하눈 문서 포맷
웹에서 하이퍼 텍스트 문서를 사용하기 떄문에 문자, 이미지, 음악, 동영상등을 볼 수 있고 링크를 클릭해서 다른 페이지로 연결 될 수 있다.
이 하이퍼 텍스트 문서를 만들 수 있는 문서 포맷이 바로 하이퍼 텍스트 마크업 언어(HTML : HyperText MakrUp Language)의 약자인 HTML포맷이다.

웹 서버란 특정 포트에서 대기하며 사용자의 HTTP요청에 HTTP응답을 전송하는 서버를 말하고 이떄 응답은 일반적으로 HTML 문서를 전송
	웹 브라우저 -> HTTP 요청 -> 웹서버
	웹 브라우저 <- HTTP 답변 <- 웹서버


// HTTP쿼리 인수 사용하기
쿼리 인수는 URL끝에 붙여넣는 인수를 말하는 것(GET)
	Http://localhost?id=1&name=abcd

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}

	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3003", nil)
}

// ServeMux 인스턴스 이용하기
// ServeMux를 이용하면 핸들러에 다양한 기능을 추가 하기가 쉬움
// Mux? multiplexer(멀티 플렉서)의 약자로 여러 입력중 하나를 선택해서 반환하는 디지털 장치
// 웹서버는 각 URL에 해당하는 핸들러들을 등록한 다음 HTTP요청이 왔을 때 URL에 해당하는 핸들러를 선택해서 실행하는 방식
// 이 핸들러를 선택하고 실행하는 구조체이름이 MUX를 제공한다고 해서 ServeMux라고 부른다.
// 비슷한 의미인 "라우터"라고 말하기도함
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListenAndServe(": 3003", mux)
}

// 파일서버
웹 서버로 파일을 제공하는 방법
이미지같은경우 url같은거로 배포하는것을 말함
파일 읽어오는 법
http://localhost:3003/gopher.jpg 로 불러온다.
func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(": 3003", nil)
}
실제 웹 서비스에서는 파일을 웹서버에서 직접 전달하는 방식 대신 대부분은 콘텐츠 전송 네트워크(CDN Content delivery network, CDN) 서비스를 이용하는 방식으로 제공
이 서비스를 이용하면 파일을 사용자에서 가장 가까운 데이터 센터에서 바로 제공하기 떄문에 매우 빠르게 파일 데이터를 제공할 수 있다.

// 웹서버 테스트 코드 작성(main_test.go참조)
// Handler Instance 작성
func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	return mux
}

func main() {
	http.ListenAndServe(":3003", MakeWebHandler())
}

// JSON 데이터 전송
JSON(Java Script Object Notation)의 약자
자바스크립트에서 오브젝트를 표현하는 방법으로 사용되는 포맷

규칙
1. 오브젝트 시작은 {로 표기하고 }로 종료한다
2. 필드는 "key": value형태로 표기한다.
3. 각 필드는 , 로 구분한다.
4. 배열은 []로 표기한다.
5. 문자열은 ""로 묶어서 표기한다.

encoding.json패키지를 통해 구조체를 json데이터로 변환하고 다시 json데이터를 구조체로 변환 할 수 있음
type Student struct {
	Name  string
	Age   int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa", 16, 87}
	data, _ := json.Marshal(student)                   //Json을 []byte로 변환
	w.Header().Add("content-type", "application/json") //Json 포맷임을 표시
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":3003", MakeWebHandler())
}


// HTTPS 웹 서버 만들기
HTTPS를 지원하는 웹서버
HTTP는 보안을 염두해 두지 않고 설계된 프로토콜이므로 모든 요청이 일반 문자열이라 보안에 매우 취약
특히 스니핑과 같은 해킹으로 HTTP전문을 볼 수 있어 비밀번호, 개인정보와 같은 중요 데이터를 보호하지 못함
HTTPS는 기촌 HTTP요청과 응답을 ⭐️⭐️⭐️공개키 암호화 방식을 사용해서 암호화한 프로토콜
패킷이 암호화 되기 때문에 해커가 스니핑을 해도 어떤 내용인지 알수 없어서 비밀번호나 개인정보가 노출되지 않음

⭐️⭐️⭐️공개키 암호화 방식
공개키와 비밀키 두 가지 키를 생성해서 공개키는 클라이언트에 알려주고 비밀키는 서버에 비공개 상태로 놔둔다.
클라이언트에서 HTTPS요청을 보낼 때 공개키로 암호화 하고 서버는 비밀키를 이용해서 다시 원문으로 돌리는 복호화를 한다.
공개키 암호화 방식은 암호화와 복호화에 쓰이는 키가 서로 다르기 때문에 비대칭 암호화 방식이라고 한다.
HTTPS서버를 실행하려면 인증서와 비밀키를 생성해야하고 인증서는 공식 인증기관에서 발급해야하지만 개인 서버의 경우
발급이 되지 않으므로 openssl을 이용해서 인증서를 발급할 수 있다.

⭐️⭐️⭐️ 공개키암호화 방식을 도입했음에도 불구하고 또 인증서를 이용해 인증을 받는 이유
공개키 암호화 방식은 통신 내용이 암호화 되기 때문에 해커가 통신 내용을 감청해도 안전하다.
그러나 또 공인 기관의 인증을 받는 이유는 해커가 웹 사이트를 가장할 수 있기때문. 이것을 피싱 사이트라고한다.
예를 들어 겉모양을 은행 사이트와 똑같이 만든 뒤 사용자의 비밀번호와 같은 개인정보를 입력하라고 요청 할 수 있다.
이떄 HTTPS 프로토콜을 지원하기 위해서 해커가 만든 공개키를 클라이언트에 보내서 암호화 하라고 할 수 있다.
만약 클라이언트에서 이 웹사이트를 신뢰 할 수 있는지 여부를 모른다면 해커가 준 공개키로 패킷을 암호화 하게 되고
해커는 자신이 가진 비밀키로 이것을 복호화해서 알 수 있게 된다. 이것을 막기 위해서 별도 외부 공인기관을 통해 신뢰 할 수 있는 웹사이트인지 인증을 한다.

웹사이트 공개키는 외부 공인기관의 비밀키로 다시 암호화 되어서 공인기관내 저장된다. 사용자는 인증서를 받게 되고,
그 인증서에는 웹 사이트 공개키 대신 인증정보와 인증기관의 공개키가 들어있어서 인증기관을 통해 웹 사이트 공개키를 요청하게 된다.
그럼으로서 클라이언트는 안전하게 인증된 웹 사이트의 공개키를 받게 된다.
따라서 피싱 사이트라고 하더라도 해커가 제공한 공개키가 아닌 공인 인증기관에 저장된 웹 사이트 본래의 공개키로 암호화 되어서 해커가 볼 수 없게된다.

클라이언트	<- 인증정보, 인증기관의 공개키 <- 웹사이트
클라이언트	-> 웹 사이트 공개키 요청 	 -> 인증기관
클라이언트	<- 웹 사이트의 본래 공개키	  <- 인증기관

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello World")
	})

	// port, 인증서, 비밀키
	err := http.ListenAndServeTLS(":3003", "localhost.crt", "localhost.key", nil)

	if err != nil {
		log.Fatal(err)
	}
}
*/

/* 29. RESTful API 서버 만들기
//RESTful API
REST는 자원을 이름으로 구분하여 자원상태(정보를 주고 받는 소프트웨어 아키텍쳐)
RESTful API는 REST 규약을 따르는 API
웹서버에서는 URL과 HTTP 메서드로 데이터와 동작을 정의하는 방식을 의미.

REST : Representatlonal State Transfer 직역하면 표현식으로 데이터를 전송한다.

예를들어 웹서버에서 학생 데이터를 가져오는 URL의 예
GET https://somesite.com/getstudent.aspx?id=3

GET https://somesite.com/students/3
밑의 도메인이 더 자기 표현적이어서 어떻게 동작하는지 알기가 쉽다.

// ⭐️⭐️⭐️ HTTP 메서드
GET, POST, PUT, PATCH, DELETE 같은 메서드를 지원.
이것을 사용해 데이터에 대한 동작을 정의한다.
GET -> /student	-> 전체 학생 데이터 반환
GET -> /student/id -> id에 해당하는 학생 데이터 반환
POST -> /students -> 새로운 학생 등록
PUT -> /students/id -> id에 해당하는 학생 데이터 변경
DELETE -> /students/id -> id에 해당하는 학생 데이터 학제

GET	리소스 조회
POST	리소스 생성
PUT	리소스 생성/전체 교체
PATCH	리소스 부분 수정
DELETE	리소스 삭제
HEAD	리소스 메타데이터 조회
OPTIONS	서버 지원 메서드 확인
TRACE	요청이 서버에 도달했는지 확인
CONNECT	터널 연결

이와같이 URL과 메서드의 조합으로 데이터와 동작을 정의 할 수 있다.

가장 큰 특징들을 정리하면
1. 자기 표현적인 URL : URL만으로도 어떤 데이터에 대한 요청인지 알 수 있다.
2. 메서드로 행위 표현 : 메서드로 데이터에 대한 행위를 표현한다. URL과 메서드 조합으로 데이터에 대한 조작을 정의
3. 서버/클라이언트 구조 : 서버는 데이터 제공자로 존재하고 클라이언트는 데이터 사용자로 동작한다.
프론트와 백엔드로 분리하고 백엔드는 데이터만 제공하고 프론트에서 데이터를 처리하고 화면에 표시하는 역할을 함
4. 무상태 : 서버는 클라이언트의 상태를 유지하지 않음. 서버가 상태를 보관할 필요가 없기때문에 서버를 손쉽게 교체할 수 있어서 빠른 장애 대응이나 분산처리에 유용
5. 캐시 처리 : REST 구조로 서버가 단순해져서 더 쉽게 캐시 정책을 적용해서 성능을 개선

// RESTful API 서버 만들기
gorilla/mux 패키지를 이용하면 편하다
type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
	return mux
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func main() {
	http.ListenAndServe(":3003", MakeWebHandler())
}

// 특정 학생 데이터 반환하기
type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
	return mux
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)

	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // id를 가져옴
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func main() {
	http.ListenAndServe(":3003", MakeWebHandler())
}
// 학생 데이터 추가와 삭제
type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	mux.HandleFunc("/students/{id:[0-9]+}", DeleteStudentHandler).Methods("DELETE")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
	return mux
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)

	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // id를 가져옴
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(students, id)
	w.WriteHeader(http.StatusOK)
}
func main() {
	http.ListenAndServe(":3003", MakeWebHandler())
}

// RESTful API로의 발전
웹서버는 예전에는 완전한 HTML문서를 만들어서 반환했고 웹 브라우저가 하는 일은 웹 서버로 받아온 문서를 화면에 그리기만 하면 되었다.
이 방식을 서버 사이드 렌더링 방식이라고 함

	웹 브라우저 -> 웹 페이지 요청 -> 웹서버
	웹 브라우저 <- 전체 HTML문서 <- 웹서버

서버 사이드 렌더링 방식은 웹 브라우저에서 URL을 입력하면 웹서버에 웹페이지를 요청하고
웹서버는 전체 페이지 HTML을 만들어서 제공하는 방식
www에 대한 수요가 폭발적으로 증가하면서 서버사이드렌더링 방식의 문제점이 드러났다.

가장 큰 문제는 웹 서버로 성능 부하가 집중된다는 점. 요청이 많아지면 웹 서버 성능이 저하되어 느리게 반응하는 문제가 발생
사용자들은 빠른 반응에 대한 요구치가 높아져 다른 방식의 웹서버가 필요해짐

다양한 환경에 대한 유연한 대응이 힘듬
스마트폰이 보급되면서 컴퓨터없이도 어디서든 웹페이지에 접근하는 시대가 됨
사용자 환경이 다양해지면서 다양한 크기의 디스플레이에서 웹페이지가 생성되야했고 웹 브라우저뿐 아니라 단말기나 다양한 디바이스에서 웹 요청이 증가
웹서버가 HTML 문서를 모두 만드는 상황에서 다양한 환경에 대한 유연한 대응이 힘들었음

이떄 AJAX(비동기 자바스크립트)와 같은 동적 웹 기술이 발전하면서 서버에서 모든 HTML문서를 만드는게 아닌 CDN(Content delivery network 콘텐츠 전송 네트워크)에서
빠르게 껍데기 문서만 제공하고 내용은 웹 서버에서 데이터를 읽어와서 채우는 클라이언트 렌더링 방식으로 변화함

		웹 브라우저	-> HTTP요청 -> CDN(콘텐츠 전송 네트워크)
		웹 브라우저	<- 페이지 템플릿 <- CDN(콘텐츠 전송 네트워크)

		웹 브라우저 -> 페이지에 채울 데이터 요청 -> 웹 서버(데이터 제공자)
		웹 브라우저 <- 데이터 <- 웹 서버(데이터 제공자)

클라이언트인 웹 브라우저에서 URL이 입력되면 먼저 가까운 CDN으로부터 빠르게 템플릿 HTML문서를 받아옴
이 HTML문서는 완성된 형태가 아닌 중간 중간 데이터가 비어있는 형태로 빠르게 전송
CDN은 데이터를 채울 필요 없이 단순한 파일 서버형태로 빠르게 사용자에게 일정한 템플릿 HTML문서를 제공 할 수 있음
그 뒤 웹 브라우저에서 AJAX와 같은 동적 웹 기술을 이용해 페이지를 완성시킬 데이터를 웹 서버로 요청하게 된다
이때 웹서버는 순수하게 대ㅔ이터 제공자로 동작하고 주로 JSON포맷으로 데이터를 제공함
이렇게 데이터를 받으면 클라이언트에서 HTML문서를 그때그때 채워서 페이지를 완성하게됨

그래서 과거에는 웹페이지에 접속하면 화면 전체가 한꺼번에 그려졌지만 요즘에는 첫화면은 중간 중간 빈 형태로 빠르게 그려지고 점차 내용이 채워지는 것

클라이언트 렌더링 방식에서 웹 서버는 데이터 제공자로서 동작하기 떄문에 범용적인 인터페이스가 필요해졌음
이때 REST가 등장하면서 어떤환경에서도 URL,HTTP메서드만으로 데이터와 동작을 표현하게 됨으로서 범용성을 가질수 있게 됨

덕분에 네이버, 페이스북같은 거대 웹 서비스에서 수많은 웹 서버가 범용적인 인터페이스를 이용해 유기적으로 연결됨
성능 부하를 여러 웹서버로 분산해 더욱 빠르며 일부 서버에 장애가 발생해도 전체 서비스는 중단되지 않음
기능 변경시에도 전체 웹 서버를 변경하지 않고 일부만 변경하면되어 유지보수또한 쉬워짐
하지만 관리해야하는 웹 서버수가 급격히 늘어나서 사람이 일일이 관리가 안되게 됨
각종 자동 관리 툴이 탄생했고 다양한 툴을 능숙하게 다룰 인력이 필요해짐
이 인력을 DevOps라고 한다.

// Gin으로 서버만들기
// Gin 프레임 워크를 이용한 서버. 고에서 가장 대표적인 프레임워크

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score,omitempty"`
}

var students map[int]Student
var lastId int

func SetupHandlers(g *gin.Engine) {
	g.GET("/students", GetStudentsHandler)
	g.GET("/student/:id", GetStudentHandler)
	g.POST("/students", PostStudentHandler)
	g.DELETE("/students/:id", DeleteSttudentHandler)

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)
	c.JSON(http.StatusOK, list)
}

func GetStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}

func PostStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id:%d", lastId)
}

func DeleteSttudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	delete(students, id)
	c.String(http.StatusOK, "success to delete")
}

func main() {
	r := gin.Default()
	SetupHandlers(r)
	r.Run(":3003")
}
*/

/* 30. gnet과 gRPC로 채팅 앱 만들기
gnet : ChatApp폴더 참조
gRPC : 구글에서 만든 오픈 소스 원격 프로시져 콜 프레임워크
네트워크를 통해서 다른 컴퓨터에서 원하는 함수를 실행하는 것

프로토 버퍼
gRPC는 메세지의 직렬화를 위해 프로토버퍼 컴파일러를 사용.
직렬화란 구조체형태의 데이터를 하나의 바이너리 배열로 바꾸는 것.
역직렬화는 바이너리 배열을 다시 구조체로 바꾸는 것.
그래서 gRPC를 사용하려면 프로토 버퍼 컴파일러를 설치해야한다.
이후는 grpcchat폴더 참조
*/

/* A.문법 보충 수업
// 배열과 슬라이스
// 슬라이스를 만드는 여러가지 방법
func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice1 []int = array[1:5]           // 배열 슬라이싱
	var slice2 []int = slice1[1:8:9]        // 슬라이스 슬라이싱
	var slice3 []int = make([]int, 5)       // make()
	var slice4 []int = make([]int, 0)       // 길이 0인 슬라이스
	var slice5 []int = []int{1, 2, 3, 4, 5} // 초기화
	var slice6 []int                        // 기본값은 nil

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2, cap(slice2))
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	if slice4 != nil {
		fmt.Println("slice4 is not nil")
	}

	fmt.Println("slice5", slice5)
	fmt.Println("slice6", slice6)

	if slice6 == nil {
		fmt.Println("slice6 is nil")
	}
}
유용한 go 명령어
bug : Go 언어 자체의 버그를 리포트 할 수 있는 사이트를 브라우저로 접속. 버그 리포트를 시작
build : 패키지를 컴파일
clean : 컴파일 시 생성되는 패키지 목적 파일들을 삭제
doc : 패키지 문서를 출력. 특정 패키지 설명을 볼 때 유용
env : Go 환경 설정을 출력
fix : 오래된 API를 사용하는 Go 프로그램을 찾아서 새로운 API로 업데이트
fmt : 패키지를 리포맷한즌 gofmt툴을 실행. Go 코딩 규약에 맞춰 소스코드를 수정
generate : 패키지 파일 안에 파일 생성 절차가 정의되어있으면 그에 따라서 go 파일을 생성
get : 현재 모듈 패키지 목록에 패키지를 추가하고 다운받음
install : 컴파일한 뒤 결과를 GOPATH/bin 경로에 설치
list : 패키지나 모듈 목록을 출력
mod : 새로운 모듈을 만들거나 관리
run : 컴파일한 뒤 결과 프로그램을 실행. 실행파일을 생성하지 않음
test : 패키지를 테스트
tool : 특정 go 도구를 실행
version : 버전 출력
vet :  패키지내 버그로 의심되는 부분을 보고
*/
