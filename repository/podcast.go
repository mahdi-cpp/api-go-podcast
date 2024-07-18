package repository

import (
	"github.com/mahdi-cpp/api-go-podcast/model"
	"github.com/mahdi-cpp/api-go-podcast/utils"
	"log"
	"os"
)

type PodcastDTO struct {
	View            int      `json:"view"`
	HeaderText      []string `json:"headerText"`
	PagePhotosCount int      `json:"pagePhotosCount"`
	MarginX         float64  `json:"marginX"`
	Round           float64  `json:"round"`
	MaxPageCount    int      `json:"maxPageCount"`

	PhotoWidth  float64 `json:"photoWidth"`
	PhotoHeight float64 `json:"photoHeight"`

	CellWidth  float64 `json:"cellWidth"`
	CellHeight float64 `json:"cellHeight"`

	Images []model.Image `json:"albums"`
}

func GetPodcastDTO(MaxPageCount int, PagePhotosCount float64, MarginX float64, Round float64, HeaderText []string) PodcastDTO {

	var PhotoWidth float64 = 0
	var PhotoHeight float64 = 0

	var albumDTO = PodcastDTO{}

	MarginX = utils.DP(MarginX)
	PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount * MarginX) + MarginX)) / PagePhotosCount
	PhotoHeight = PhotoWidth

	albumDTO = PodcastDTO{
		Images:          GetImages(),
		HeaderText:      HeaderText,
		MaxPageCount:    MaxPageCount,
		PagePhotosCount: int(PagePhotosCount),
		MarginX:         MarginX,
		Round:           Round,
		PhotoWidth:      PhotoWidth,
		PhotoHeight:     PhotoHeight,

		CellWidth:  utils.ScreenWidth,
		CellHeight: PhotoHeight + utils.DP(140),
	}

	return albumDTO
}

func GetImages() []model.Image {
	var images []model.Image
	var root = "/var/instagram/"
	var path = "music/ali3/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		images = append(images, model.Image{Id: 1, AlbumName: e.Name(), AlbumPath: path + e.Name()})
	}
	return images
}
