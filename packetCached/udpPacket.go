package packetCached

import (
	"net"
	"sync"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"fmt"
	"log"
)


var Max_queue_length int = 10000
var Queue *QueueStruct


type UDPPacket struct{
	Buf []byte
	Octets int
	RouterAddr *net.UDPAddr
}

func QueueInit(){
	Queue = new(QueueStruct)
	Queue.Data_queue = lls.New()
	Queue.Locker = new(sync.Mutex)
}

//****************************************************
//Queue
//****************************************************
type QueueStruct struct{
	Data_queue *lls.Stack
	Locker *sync.Mutex
}

func (q *QueueStruct) PushQueue(buf []byte, octets int, routerAddr *net.UDPAddr) {
	q.Locker.Lock()
	defer q.Locker.Unlock()
	udp_packet := UDPPacket{buf, octets, routerAddr}

	if q.Data_queue.Size()>Max_queue_length {
		fmt.Println("Reader - the queue is bigger then max_queue_length ", q.Data_queue.Size(), "/", Max_queue_length)
		return
	}
	q.Data_queue.Push(udp_packet)
	if q.Data_queue.Size() > 10 {
		log.Print("queue size:", q.Data_queue.Size())
	}

}

func (q *QueueStruct) GetQueueData()(value interface{}, ok bool){
	q.Locker.Lock()
	defer q.Locker.Unlock()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("GetQueueData ERROR:", err)
		}
	}()

	if q.Data_queue != nil && q.Data_queue.Size()>0{
		data, err := q.Data_queue.Pop()
		return data, err
	}
	return nil, false
}
//****************************************************
//Queue   End
//****************************************************
