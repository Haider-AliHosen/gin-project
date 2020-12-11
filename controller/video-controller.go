package controller

// import (
// 	"github.com/Haider-AliHosen/gin-project/entity"
// 	"github.com/Haider-AliHosen/gin-project/service"
// 	"github.com/gin-gonic/gin"
// )

// //VideoController for controll ineterfacess
// type VideoController interface {
// 	FindAll() []entity.Video
// 	Save(ctx *gin.Context) entity.Video
// }
// type controller struct {
// 	service service.VideoService
// }

// //New for use servicess
// func New(service service.VideoService) VideoController {
// 	return controller{
// 		service: service,
// 	}
// }
// func (c *controller) FindAll() []entity.Video {
// 	return c.service.FindAll()
// }
// func (c *controller) Save(ctx *gin.Context) entity.Video {
// 	var video entity.Video
// 	ctx.BindJSON(&video)
// 	c.service.Save(video)
// 	return video
// }
