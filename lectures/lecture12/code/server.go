package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

const tpl = `<html>
<head></head>
<body>
    <script type="text/javascript">
        var sock = null;
        var wsuri = "ws://127.0.0.1:1234";

        window.onload = function() {

            console.log("onload");

            sock = new WebSocket(wsuri);

            sock.onopen = function() {
                console.log("connected to " + wsuri);
            }

            sock.onclose = function(e) {
                console.log("connection closed (" + e.code + ")");
            }

            sock.onmessage = function(e) {
                console.log("message received: " + e.data);
            }
        };

        function send() {
            var msg = document.getElementById('message').value;
            sock.send(msg);
        };
    </script>
    <h1>WebSocket Echo Test</h1>
    <form>
        <p>
            Message: <input id="message" type="text" value="Hello, world!">
        </p>
    </form>
    <button onclick="send();">Send Message</button>
</body>
</html>`

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("tpl")
	content, _ := t.Parse(tpl)
	content.Execute(w, nil)
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.Handle("/send", http.HandlerFunc(myHandler))

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
