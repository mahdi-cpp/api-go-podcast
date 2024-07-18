package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-podcast/repository"
)

func AddPodcastRoutes(rg *gin.RouterGroup) {

	podcast := rg.Group("/podcast")

	podcast.GET("/feed", func(context *gin.Context) {

		podcastDTO := repository.GetPodcastDTO(
			6,
			3,
			15,
			0,
			[]string{"Albums", "آلبوم ها"},
		)

		context.JSON(210, gin.H{
			"podcastDTO": podcastDTO,
		})
	})

	podcast.GET("/shapes", func(context *gin.Context) {

		repository.StartProcessHtmlFile()
		context.JSON(210, repository.GetViewSliderDTO())

		//rectangles := repository.GetRectangle()
		//images := repository.GetImageView()
		//textViews := repository.GetTextView()
		//
		//context.JSON(210, gin.H{
		//	"rectangles": rectangles,
		//	"images":     images,
		//	"textViews":  textViews,
		//})

	})

}
