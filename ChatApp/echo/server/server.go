/* 일반적인 서버를 만드는 예

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"

// 	"github.com/panjf2000/gnet/v2"
// )

// type echoServer struct {
// 	gnet.BuiltinEventEngine

// 	eng       gnet.Engine
// 	addr      string
// 	multicore bool
// }

// func (es *echoServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
// 	log.Printf("client connected. address:%s", c.RemoteAddr().String())
// 	return nil, gnet.None
// }

// func (es *echoServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
// 	log.Printf("client disconnected. address:%s", c.RemoteAddr().String())
// 	return gnet.None
// }

// func (es *echoServer) OnBoot(eng gnet.Engine) gnet.Action {
// 	es.eng = eng
// 	log.Printf("echo server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
// 	return gnet.None
// }

// func (es *echoServer) OnTraffic(c gnet.Conn) gnet.Action {
// 	buf, _ := c.Next(-1)
// 	c.Write(buf)
// 	return gnet.None
// }

// func main() {
// 	var port int
// 	var multicore bool

// 	// flag패키지를 통해 실행 인수를 읽어옴.
// 	flag.IntVar(&port, "port", 9000, "--port 9000")
// 	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
// 	flag.Parse()

// 	// gnet.Run을 호출해서 서버를 실행
// 	// 첫번째 인수는 이벤트 핸들러로 gnet내부에서 이벤트들이 클라이언트 접속, 연결해재, 데이터 수신등이 발생하면 그에 맞는 메서드를 호출해줌
// 	// 이 이벤트 핸들러 인터페이스를 통해 서버에서 데이터를 수신하거나 클라이언트 접속/해제를 알 수 있음

// 	// 두번쨰 인수는 protoAddr로 서버가 어떤 프로토콜을 통해서 통신하게 되는지, 그 주소는 어디에 바인딩하는지 나타냄
// 	// 세번쨰 인수는 옵션 리스트로 멀티코어 옵션을 적용하여 여러 CPU코어를 사용해서 동작하게함
// 	// 더 빠른 성능을 가지게 되나, 멀티 스레딩 환경에서 동작 하기 떄문에 메모리 자원점유에 주의 해야함
// 	echo := &echoServer{
// 		addr:      fmt.Sprintf("tcp://%d", port),
// 		multicore: multicore,
// 	}
// 	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
// }

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/panjf2000/gnet/v2"
)

type chatServer struct {
	gnet.BuiltinEventEngine

	// 연결된 커넥션을 보관하는 앱
	cliMap sync.Map
}

func (cs *chatServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("client connected. address:%s", c.RemoteAddr().String())
	// 새로운 연결이 되면 커넥션을 맵에 보관
	cs.cliMap.Store(c, true)
	return nil, gnet.None
}

func (cs *chatServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("client disconnected. address:%s", c.RemoteAddr().String())
	// 연결이 해제되면 맵에서 삭제
	if _, ok := cs.cliMap.LoadAndDelete(c); ok {
		log.Printf("connection removed")
	}
	return gnet.None
}

func (cs *chatServer) OnBoot(eng gnet.Engine) gnet.Action {
	log.Printf("chat server is listening\n")
	return gnet.None
}

func (cs *chatServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1)
	// 데이터 수신 시 모든 커넥션에 데이터를 전송한다.
	cs.cliMap.Range(func(key, value any) bool {
		if conn, ok := key.(gnet.Conn); ok {
			conn.AsyncWrite(buf, nil)
		}
		return true
	})
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()

	chat := &chatServer{}
	log.Fatal(gnet.Run(chat, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
