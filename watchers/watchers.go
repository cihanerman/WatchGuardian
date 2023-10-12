package watchers

import (
	"bufio"
	"github.com/cihanerman/WatchGuardian/utils"
	"github.com/fsnotify/fsnotify"
	"log"
)

func CreateWatcher() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	return watcher
}

func AddFileToWatcher(watcher *fsnotify.Watcher, filePath string) {
	err := watcher.Add(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func StartWatcher(watcher *fsnotify.Watcher, closeWatcher chan struct{}, fileReader *bufio.Reader, postUrl string, headerVariable string, token string) {
	defer close(closeWatcher)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Has(fsnotify.Write) {
				line, err := fileReader.ReadString('\n')
				utils.CheckError(err)
				line = utils.TrimInput(line)
				go utils.SendUpdate(line, event.Name, event.Op.String(), postUrl, headerVariable, token)
			} else if event.Has(fsnotify.Remove) {
				return
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("error:", err)
		}
	}
}
