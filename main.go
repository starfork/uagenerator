package main

import (
	"flag"
	"os"

	"github.com/starfork/uagenerator/generator"
)

var (
	cFile = flag.String("c", "", "config file,empty for default")
	//generator type
	t = flag.String("t", "", "type to generate,m for mobile default for pc")
	n = flag.Int("n", 50, "number to generator")
)

func main() {
	flag.Parse()
	if *cFile != "" {
		generator.LoadConfig(*cFile)
	}

	tf, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		os.Exit(1)
	}
	for i := 0; i < *n; i++ {
		if *t == "m" {
			tf.WriteString(generator.RandomMobileUserAgent() + "\r\n")

		} else {
			tf.WriteString(generator.RandomUserAgent() + "\r\n")
		}
	}

	tf.Close()

}
