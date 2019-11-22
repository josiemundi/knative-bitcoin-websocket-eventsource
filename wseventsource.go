package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	"github.com/kelseyhightower/envconfig"
	"github.com/rgamba/evtwebsocket"
	"knative.dev/eventing-contrib/pkg/kncloudevents"
)

var (
	eventSource string
	eventType   string
	sink        string
	label       string
	periodStr   string
)

func init() {
	flag.StringVar(&eventSource, "eventSource", "", "the event-source (CloudEvents)")
	flag.StringVar(&eventType, "eventType", "dev.knative.samples.wsevent", "the event-type (CloudEvents)")
	flag.StringVar(&sink, "sink", "http://default-broker.default.svc.cluster.local", "the host url to heartbeat to")
	flag.StringVar(&label, "label", "", "a special label")
	flag.StringVar(&periodStr, "period", "5", "the number of seconds between heartbeats")
}

type envConfig struct {
	// Sink URL where to send heartbeat cloudevents
	Sink string `envconfig:"SINK"`

	// // Name of this pod.
	// Name string `envconfig:"POD_NAME" required:"true"`

	// // Namespace this pod exists in.
	// Namespace string `envconfig:"POD_NAMESPACE" required:"true"`
}

func main() {

	flag.Parse()

	var env envConfig
	if err := envconfig.Process("", &env); err != nil {
		log.Printf("[ERROR] Failed to process env var: %s", err)
		os.Exit(1)
	}

	if env.Sink != "" {
		sink = env.Sink
	}

	fmt.Println("this is my sink:", sink)
	c, err := kncloudevents.NewDefaultClient(sink)
	if err != nil {
		log.Fatalf("failed to create client: %s", err.Error())
	}

	if eventSource == "" {
		eventSource = fmt.Sprintf("https://knative.dev/knative-eventing-websocket-source/") //, env.Namespace, env.Name)
		log.Printf("Heartbeats Source: %s", eventSource)
	}

	websocketclient := evtwebsocket.Conn{

		// When connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			log.Println("Connected")
		},

		// When a message arrives
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			//log.Printf("OnMessage: %s\n", msg)
			//fmt.Println(msg)
			fmt.Printf("MESSAGE: %s\n", msg)
			// var transact Transaction
			// //data := []byte(msg)
			// err := json.Unmarshal(msg, &transact)
			// if err == nil {
			// 	//fmt.Printf("%s", msg)
			// 	fmt.Printf("INPUTS:\n")
			// 	for k, v := range transact.X.Inputs {
			// 		//fmt.Println(parsed["inputs"])
			// 		fmt.Println(k, v)
			// 	}
			// 	fmt.Println("")
			// }

			event := cloudevents.Event{
				Context: cloudevents.EventContextV03{
					Type:   eventType,
					Source: *types.ParseURLRef(eventSource),
				}.AsV02(),
				Data: msg,
			}

			//should this be c.StartReceiver?
			if _, _, err := c.Send(context.Background(), event); err != nil {
				log.Printf("failed to send cloudevent: %s", err.Error())
			}
		},

		// When the client disconnects for any reason
		OnError: func(err error) {
			log.Printf("** ERROR **\n%s\n", err.Error())
		},

		// This is used to match the request and response messagesP>termina
		MatchMsg: func(req, resp []byte) bool {
			return string(req) == string(resp)
		},

		// Auto reconnect on error
		Reconnect: true,

		// Set the ping interval (optional)
		PingIntervalSecs: 5,

		// Set the ping message (optional)
		PingMsg: []byte("PING"),
	}

	// Connect
	wsErr := websocketclient.Dial("wss://ws.blockchain.info/inv", "")
	if wsErr != nil {
		log.Fatal(wsErr)
	}

	outmsg := evtwebsocket.Msg{
		Body: []byte("{\"op\":\"unconfirmed_sub\"}"),
		Callback: func(resp []byte, w *evtwebsocket.Conn) {
			log.Printf("[%d] Callback: %s\n", 0, resp)
		},
	}

	//err = c.Send(msg)
	if err := websocketclient.Send(outmsg); err != nil {
		log.Println("Unable to send: ", err.Error())
	}

	time.Sleep(time.Second * 2 * 60)
}
