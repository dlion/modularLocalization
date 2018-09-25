package main

import "fmt"
import "os"
import "plugin"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: ./%s <language plugin>", os.Args[0])
	}

	lang := os.Args[1]

	symLanguage := lookupPlugin("./languages/"+lang+"/"+lang+".so", "Speak")

	speak, ok := symLanguage.(func() string)
	if !ok {
		fmt.Println("The function signature is different")
		os.Exit(1)
	}

	fmt.Printf("%s\n", speak())

}

func lookupPlugin(p, s string) plugin.Symbol {
	plug, err := plugin.Open(p)
	errorHandler(err)

	sym, err := plug.Lookup(s)
	errorHandler(err)
	return sym
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
