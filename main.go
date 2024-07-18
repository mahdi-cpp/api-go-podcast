package main

func main() {
	//javascript.StarScriptParse()
	//tutorial.ShowInterface()
	//tutorial.Methods()
	//tutorial.MainPointInPolygon()
	Run()
}

//
//// Selection represents a collection of integers
//type Selection struct {
//	elements []int
//	Name     string
//	title    string
//	Width    float32
//	Height   float32
//}
//
//// Each applies a function to each element in the selection
//func (s *Selection) Each(f func(int, *Selection)) *Selection {
//	for _, element := range s.elements {
//		f(element, s)
//	}
//	return s
//}
//
//func main() {
//	// Create a Selection instance with some elements
//	selection := &Selection{elements: []int{1, 2, 3, 4, 5, 6, 7}, Name: "Image"}
//
//	////Define a function to be applied to each element
//	//applyFunc := func(num int, s *Selection) {
//	//	fmt.Printf("Element: %d, Selection: %v\n", num, s.elements)
//	//}
//
//	// Call the Each method on the Selection instance
//	selection.Each(func(i int, selection *Selection) {
//		fmt.Printf("Element: %d, Selection: %v\n", i, selection.Name)
//	})
//}
