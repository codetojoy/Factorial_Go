
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/codetojoy/config"
    "github.com/codetojoy/factorial"
    "github.com/codetojoy/util"
    "github.com/codetojoy/worker"
)

func processN(n int) {
    config := config.New(n)

    for i:= 2; i <= n; i++ {
        f := factorial.NewFactorial(i, config)
        f.Compute()
        fmt.Printf("TRACER f: %d -> %v \n", i, f.GetPureFactorsString())
    }
}

func searchChunk(chunkLow int, chunkHigh int, factorialIndex *factorial.FactorialIndex,
                    myChannel chan worker.Payload) {
    go worker.SearchChunk(chunkLow, chunkHigh, factorialIndex, myChannel)

    /*
    result := <- myChannel
    if result.Err == nil {
        for i := 0; i < len(result.Results); i++ {
            result := result.Results[i]
            fmt.Printf("TRACER %d! = %d! x %d!\n", result.A, result.B, result.C)
        }
    } else {
        log.Fatal(result.Err)
    }
    */
}

func reportResults(result worker.Payload) {
    if result.Err == nil {
        for i := 0; i < len(result.Results); i++ {
            result := result.Results[i]
            fmt.Printf("TRACER %d! = %d! x %d!\n", result.A, result.B, result.C)
        }
    } else {
        log.Fatal(result.Err)
    }
}

func search(n int) {
    config := config.New(n)
    factorialIndex := factorial.NewFactorialIndex(config)
    isDone := false
    chunkIndex := 0
    const chunkSize = 20

    const numChannels = 5
    channel1 := make(chan worker.Payload)
    channel2 := make(chan worker.Payload)
    channel3 := make(chan worker.Payload)
    channel4 := make(chan worker.Payload)
    channel5 := make(chan worker.Payload)

    for ! isDone {
        index := (chunkIndex % numChannels) + 1;
        channel := channel1
        if index == 2 {
           channel = channel2
        } else if index == 3 {
           channel = channel3
        } else if index == 4 {
           channel = channel4
        } else if index == 5 {
           channel = channel5
        }

        chunkLow, chunkHigh := util.GetChunk(chunkIndex, chunkSize, n)
        searchChunk(chunkLow, chunkHigh, factorialIndex, channel)

        if chunkHigh >= n {
            isDone = true
        }

        chunkIndex++
    }

    // wait on channels
    //
    // TODO: I'm not entirely convinced this is correct.
    // I think it possible for all channels to have nothing happening, and yet
    // the workload is not yet complete.
    // Stack O threads illustrate a dedicated 'quit' channel.
    // e.g. https://stackoverflow.com/questions/11117382
    hasStarted := false
    areChannelsDone := false

    fmt.Println("TRACER main selecting ...")
    for ! areChannelsDone {
        time.Sleep(100 * time.Millisecond)

        select {
        case result1 := <- channel1:
            hasStarted = true
            reportResults(result1)
        case result2 := <- channel2:
            hasStarted = true
            reportResults(result2)
        case result3 := <- channel3:
            hasStarted = true
            reportResults(result3)
        case result4 := <- channel4:
            hasStarted = true
            reportResults(result4)
        case result5 := <- channel5:
            hasStarted = true
            reportResults(result5)
        default:
            if hasStarted {
                fmt.Println("TRACER main channels done ...")
                areChannelsDone = true
            }
        }
    }

    fmt.Println("TRACER main sleeping ...")
    time.Sleep(3 * time.Second)
}

func main() {
    numArgs := len(os.Args)

    if numArgs > 1 {
        mode := os.Args[1]

        if mode == "process50" {
            n := 50
            processN(n)
        } else if mode == "search" {
            x := 121
            search(x)
        } else {
            fmt.Println("ERROR: unknown mode")
            os.Exit(1)
        }
    } else {
        fmt.Println("USAGE: requires mode param")
    }

    fmt.Println("Ready.")
}
