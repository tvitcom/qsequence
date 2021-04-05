package main

import (
    "fmt"
    "math/big"
    "crypto/rand"
)

func main() {
    xs := []uint64{1,2,3,4,5,6,7,8,9,10,11,12,12,14,15,16,17,18,19,20}

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
        xs = append(xs[:index], xs[index+1:]...)
        return result 
    }

    lenXs := len(xs)
    fmt.Println("lenXs:",lenXs)
    for i :=0 ;i < lenXs;i++ {
        fmt.Println(shuffleElement())
    }
}

