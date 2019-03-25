package goback

import (
	"log"
)

// 外部使用方式示例:
//	obj := goback.Regist("MainWnd")
//	go func() {
//		for {
//			_, ok := <-obj.BufCh
//			if !ok {
//				break
//			}
//		}
//    close(obj.BufCh)
//	}()
//  goback.Wait()



type BackServer struct {
	BufCh   chan []byte
	backWnd *BackWnd
}

var backsev *BackServer

func Regist(cppwndName string) *BackServer {
	backsev = new(BackServer)
	backsev.BufCh = make(chan []byte)

	backsev.backWnd = NewBackWnd(cppwndName)
	backsev.Push("goback for golang is OK")
	return backsev
}

func Wait() {
	backsev.backWnd.WaitMessage()
}

func (p *BackServer) Push(data string) {
	p.backWnd.SendSyn(data)
	log.Println("Push message is:", data)
}

func (p *BackServer) Accept(ldata interface{}) {
	log.Println("Accept message is:", ldata.(string))
	params := ldata.(string)
	p.BufCh <- ([]byte)(params)
}