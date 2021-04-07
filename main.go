package main

import (
	"crypto/rand"
	"fmt"
	"github.com/recoilme/pudge"
	"log"
	"math/big"
)

const repeatNumber uint64 = 20

func main() {
	startQueueNumbers := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	dbFile := "../data/pudgedb"
	cfg := &pudge.Config{
		SyncInterval: 1,
		StoreMode:    2,
	} //disable every second fsync
	db, err := pudge.Open(dbFile, cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.DeleteFile()

	for i := 0; i < 1000000; i++ {
		qQueueNumbers, stringQueue := getQueueNumbers(startQueueNumbers)

		if ok, _ := db.Has(stringQueue); !ok {
			pudge.Set(dbFile, stringQueue, qQueueNumbers)
			fmt.Println(stringQueue)
		} else {
			fmt.Println("Sequence last occured before.", qQueueNumbers, "Try again.")
		}
		qQueueNumbers = []uint64{}
		stringQueue = string("")
	}
}

func getQueueNumbers(xs []uint64) ([]uint64, string) {
	var debug bool = false

	if debug {
		fmt.Println("Debug: start func------------> xs", xs)
	}

	lenXs := len(xs)
	var lastElements = make([]uint64, lenXs)
	_ = copy(lastElements, xs)

	shuffleElement := func() uint64 {
		var result uint64

		if debug {
			fmt.Println("Debug: shuffleElements START-> lastElements:", lastElements)
		}

		// get rand number for index
		b := len(lastElements)
		big := big.NewInt(int64(b))
		index := uint64(0)

		num, ok := rand.Int(rand.Reader, big)
		if ok != nil {
			panic(ok)
		}

		index = uint64(num.Int64())

		if debug {
			fmt.Println("Debug: shuffleElements internal-> cycle:", "index:", index, "num:", num)
		}

		result = lastElements[index]
		if debug {
			fmt.Println("Debug: shuffleElements -> lastElements[index]:", lastElements[index])
		}
		//Удаляем элемент с массива
		lastElements = append(lastElements[:index], lastElements[index+1:]...)
		if debug {
			fmt.Println("Debug: shuffleElements STOP-> OUTPUT:", result, "lastElements:", lastElements)
		}
		return result
	}

	if debug {
		fmt.Println("lenXs:", lenXs)
	}

	var queueString string // For store string representative
	var qQueueNumbers []uint64
	for i := 0; i < lenXs; i++ {
		next := shuffleElement()
		if debug {
			fmt.Println("Debug:OUTER CYCLE-> i:", i, "next:", next, "|", "queueString:", queueString)
		}
		if next != 0 {
			qQueueNumbers = append(qQueueNumbers, next)
			queueString = queueString + fmt.Sprintf("%v,", next) //
		}
		next = uint64(0)
	}

	if debug {
		fmt.Println("Debug: end func. qQueueNumbers:", qQueueNumbers)
	}

	return qQueueNumbers, queueString
}
