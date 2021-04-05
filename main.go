package main

import (
	"crypto/rand"
	"fmt"
	"github.com/recoilme/pudge"
	"log"
	"math/big"
)

func main() {
	queueNumbers := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 14, 15, 16, 17, 18, 19, 20}
	qQueueQuestionNumbers, stringQueue := getQueueQuestionNumbers(queueNumbers, false)

	dbFile := "../data/pudgedb"
	cfg := &pudge.Config{
		SyncInterval: 1,
		// StoreMode: 2,
	} //disable every second fsync
	db, err := pudge.Open(dbFile, cfg)
	if err != nil {
		log.Panic(err)
	}
	defer db.DeleteFile()

	if ok, _ := db.Has(stringQueue); !ok {
		pudge.Set(dbFile, stringQueue, qQueueQuestionNumbers)
	} else {
		log.Println("Sequence last occured before. Try again.")
	}

	log.Println(stringQueue)
	log.Println(qQueueQuestionNumbers)
}

func getQueueQuestionNumbers(xs []uint64, debug bool) ([]uint64, string) {
	queueString := ""
	var qQueueNumbers []uint64
	shuffleElement := func() uint64 {
		// get rand number for index
		b := len(xs)
		big := big.NewInt(int64(b))
		num, ok := rand.Int(rand.Reader, big)
		if ok != nil {
			panic(ok)
		}
		index := uint64(num.Int64())
		result := xs[index]
		queueString = queueString + fmt.Sprintf("%v,", xs[index])
		xs = append(xs[:index], xs[index+1:]...)
		return result
	}

	lenXs := len(xs)
	next := uint64(0)

	if debug {
		log.Println("lenXs:", lenXs)
	}

	for i := 0; i < lenXs; i++ {
		next = shuffleElement()
		if debug {
			log.Println(next, queueString)
		}
		qQueueNumbers = append(qQueueNumbers, next)
	}

	return qQueueNumbers, queueString
}
