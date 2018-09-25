package main

import "fmt"
import "os"
import "plugin"

func main() {
	for {
		var lang string
		fmt.Print("Insert which language do you prefer: ")
		fmt.Scanf("%s", &lang)

		symLanguage := lookupPlugin("./languages/"+lang+"/"+lang+".so", "Speak")

		speak, ok := symLanguage.(func() string)
		if !ok {
			fmt.Println("The function signature is different")
			os.Exit(1)
		}

		fmt.Printf("%s\n", speak())
	}
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
