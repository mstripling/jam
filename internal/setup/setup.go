package setup

import (
    "fmt"
    "github.com/mstripling/jam/internal/async"
    "strings"
    "time"
)

func WaitForMonitor(){
    for{
        // iwconfig wlp4s0mon | awk '/Mode:Monitor/{print "Monitor"}'
        outputBytes := async.Sync("iwconfig wlp4s0mon | awk '/Mode:Monitor/{print \"Monitor\"}'")
        mode := strings.TrimSpace(string(outputBytes))
        fmt.Println("Current mode:", mode)

        if mode == "Monitor" {
            break
        }
        time.Sleep(500 * time.Millisecond)
    }
}

func Setup() string {
    async.Sync("apt install aircrack-ng")
    wifiCardNameBytes := async.Sync("iw dev | awk '$1==\"Interface\"{print $2}'")
    var wifiCardName = strings.TrimSpace(string(wifiCardNameBytes))
    command := "sudo airmon-ng start " + wifiCardName
    async.Sync(command)
    WaitForMonitor()
    fmt.Println("waitForMonitor() completed")
    wifiCardNameBytes = async.Sync("iw dev | awk '$1==\"Interface\"{print $2}'")
    wifiCardName = strings.TrimSpace(string(wifiCardNameBytes))
    return wifiCardName
}
