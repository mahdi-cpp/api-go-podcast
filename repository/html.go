package repository

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/mahdi-cpp/api-go-podcast/utils"
	"log"
	"strconv"
	"strings"
)

type Shape struct {
	Dx     float64 `json:"dx"`
	Dy     float64 `json:"dy"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Round  float64 `json:"round"`
}

type Image struct {
	Shape Shape `json:"shape"`

	Url string `json:"url"`

	TextViews []TextView `json:"textViews"`
}

type TextView struct {
	Shape Shape `json:"shape"`

	Color    int    `json:"color"`
	Text     string `json:"text"`
	TextSize int    `json:"textSize"`

	Align  string `json:"align"`
	Anchor string `json:"anchor"`
}

//type Button struct {
//	Shape Shape `json:"shape"`
//
//	Text     string `json:"text"`
//	Icon     string `json:"icon"`
//	TextSize int    `json:"textSize"`
//	Color    int    `json:"color"`
//
//	Anchor string `json:"anchor"`
//}

//type Rectangle struct {
//	Color  int     `json:"color"`
//	Dx     float64 `json:"dx"`
//	Dy     float64 `json:"dy"`
//	Width  float64 `json:"width"`
//	Height float64 `json:"height"`
//	Round  float64 `json:"round"`
//}

type ViewSliderDTO struct {
	Images []Image `json:"images"`
	//Rectangles []Rectangle `json:"rectangles"`
	//TextViews  []TextView  `json:"textViews"`

	Margin         float64 `json:"margin"`
	PagePhotoCount float64 `json:"pagePhotoCount"`
	MaxPageCount   float64 `json:"maxPageCount"`
}

var viewSliderDTO ViewSliderDTO

func StartProcessHtmlFile() {
	ProcessHtmlFile("android-music.html")
}

func GetViewSliderDTO() ViewSliderDTO {
	return viewSliderDTO
}

func ProcessHtmlFile(html string) {

	filePath := html // Path to the HTML file

	// Read the HTML content from the file
	htmlContent, err := utils.ReadHTMLFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	FindViewScroll(doc)
}

func FindViewScroll(doc *goquery.Document) {

	ViewSlider := doc.Find("#ViewSlider")

	if ViewSlider.Size() == 0 {
		fmt.Println("ViewSlider not found")
		return
	}

	//ViewSlider.Each(func(i int, cell *goquery.Selection) {
	//	src, exists := cell.Attr("margin")
	//	if exists {
	//		fmt.Println("margin", src)
	//
	//	}
	//
	//	src, exists = cell.Attr("page_photo_count")
	//	if exists {
	//		fmt.Println("page_photo_count", src)
	//	}
	//
	//	src, exists = cell.Attr("max_page_count")
	//	if exists {
	//		fmt.Println("max_page_count", src)
	//	}
	//})

	//doc.Find("#image45").Children().Each(func(i int, s *goquery.Selection) {
	//	fmt.Print("Child %d: %s", i+1, s.Text())
	//})
	//
	//doc.Find("#image33").Children().Each(func(i int, s *goquery.Selection) {
	//	fmt.Print("Child %d: %s\n", i+1, s.Text())
	//})

	ViewSlider.Find("ImageView").Each(func(i int, cell *goquery.Selection) {

		var image Image

		src, exists := cell.Attr("id")
		if exists {

			doc.Find("#" + src).Children().Each(func(i int, textViewChild *goquery.Selection) {

				//fmt.Print(src + " child: " + s.Text())
				var textView TextView

				src, exists := textViewChild.Attr("text")
				if exists {
					textView.Text = src
				}

				src, exists = textViewChild.Attr("text_size")
				if exists {
					textView.TextSize, _ = strconv.Atoi(src)
				}

				src, exists = textViewChild.Attr("color")
				if exists {
					color, _ := utils.Parse(src)
					var rgb = int(color.ToRGB().R)
					rgb = (rgb << 8) + int(color.ToRGB().G)
					rgb = (rgb << 8) + int(color.ToRGB().B)
					textView.Color = rgb
				}

				src, exists = textViewChild.Attr("width")
				if exists {
					textView.Shape.Width, _ = strconv.ParseFloat(src, 64)
				}

				src, exists = textViewChild.Attr("height")
				if exists {
					textView.Shape.Height, _ = strconv.ParseFloat(src, 64)
				}
				src, exists = textViewChild.Attr("dx")
				if exists {
					textView.Shape.Dx, _ = strconv.ParseFloat(src, 64)
				}
				src, exists = textViewChild.Attr("dy")
				if exists {
					textView.Shape.Dy, _ = strconv.ParseFloat(src, 64)
				}
				src, exists = textViewChild.Attr("align")
				if exists {
					textView.Align = src
				}
				src, exists = textViewChild.Attr("anchor")
				if exists {
					textView.Anchor = src
				}
				fmt.Println(textView)
				image.TextViews = append(image.TextViews, textView)
			})
		}

		src, exists = cell.Attr("url")
		if exists {
			image.Url = src
		}
		src, exists = cell.Attr("width")
		if exists {
			image.Shape.Width, _ = strconv.ParseFloat(src, 64)
		}
		src, exists = cell.Attr("height")
		if exists {
			image.Shape.Height, _ = strconv.ParseFloat(src, 64)
		}
		src, exists = cell.Attr("dx")
		if exists {
			image.Shape.Dx, _ = strconv.ParseFloat(src, 64)
		}
		src, exists = cell.Attr("dy")
		if exists {
			image.Shape.Dy, _ = strconv.ParseFloat(src, 64)
		}
		src, exists = cell.Attr("round")
		if exists {
			image.Shape.Round, _ = strconv.ParseFloat(src, 64)
		}
		//fmt.Println(image)
		viewSliderDTO.Images = append(viewSliderDTO.Images, image)
	})

	//RectangleFind := ViewSlider.Find("Rectangle")
	//RectangleFind.Each(func(i int, rectangle *goquery.Selection) {
	//
	//	var rect Rectangle
	//
	//	src, exists := rectangle.Attr("color")
	//	if exists {
	//		color, _ := utils.Parse(src)
	//		var rgb = int(color.ToRGB().R)
	//		rgb = (rgb << 8) + int(color.ToRGB().G)
	//		rgb = (rgb << 8) + int(color.ToRGB().B)
	//		rect.Color = rgb
	//
	//		//fmt.Println(rgb)
	//	}
	//	src, exists = rectangle.Attr("width")
	//	if exists {
	//		rect.Width, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = rectangle.Attr("height")
	//	if exists {
	//		rect.Height, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = rectangle.Attr("dx")
	//	if exists {
	//		rect.Dx, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = rectangle.Attr("dy")
	//	if exists {
	//		rect.Dy, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = rectangle.Attr("round")
	//	if exists {
	//		rect.Round, _ = strconv.ParseFloat(src, 64)
	//	}
	//	//fmt.Println(rect)
	//	viewSliderDTO.Rectangles = append(viewSliderDTO.Rectangles, rect)
	//})

	//TextViewFind := ViewSlider.Find("TextView")
	//TextViewFind.Each(func(i int, cell *goquery.Selection) {
	//
	//	var textView TextView
	//
	//	src, exists := cell.Attr("text")
	//	if exists {
	//		textView.Text = src
	//	}
	//
	//	src, exists = cell.Attr("text_size")
	//	if exists {
	//		textView.TextSize, _ = strconv.Atoi(src)
	//	}
	//
	//	src, exists = cell.Attr("color")
	//	if exists {
	//		color, _ := utils.Parse(src)
	//		var rgb = int(color.ToRGB().R)
	//		rgb = (rgb << 8) + int(color.ToRGB().G)
	//		rgb = (rgb << 8) + int(color.ToRGB().B)
	//		textView.Color = rgb
	//	}
	//
	//	src, exists = cell.Attr("width")
	//	if exists {
	//		textView.Width, _ = strconv.ParseFloat(src, 64)
	//	}
	//
	//	src, exists = cell.Attr("height")
	//	if exists {
	//		textView.Height, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = cell.Attr("dx")
	//	if exists {
	//		textView.Dx, _ = strconv.ParseFloat(src, 64)
	//	}
	//	src, exists = cell.Attr("dy")
	//	if exists {
	//		textView.Dy, _ = strconv.ParseFloat(src, 64)
	//	}
	//	//fmt.Println(textView)
	//	viewSliderDTO.TextViews = append(viewSliderDTO.TextViews, textView)
	//})

	//sort.Slice(rectangles, func(i, j int) bool {
	//	return rectangles[i].Z < rectangles[j].Z
	//})
	//sort.Slice(rectangles, func(i, j int) bool {
	//	return rectangles[i].Z < rectangles[j].Z
	//})

	fmt.Println("-------------------------------")
}
