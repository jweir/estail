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
  "flag"
	"os"
)

var Stdin *bufio.Reader
var printOut bool
var port string

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
      if printOut == true{
        fmt.Println(string(line))
      }
		}
	}()

	log.Printf("ESTail starting at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	Stdin = bufio.NewReader(os.Stdin)
  flag.BoolVar(&printOut, "v", false, "print the STDOUT (quiet by default)")
  flag.StringVar(&port, "p", "8080", "http port")
  flag.Parse()
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
        if(addMark){ target.append($("<hr/>")); addMark = false}
      },3000)

      source.onmessage = function(e){
        addMark = true;
        var d = $("<div/>")
        target.append(d.text(parse(e.data)))
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
