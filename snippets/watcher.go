package snippets

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func Watch() {
	w := watcher.New()

	// If SetMaxEvents is not set, the default is to send all events.
	// w.SetMaxEvents(1)

	// Uncomment to only notify rename and move events.
	// w.FilterOps(watcher.Rename, watcher.Move)

	r := regexp.MustCompile(".go")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Printf("\nSnippets: Event %v info\n", event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add("./snippets/gowatcher/"); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	fmt.Println("Files being watched!!!")
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	// Trigger 2 events after watcher started.
	//go func() {
	//	w.Wait()
	//	w.TriggerEvent(watcher.Create, nil)
	//	w.TriggerEvent(watcher.Remove, nil)
	//}()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
