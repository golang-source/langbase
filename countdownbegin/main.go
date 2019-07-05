package main

import (
	"sync/atomic"
	"fmt"
	"time"
)

type CountDownBegin struct{
	num uint32
	sigMap  map[uint32] chan Sig
	readNn uint32
	finishChan chan int  //fire执行完毕，会通知完成

}

func (this *CountDownBegin)SetSize(s uint32)*CountDownBegin {
	this.num = s
	return this
}

type Sig struct {
	SigCode int
}


func (this *CountDownBegin) WaitFinish(){

	<-this.finishChan

}

func newCountDownBegin() *CountDownBegin {
	ret:=CountDownBegin{}
	ret.sigMap = make(map[uint32]chan Sig)
	ret.num=0
	ret.readNn=0
	ret.finishChan =make(chan int,1)
	return &ret
}
const CLIENT_READY =1
const SVR_SEND_CALLBACK_EXECUTE =1

func newSigChan() chan Sig {
	return make(chan Sig,1)
}
var uuid uint32=0
func getUUID() uint32 {
	return atomic.AddUint32(&uuid,1)
}
func (this *CountDownBegin)save(s chan Sig){
	//这里不是 ，map的内部的索引计算部分，并非 支持并发 （map 非并发map,虽然外部已经有并发支持getUUID（））
    this.sigMap[getUUID()] = s
}
func (this *CountDownBegin) switchToReadyQueue(sigChan <-chan Sig)  {
	res:=atomic.AddUint32(&this.readNn,1)
	if res==this.num &&res!=0{
		this.fire()
	}
}
//如果只有 reg 和 fire 那么不好控制
//所以分成，reg  客户端 ready  ，服务器端 设定 ready queue ，如果read queue==num ,那就执行 fire
//需要客户端更多的步骤和特征来完成控制协调工作
func (this *CountDownBegin) reg() chan Sig {
		s:= newSigChan()
		this.save(s)
		//ready 检查   //go routine 能否被调度是全局对它的调度问题,函数范围内无法决定它的调度权利
        var alarm =make (chan byte ,1)
		go func(sigChan chan Sig,alar chan byte) {
				select {
				    case s:=<-sigChan:
				    	if s.SigCode==CLIENT_READY{
				    		this.switchToReadyQueue(sigChan)
				    		fmt.Println("client say am ready")
							alar<-1
						}
				}
		}(s,alarm)

         <-alarm

	return s
}


//通知他们去做事
func (this * CountDownBegin)fire()  {

	for _, recv := range this.sigMap {
		s:=Sig{SigCode:SVR_SEND_CALLBACK_EXECUTE}
		recv<-s
	}
	fmt.Println("fire finished ")
	this.finishChan<-1
}

func waitForRun(say chan Sig,alarmQuit chan int,runCallback func())  {
	//for {
		select {
			case s:= <-say:
				if s.SigCode==SVR_SEND_CALLBACK_EXECUTE{
					runCallback()
					alarmQuit <-1
				}
		//default:// 没有default 就没有必要外层的 for ,因为 case s:= <-say: 本身就是执行一次，并且没有信号会堵塞的
		}
	//}
}

func startThread(say chan Sig )  {
	var alarmQuit = make(chan int,1)
	go waitForRun(say, alarmQuit,func() {
		fmt.Println("frondend is start running")
	})
//我准备好了
	say <- Sig{SigCode:CLIENT_READY}
}

const WORKER = 10
func main() {
	var begin =newCountDownBegin()
	begin.SetSize(WORKER)
	for i:=0;i<WORKER;i++{
		go startThread(begin.reg()) //两边的系统位置一个双向的 对话数据通道 即可，无需依赖别的
	}
	//begin.WaitFinish()
	time.Sleep(time.Second*120)
}