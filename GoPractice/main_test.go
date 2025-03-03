package main

/* 26. 테스트와 벤치마크하기
// 3가지 표현 규약을 따라 테스트 코드를 작성. go test 명령으로 실행
//⭐️ 1. 파일명이 _test.go로 끝나애 한다.
//⭐️ 2. testing 패키지를 임포트 해야한다.
//⭐️ 3. 테스트 코드는 func TestXxxx(t *testing.T)형태여야 한다.
// 특정 테스트만 실행하는 코드는 go test -run 테스트명
// 테스트 코드 작성 방법

func TestSquare1(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Errorf("square(9) should be 81 but square(9) returns %d", rst)
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)
	if rst != 9 {
		t.Errorf("square(3) should be 81 but square(9) returns %d", rst)
	}
}

// 테스트를 돕는 외부 패키지
// stretchr/testfy 패키지 -> 테스트를 하고 테스트의 실패를 알릴 수 있는 함수를 제공.
// 코드가 간략하여 인기가 높음
func TestSquare1(t *testing.T) {

	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}
func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
// stretchr/testfy/assert 패키지에서 제공하는 유용한 함수
// Equal(기대값, 실제값) -> 기대값과 실제값을 비교하여 다를경우 테스트를 실패하고 메세지 출력
// Greater(e1, e2) -> e1이 e2보다 크지 않으면 테스트를 실패하고 메세지 출력
// len() -> 오브젝트의 항목 개수가 lenth와 다르면 테스트를 실패하고 메세지를 출력
// NotNilf() -> 오브젝트가 널이면 테스트를 실패하고 메시지를 출력합니다.
// NotEqualf() -> 기대값과 실제값을 비교하여 같을 경우 테스트를 실패하고 메세지 출력

//mock
// 모듈의 행동을 가장하는 목업 객체를 제공. 온라인 기능을 테스트할 떄 하위 영역인 네트워크 기능까지 모두 테스트하기는 힘들다.
// 그래서 네트워크 객체를 가장하는 목업 객체를 만들 때 유용하다.

// suite
// 테스트 준비 작업이나 테스트 종류후 뒤처리 작업을 쉽게 할 수 있도록 도와줌
// 예) 테스트에 특정 파일이 있어야 한다면 테스트 시작 전 임시 파일을 생성하고
// 테스트 종료 후 생성한 임시 파일을 삭제 해주는 작업을 만들 때 유용함

// 테스트 주도 개발
// 과거에 비해서 프로그램 규모가 커지고 고가용성에 대한 요구사항이 높아져 테스트의 중요성이 높아짐
// 가용성 -> 프로그램이나 웹 서비스가 얼마나 오랫동안 정상 동작하는가를 의미
// 현재 거의 모든 웹이나 프로그램은 상시 서비스를 기본으로 제공하므로 중간에 서비스가 멈추어선 안됨

// 블랙 박스 테스트
// 제품 내부를 오픈하지 않은 상태에서 진행되는 테스트
// 사용자 입장에서 테스트한다고 해서 사용성 테스트라고 하기도 함
// 프로그램 내부 코드를 검증하는 것이 아닌 프로그램을 실행하여 실행 동작을 검사하는 방식. 전문 테스터 QV,QA에서 보통 담당.
// 결합테스트와 비슷한 개념인 것 같다.

// 화이트 박스 테스트
// 화이트 박스 테스트는 프로그램 내부 코드를 직접 검증하는 방식
// 유닛테스트라고 부름(단위 테스트)
// 코드작성->테스트하고 버그발견->코드수정 이 일반적인방법이지만
// 테스트 케이스가 빈약하면 메인 시나리오에 의존해 테스트하고 여외 상황이나 경계체크가 무시되기 쉬움
// 테스트 통과를 목적으로 하는 형식적인 테스트 코드를 작성하기가 쉽다.

// 위와 같은 문제점을 해결하기 위해 테스트 주도개발이 대안이 될 수 있다.
// 먼저 테스트 코드를 작성
// 테스트를 성공시키는 코드를 작성해서 테스트를 성공시킴
// 개선작업을 통해 코드를 개선
// 개선은 SOLID원칙에 입각해 진행 -> 리팩터링
// 이 과정을 반복함

// 테스트 주도 개발의 이점
// 테스트 코드가 자연적으로 촘촘해짐.
// 작은 목표 설정 -> 실패 -> 달성 -> 달성 강화-> 새로운 작은 목표설정이라는 과정으로 개발이 재미있어진다.
// 단점은 나무위키를 참고

// 벤치마크
// 코드 성능을 검사하는 벤치마크 기능도 지원.
// 벤치마크 표현규약
// 1.파일명이 _test.go으로 끝나야한다.
// 2.testing 패키지를 임포트해야한다.
// 3.벤치마크 코드는 func BenchmarkXxxx(b *testing.B)형태여야한다.

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibonacci1(2) should be 2")
	assert.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibonacci2(2) should be 2")
	assert.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

// 벤치 마크 코드 작성의 예
// 실행은 go test -bench . 로 한다.(패키지내 모든 벤치마크를 실행하는 코드
func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
*/

/* 27. 웹서버 테스트 코드
func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}

// Json Test
func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/student", nil) //student 경로 테스트

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	student := new(Student)
	err := json.NewDecoder(res.Body).Decode(student) // 결과 변환
	assert.Nil(err)                                  // 결과 확인
	assert.Equal("aaa", student.Name)
	assert.Equal(16, student.Age)
	assert.Equal(87, student.Score)
}
*/

/* 28. RESTful API 테스트 코드
func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("aaa", list[0].Name)
	assert.Equal("bbb", list[1].Name)
}

// 특정 데이터 반환 테스트 코드
func TestJsonHandler2(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("aaa", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("bbb", student.Name)
}
func TestJsonHandler3(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students", strings.NewReader(`{"Id":0, "Name":"ccc", "Age":15, "Score":78}`))

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("ccc", student.Name)
}

func TestJsonHandler4(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("bbb", list[0].Name)
}
*/
