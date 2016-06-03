package update

import (
	"log"
	"fmt"
	//"log"
	"testing"

	"github.com/zabawaba99/firego"
)

type Message struct {
	data string
}

func TestConnect(t *testing.T) {
	// authToken := "3T2ARevFXMY9y0bWnfGf1cBm3ZZfpPdkxcRtxlth"
	// url := "https://firego-demo.firebaseio.com/messages/data/greeting"
	//fireBase := firebase.NewReference(url).Auth(authToken)

	f := firego.New("https://firego-new.firebaseio.com/", nil)
	f.Auth("AkKUqr3I3Nx26nxFcyttIighPdrkdbxEmHydtg10")

	// v := "bar"
	// _, err := f.Push(v)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	notifications := make(chan firego.Event)
	if err := f.Watch(notifications); err != nil {
		t.Fatal(err)
	}

	defer f.StopWatching()
	// for event := range notifications {
	// 	fmt.Printf("Event %#v\n", event)
	//     fmt.Println("Event Received")
	// 	fmt.Printf("Type: %s\n", event.Type)
	// 	fmt.Printf("Path: %s\n", event.Path)
	// 	fmt.Printf("Data: %v\n", event.Data)
	// }

	//outType := getType(notifications)
    outData := getData(getPath(getType(notifications)))
    
    //outPath := getPath(notifications)
	//fmt.Println(<-outType)
    for event := range outData {
		fmt.Printf("Event %#v\n", event)
	    fmt.Println("Event Received")
		fmt.Printf("Type: %s\n", event.Type)
		fmt.Printf("Path: %s\n", event.Path)
		fmt.Printf("Data: %v\n", event.Data)
	}
    //fmt.Println(<-outData)
	fmt.Printf("Notifications have stopped")	

}

func getType(evts <-chan firego.Event) <-chan firego.Event {
	out := make(chan firego.Event)
	go func() {
		for n := range evts {
            //log.Println
            log.Println("type is :: ",n.Type)
			out <- n
		}
		close(out)
	}()
	return out
}
func getPath(evts <-chan firego.Event) <-chan firego.Event {
	out := make(chan firego.Event)
	go func() {
		for n := range evts {
            log.Println("path is :: ",n.Path)
			out <- n
		}
		close(out)
	}()
	return out
}
func getData(evts <-chan firego.Event) <-chan firego.Event {
	out := make(chan firego.Event)
	go func() {
		for n := range evts {
            log.Println("data is :: ",n.Data)
			out <- n
		}
		close(out)
	}()
	return out
}
