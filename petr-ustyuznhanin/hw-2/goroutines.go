package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex //добавляется для блокировки участков кода
	balance int
)

func deposit(value int, wg *sync.WaitGroup) {
	fmt.Printf("Текущий баланс %d. Добавляем на счет %d\n", balance, value)

	mutex.Lock() //блокировка чтобы со значением balance могла работать только 1 функция
	balance+=value
	mutex.Unlock()
	wg.Done()
}

func credit(value int, wg *sync.WaitGroup) {
	fmt.Printf("Текущий баланс %d. Снимаем со счета %d\n", balance, value)

	mutex.Lock()
	balance-=value
	mutex.Unlock()
	wg.Done()
}

func main()  {
	fmt.Println("Start")

	balance = 1000
	var wg sync.WaitGroup

	wg.Add(2) //отсюда ждем 2 сигнала от wg.Done и только тогда подолжаем
	go deposit(500, &wg)
	go credit(400, &wg)
	wg.Wait() //код выполняется только до сюда пока не будет 2 сигнала wg.Done

	fmt.Printf("Окончательный баланс счета %d\n", balance)
}