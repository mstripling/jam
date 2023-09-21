package async

import (
    "bufio"
    "fmt"
    "os/exec"
    "sync"
    "time"
)

type ErrMessage struct{
    command string
    err error
    notes string
}

func Check(e ErrMessage){
    if e.err != nil {
        fmt.Printf("Error %s: %v\nNotes: %s\n", e.command, e.err, e.notes)
        return
    }
}

func Async(command string, delay int, wg *sync.WaitGroup, outputChan chan<- string, stopChan <-chan struct{}) {
    defer wg.Done()

    cmd := exec.Command("bash", "-c", command)
    stdout, err := cmd.StdoutPipe()
    Check(ErrMessage{command,err,"StdoutPipe()"})
    /*    if err != nil {
        fmt.Printf("Error creating stdout pip for cmd %s: %v\n", command, err)
        return
    }
    */
    
    err = cmd.Start()
    Check(ErrMessage{command,err,"Start()"})
    /*
    if err:= cmd.Start(); err !=nil{
        fmt.Printf("Error starting cmd %s: %v\n", command, err)
        return
    }
    */
    scanner := bufio.NewScanner(stdout)
    go func() {
        for scanner.Scan(){
            time.Sleep(time.Duration(delay)*time.Second)
            outputChan <- scanner.Text()
        }
    }()
    
    select {
    case <-stopChan:
        if err:= cmd.Process.Kill(); err != nil{
            fmt.Printf("Failed to kill process: %s\n", err)
        }
    }
}

func Sync(command string) []byte {
    cmd := exec.Command("bash", "-c", command)
    output, err := cmd.Output()
    Check(ErrMessage{command,err,"Output()"})
    return output
}

func SimpleDelay(command string, delay int) {
    cmd := exec.Command("bash", "-c", command)

    // Start the command but do not wait for it to complete
    err := cmd.Start()
    Check(ErrMessage{command, err, "Start()"})
    
 
    timer := time.NewTimer(time.Duration(delay) * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		// If the delay is reached, kill the command
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Printf("Error killing command: %v\n", err)
		}
		return
	}
}

