package cook

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strings"
)

type DeviceEntry struct{
    MAC string
    BSSID string
}

func CleanCSV(inputFilename, outputFilename, keyword string) error {
	inputFile, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	foundKeyword := false

	for scanner.Scan() {
		line := scanner.Text()

		if foundKeyword || strings.HasPrefix(line, keyword) {
			// Write the line to the output file
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				return err
			}
			foundKeyword = true
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("Lines including and after '%s' copied successfully to %s.\n", keyword, outputFilename)
	return nil
}

func ReadCSV(inputFilename string) ([]DeviceEntry, error){
    var devices []DeviceEntry

    file, err := os.Open(inputFilename)
    if err != nil {
        return devices, err
    }
    defer file.Close()
    fmt.Println("csv opened successfully")

    reader := csv.NewReader(file)

	// Read the header line
	header, err := reader.Read()
	if err != nil {
        fmt.Println("error:", err)
		return devices, err
	}

	// Find the indices of columns based on partial matches
	macIndex := -1
	bssidIndex := -1
	for i, column := range header {
		if strings.Contains(column, "Station MAC") {
			macIndex = i
		}
		if strings.Contains(column, "BSSID") {
			bssidIndex = i
		}
	}

	if macIndex == -1 || bssidIndex == -1 {
		return devices, fmt.Errorf("Columns 'Station MAC' and 'BSSID' not found in CSV header")
	}

	// Read lines until the end of the file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return devices, err
		}

		// Extract MAC and BSSID from the current record
		mac := record[macIndex]
		bssid := record[bssidIndex]

		// Create a DeviceEntry and append it to the devices slice
		device := DeviceEntry{MAC: mac, BSSID: bssid}
		devices = append(devices, device)
	}

	return devices, nil
}

func Strip(s []DeviceEntry) []DeviceEntry {
    var result []DeviceEntry
    for _, entry := range s{
        if entry.BSSID != " (not associated) " {
            result = append(result, entry)
        }
    }
    return result
}



//************************************************************************************\\
/*type NetworkEntry struct{
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
*/



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


