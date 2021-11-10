package tools

import (
	"flag"
	"fmt"
	"os"
)


func ReadInput(){
	textPtr := flag.String("text", "", "Text to parse. (Required)")
    metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};. (Required)")
    uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
    flag.Parse()

    if *textPtr == "" {
        flag.PrintDefaults()
        os.Exit(1)
		
    }
	SaveToFile(*textPtr)
    fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
}