package main

import (
    "fmt"
    "math/big"
    // "strconv"
    "crypto/rand"
)

const (
    QUERY_CHAIN_LENGTH = 20 // 20 вопросов может быть максимально задано
)
type (
    QueryElement struct {
        QueryChain string
        Tries uint64
    }
)

func main() {

    queueNumbers := []uint64{1,2,3,4,5,6,7,8,9,10,11,12,12,14,15,16,17,18,19,20,}
    

    qQueueQuestionNumber, stringQueue := getQueueQuestionNumbers(queueNumbers, false)
    fmt.Println(stringQueue)
    fmt.Println(qQueueQuestionNumber)
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
    fmt.Println("lenXs:",lenXs)

    for i :=0; i < lenXs; i++ {
        next = shuffleElement()
        if (debug) {
            fmt.Println(next, queueString)
        }
        qQueueNumbers = append(qQueueNumbers, next)
    }

    return qQueueNumbers, queueString

}