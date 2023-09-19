package main

import (
    "github.com/mstripling/jam/internal/async"
    "fmt"
//    "github.com/mstripling/jam/internal/cook"
//    "os/signal"
    "github.com/mstripling/jam/internal/setup"
    "sync"
//    "syscall"
    "time"
)



func main(){
    wifiCardName := setup.Setup()
    fmt.Println("setup completed")
    washScan := fmt.Sprintf("sudo wash -i %s -j > /tmp/wps.txt", wifiCardName)
    // phase 1
    fmt.Println("washScan created")
    
    outputChan := make(chan string, 100)
    stopChan := make(chan struct{})
    fmt.Println("output and stopchan created")
    
    var wg sync.WaitGroup

    wg.Add(1)
    //async.AsyncSimpleDelay(washScan, 10)   
    go async.Async(washScan, 0, &wg, outputChan,stopChan)
    timer := time.NewTimer(30 * time.Second)
    defer timer.Stop()

    go func() {
        select{
        case <-timer.C:
            close(stopChan)
            fmt.Println("stopChan closed after 30 sec")
        }
    }()
    /* signalChan no longer needed to end washScan
    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    fmt.Println("signal chan created")

    
    go func(){
        for{
            select{
            case<-signalChan:
                close(stopChan)
                fmt.Println("stopChan closed")
                return
            case output := <-outputChan:
                fmt.Println(output)
            }
        }
    }()
        */
    wg.Wait()
    fmt.Println("post wg.Wait()")
    async.Sync("cp /tmp/wps.txt wps.txt")
    fmt.Println("wps.txt copied here")
    //clean up
    // hard coded wifi card name
    // listen to a key and execute code below
    toManagedCommand := fmt.Sprintf("sudo airmon-ng stop wlp4s0mon") //%s", wifiCardName)
    async.Sync(toManagedCommand)
    async.Sync("clear")

    fmt.Println("All workers done.")
}
//commands := []string{
//    washScan,
//    "pwd"}
//var wg sync.WaitGroup
//
//for _, cmd := range commands {
//    wg.Add(1)
//    go async.Sync(cmd, &wg)
//}
//wg.Wait()
