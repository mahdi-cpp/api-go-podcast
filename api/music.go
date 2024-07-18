package api

import (
	"github.com/gin-gonic/gin"
)

func AddMusicRoutes(rg *gin.RouterGroup) {

	music := rg.Group("/music")

	music.GET("/feed", func(context *gin.Context) {

		//stories := podcast.GetViewSliderDTO(
		//	0,
		//	5,
		//	1,
		//	15, 5,
		//	[]string{"Stories", "داستان ها"},
		//)
		//albums := podcast.GetViewSliderDTO(
		//	1,
		//	6,
		//	3,
		//	15,
		//	0,
		//	[]string{"Albums", "آلبوم ها"},
		//)
		//recentlyTracks := podcast.GetViewSliderDTO(
		//	2,
		//	8,
		//	4,
		//	15,
		//	0,
		//	[]string{"Recently Played", "اخیرا گوش داده اید"},
		//)
		//playlist := podcast.GetViewSliderDTO(
		//	3,
		//	3,
		//	4,
		//	15,
		//	0,
		//	[]string{"Play Lists", "لیست های پخش"},
		//)
		//tracks := podcast.GetViewSliderDTO(
		//	4,
		//	8,
		//	6,
		//	15, 8,
		//	[]string{"Tracks", "تک آهنگ ها"},
		//)
		//
		//context.JSON(210, gin.H{
		//	"storyDTO":         stories,
		//	"albumDTO":         albums,
		//	"recentlyTrackDTO": recentlyTracks,
		//	"playlistDTO":      playlist,
		//	"trackDTO":         tracks,
		//})
	})
}
