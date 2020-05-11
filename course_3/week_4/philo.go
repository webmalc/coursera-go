package main

import (
	"fmt"
	"sync"
	"time"
)

const maxEatAtTime = 2
const numberOfMeals = 3
const numberPhilos = 5

type ChopS struct {
	sync.Mutex
}
type Philos struct {
	num, count      int
	leftcs, rightcs *ChopS
}

func (p Philos) eat(c chan *Philos, wg *sync.WaitGroup) {
	for i := 0; i < numberOfMeals; i++ {
		c <- &p
		if p.count < numberOfMeals {
			p.leftcs.Lock()
			p.rightcs.Lock()

			fmt.Println("starting to eat ", p.num)
			p.count++
			fmt.Println("finishing eating", p.num)
			p.rightcs.Unlock()
			p.leftcs.Unlock()
			wg.Done()
		}
	}
}

func host(c chan *Philos) {
	const sleep = 20
	for {
		if len(c) == maxEatAtTime {
			<-c
			<-c
			time.Sleep(sleep * time.Millisecond)
		}
	}
}

func main() {
	var i int
	var wg sync.WaitGroup
	c := make(chan *Philos, maxEatAtTime)

	wg.Add(numberPhilos * numberOfMeals)

	ChopSticks := make([]*ChopS, numberPhilos)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)
	}

	Philosophers := make([]*Philos, numberPhilos)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philos{i + 1, 0, ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	go host(c)
	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(c, &wg)
	}
	wg.Wait()
}
