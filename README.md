Tail (or pipe) anything to a web browser or web browsers.

## ESTail (Event Source Tail)

Useage

    tail -f somefile.log | ./estail

Open a browser to localhost:8080 or maybe MachineName.local:8080

Need to clear the buffer? Refresh the browser.

### Get and Build

Requires [Go](http://golang.org/doc/install) to be installed

    git clone git@github.com:jweir/estail.git
    cd estail
    go get && go build

You should now have a binary file `estail`. Install in your PATH and you will be good to go.

### Why

I wanted to stream some log files from one machine to an ipad.

### Options

Want options? Fork it!

### License 

Copyright (c) 2013 John Weir

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
