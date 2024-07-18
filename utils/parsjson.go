package utils

//func main() {
//	// Read the JavaScript file
//	jsFile, err := ioutil.ReadFile("data.js")
//	if err != nil {
//		fmt.Println("Error reading JavaScript file:", err)
//		return
//	}
//
//	// Create a new JavaScript VM
//	vm := otto.New()
//
//	// Execute the JavaScript code
//	_, err = vm.Run(string(jsFile))
//	if err != nil {
//		fmt.Println("Error executing JavaScript code:", err)
//		return
//	}
//
//	// Get the JSON data from JavaScript
//	value, err := vm.Get("Button")
//	if err != nil {
//		fmt.Println("Error getting JSON data:", err)
//		return
//	}
//
//	fmt.Println("is function ", value.IsFunction())
//
//	// Convert JavaScript object to JSON
//	jsonObj, err := value.Export()
//	if err != nil {
//		fmt.Println("Error exporting JavaScript object:", err)
//		return
//	}
//
//	// Print the JSON data
//	fmt.Println("JSON data:", jsonObj)
//}
