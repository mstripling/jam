package setup
import (
    "github.com/mstripling/jam/internal/async"
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/mstripling/jam/cmd/jam/main"
    "os"
    "os/exec"
//    "os/signal"
    "github.com/mstripling/jam/internal/io"
    "sync"
    "strings"
//    "syscall"
    "time"
)

func WaitForMonitor(){
    for{
        // iwconfig wlp4s0mon | awk '/Mode:Monitor/{print "Monitor"}'
        outputBytes := RunSimple("iwconfig wlp4s0mon | awk '/Mode:Monitor/{print \"Monitor\"}'")
        mode := strings.TrimSpace(string(outputBytes))
        fmt.Println("Current mode:", mode)

        if mode == "Monitor" {
            break
        }
        time.Sleep(500 * time.Millisecond)
    }
}

func Setup() string {
    RunSimple("apt install aircrack-ng")
    wifiCardNameBytes := RunSimple("iw dev | awk '$1==\"Interface\"{print $2}'")
    var wifiCardName = strings.TrimSpace(string(wifiCardNameBytes))
    command := "sudo airmon-ng start " + wifiCardName
    RunSimple(command)
    WaitForMonitor()
    fmt.Println("waitForMonitor() completed")
    wifiCardNameBytes = RunSimple("iw dev | awk '$1==\"Interface\"{print $2}'")
    wifiCardName = strings.TrimSpace(string(wifiCardNameBytes))
    return wifiCardName
}
