// ESTail "broadcasts" anything piped to over HTTP using EventSource
//
// Example
//    tail -f file.log | go run estail.go
//
// Then open http://localhost:8080
//
package main

import (
	"bufio"
	"fmt"
	eventsource "github.com/antage/eventsource/http"
	"log"
	"net/http"
	"os"
)

var Stdin *bufio.Reader

func main() {
	es := eventsource.New(nil)
	defer es.Close()
	http.Handle("/log/", es)
	http.HandleFunc("/", index)

	go func() {
		for {
			line, _, err := Stdin.ReadLine()
			if err != nil {
				fmt.Println(err)
				break
			}
			es.SendMessage(string(line), "", "")
		}
	}()

	log.Println("ESTail starting at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	Stdin = bufio.NewReader(os.Stdin)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, template())
}

func template() string {
	return `
  <!DOCTYPE html>

  <html>
  <head>
  <title>ESTail</title>
  <link href='http://fonts.googleapis.com/css?family=Droid+Sans+Mono' rel='stylesheet' type='text/css'>
  <script src="//cdnjs.cloudflare.com/ajax/libs/jquery/1.9.0/jquery.min.js"></script>
  <style type="text/css">
  body {
    margin: 20px;
    font: 12px/18px 'Droid Sans Mono', san-serif;
    color: #CCCCAA;
    background: #222;
  }

  hr {
    border:0;
    height: 1px;
    line-height: 1px;
    border-top: 1px #590 solid;
    margin: 9px 0 10px;
  }

  </style>
  </head>

  <body>
  <h4>ESTail</h4>
  <div id="log"></div>

  <script>
  (function(){

    // clean up some shell color codes
    function parse(s){
      return s.replace(/\[\d{1,}m/img,"")
    }

    function listen(channel, target){
      var source = new EventSource('/log/');
      var addMark = false;

      setInterval(function(){
        console.log("hi", addMark)
        if(addMark){ target.append($("<hr/>")); addMark = false}
      },3000)

      source.onmessage = function(e){
        addMark = true;
        target.append($("<div>"+parse(e.data)+"</div>"))
        $(document).scrollTop($(document).height())
      }
    }

    listen("log", $("#log"))
  })()
  </script>
  </body>

  </html>
  `
}
