package io
import (
    "github.com/mstripling/jam/internal/async"
    "bufio"
    "encoding/json"
    "fmt"
    "github.com/mstripling/jam/cmd/jam/main"
    "os"
    "os/exec"
//    "os/signal"
    "github.com/mstripling/jam/internal/setup"
    "sync"
    "strings"
//    "syscall"
    "time"
)

func WriteToFile(output []byte, filename string) error{
    //check if exists
    _, err:= os.Stat(filename)
    fileExists := !os.IsNotExist(err)
    
    if fileExists {
        if err := os.Remove(filename); err != nil{
            return err
        }
    }

    file, err := os.Create(filename)
    if err != nil {
        return err
    }

    defer file.Close()
    _,err = file.Write(output)
    if err != nil {
        return err
    }
    
    if fileExists {
        fmt.Printf("Output overwritten in %s\n", filename)
    } else {
        fmt.Printf("Output saved to %s\n", filename)
    }

    return nil
}

//************************************************************************************\\
type NetworkEntry struct{
    BSSID string 'json:"BSSID"'
    Channel string 'json:"Channel "'
}

type DeviceEntry struct{
    MAC string 'json:"MAC"'
    Signal string 'json:"Signal"'
}

//Global maps
var networks = make(map[string]NetworkEntry)
var devices = make(map[string]DeviceEntry)



func processOutputLine(line string){
    //logic to parse lines coming from outputChan
    // output from terminal -> scanner -> go func -> outputChan
    // process the data from both goroutines as the correct struct, then pass to new function?
}

func storeProcessedData(data struct{}){
    //store the data in the correct map to be used to deauth
}

func deauth(wifiCardName string, NetworkEntry struct{}, DeviceEntry struct{}){
    //send out attacks!
}

//************************************************************************************\\


