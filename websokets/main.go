package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	// Лупчик
	for i:=0; i<10;i++ {
		x:=rand.Intn(100)
		y:=rand.Intn(100)
		msg:=fmt.Sprintf(`{"x":%d,"y":%d}`, x,y)
		err = c.WriteMessage(1, []byte(msg))
		if err != nil {
			log.Println("write:", err)
			break
		}
		time.Sleep(500*time.Millisecond)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
	// homeTemplate.Execute(w, "ws://127.0.0.1:1234")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<body>
<canvas id="myCanvas" width="500" height="500" style="border:1px solid #d3d3d3;">
Your browser does not support the HTML5 canvas tag.</canvas>

<script>
var canvas = document.getElementById('myCanvas');
var ctx = canvas.getContext('2d');

function sendCoordinates(x, y) {
    <!-- ctx.clearRect(0, 0, canvas.width, canvas.height); // Clear the canvas -->
    ctx.beginPath();
    ctx.arc(x, y, 1, 0, 2 * Math.PI); // Draw a circle at the new coordinates
    ctx.stroke();
}

var ws = new WebSocket('ws://localhost:8080/echo');

ws.onmessage = function(event) {
    var data = JSON.parse(event.data);
    sendCoordinates(data.x, data.y);
};

ws.onerror = function(error) {
    console.log('WebSocket Error: ' + error);
};

ws.onclose = function(event) {
    console.log('WebSocket closed: code = ' + event.code + ', reason = ' + event.reason);
};
</script>

</body>
</html>


    `))
