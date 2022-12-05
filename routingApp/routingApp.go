package routingApp

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const rootDir = "routingAppOld"
const dirName = "sendingTmp"
const backupDirName = "filesSent"

var filesToSend []string
var filesAlreadySent []string

func scanForNewFiles(dirName string, wg *sync.WaitGroup, m *sync.Mutex) {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Println("Error reading files")
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			wg.Add(1)
			scanForNewFiles(dirName+string(os.PathSeparator)+name, wg, m)
		} else {
			var canAdd = true
			for i := range filesToSend {
				if filesToSend[i] == dirName+string(os.PathSeparator)+name {
					canAdd = false
				}
			}
			if canAdd {
				m.Lock()
				filesToSend = append(filesToSend, dirName+string(os.PathSeparator)+name)
				m.Unlock()
			}
		}
	}
	wg.Done()
}

func sendFiles(wg *sync.WaitGroup, m *sync.Mutex) {
	for i := range filesToSend {
		path := filesToSend[i]
		var newPath = rootDir + string(os.PathSeparator) + backupDirName + path[len(rootDir)+1+len(dirName):]
		err := os.Rename(path, newPath)
		if err != nil {
			fmt.Println(err)
		}
		m.Lock()
		filesAlreadySent = append(filesAlreadySent, path)
		m.Unlock()
	}
	filesToSend = []string{}
	wg.Done()
}

func main() {
	var m sync.Mutex
	var w = sync.WaitGroup{}
	for {
		w.Add(1)

		go scanForNewFiles(rootDir+string(os.PathSeparator)+dirName, &w, &m)

		w.Add(1)
		go sendFiles(&w, &m)

		w.Wait()
		time.Sleep(time.Second * 10)
	}
}
