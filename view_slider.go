package main

import (
	"github.com:mahdi-cpp/api-go-podcast/model"
	podcast "github.com:mahdi-cpp/api-go-podcast/repository"
	"github.com:mahdi-cpp/api-go-podcast/utils"
)

const (
	STORY    int = 1
	ALBUM    int = 2
	PLAYLIST int = 3
	TRACK    int = 4
	RECENTLY int = 5
)

type ViewSliderDTO struct {
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

	Albums []model.Album `json:"albums"`
}

func GetViewSliderDTO(View int, MaxPageCount int, PagePhotosCount float64, MarginX float64, Round float64, HeaderText []string) ViewSliderDTO {

	var PhotoWidth float64 = 0
	var PhotoHeight float64 = 0

	var albumDTO = ViewSliderDTO{}

	switch View {

	case STORY:
		MarginX = utils.DP(MarginX)
		PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount * MarginX) + MarginX)) / PagePhotosCount
		PhotoHeight = PhotoWidth

		albumDTO = ViewSliderDTO{
			View:            STORY,
			Albums:          podcast.Stories(),
			HeaderText:      HeaderText,
			MaxPageCount:    MaxPageCount,
			PagePhotosCount: int(PagePhotosCount),
			MarginX:         MarginX,
			Round:           Round,
			PhotoWidth:      PhotoWidth,
			PhotoHeight:     PhotoHeight,

			CellWidth:  utils.ScreenWidth,
			CellHeight: utils.DP(100),
		}
		break
	case ALBUM:
		MarginX = utils.DP(MarginX)
		PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount * MarginX) + MarginX)) / PagePhotosCount
		PhotoHeight = PhotoWidth

		albumDTO = ViewSliderDTO{
			View:            ALBUM,
			Albums:          podcast.Albums(),
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
		//fmt.Println(albumDTO)

		break
	case RECENTLY:
		MarginX = utils.DP(MarginX)
		PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount*2)*MarginX + MarginX)) / PagePhotosCount
		PhotoHeight = PhotoWidth

		albumDTO = ViewSliderDTO{
			View:            RECENTLY,
			Albums:          podcast.RecentlyTracks(),
			HeaderText:      HeaderText,
			MaxPageCount:    MaxPageCount,
			PagePhotosCount: int(PagePhotosCount),
			MarginX:         MarginX,
			Round:           Round,
			PhotoWidth:      PhotoWidth,
			PhotoHeight:     PhotoHeight,

			CellWidth:  utils.ScreenWidth,
			CellHeight: PhotoHeight + PhotoHeight + utils.DP(120),
		}
		break

	case PLAYLIST:
		MarginX = utils.DP(MarginX)
		PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount*4.0)*MarginX + MarginX)) / (PagePhotosCount / 2)
		PhotoHeight = PhotoWidth

		albumDTO = ViewSliderDTO{
			View:            PLAYLIST,
			Albums:          podcast.Playlists(),
			HeaderText:      HeaderText,
			MaxPageCount:    MaxPageCount,
			PagePhotosCount: int(PagePhotosCount),
			MarginX:         MarginX,
			Round:           Round,
			PhotoWidth:      PhotoWidth,
			PhotoHeight:     PhotoHeight,

			CellWidth:  utils.ScreenWidth,
			CellHeight: PhotoHeight + PhotoHeight + utils.DP(170),
		}
		break

	case TRACK:
		MarginX = utils.DP(MarginX)
		PhotoWidth = (utils.ScreenWidth - ((PagePhotosCount/2)*MarginX + MarginX)) / (PagePhotosCount / 2)
		PhotoHeight = PhotoWidth

		albumDTO = ViewSliderDTO{
			View:            TRACK,
			Albums:          podcast.Tracks(),
			HeaderText:      HeaderText,
			MaxPageCount:    MaxPageCount,
			PagePhotosCount: int(PagePhotosCount),
			MarginX:         MarginX,
			Round:           Round,
			PhotoWidth:      PhotoWidth,
			PhotoHeight:     PhotoHeight,

			CellWidth:  utils.ScreenWidth,
			CellHeight: PhotoHeight + PhotoHeight + utils.DP(170),
		}
		break
	}

	return albumDTO
}
