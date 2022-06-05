package main

import (
	"hacktiv8-final-project/config"
	"hacktiv8-final-project/controllers"
	"hacktiv8-final-project/middlewares"
	"hacktiv8-final-project/repositories"
	"hacktiv8-final-project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	route := gin.Default()
	//User Repo
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewItemService(userRepo)
	userController := controllers.NewUserController(userService)

	//Route Group
	userRouter := route.Group("/users")
	{

		userRouter.POST("/register", userController.UserRegister)
		userRouter.POST("/login", userController.Login)
		userRouter.Use(middlewares.Auth())
		userRouter.PUT("/:userId", userController.UpdateUser)
		userRouter.DELETE("/:userId", userController.DeleteUser)
	}

	//Photo Repo
	photoRepo := repositories.NewPhotoRepo(db)
	photoService := services.NewPhotoService(photoRepo)
	photoController := controllers.NewPhotoController(photoService)

	photoRouter := route.Group("/photos")
	{
		photoRouter.Use(middlewares.Auth())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetAllPhotos)
		photoRouter.PUT("/:photoId", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoId", photoController.DeletePhoto)
	}

	//CommentRepo
	commentRepo := repositories.NewCommentRepo(db)
	commentService := services.NewCommentService(commentRepo)
	commentContoller := controllers.NewCommentContoller(commentService)

	commentRouter := route.Group("/comments")
	{
		commentRouter.Use(middlewares.Auth())
		commentRouter.POST("/", commentContoller.CreateComment)
		commentRouter.GET("/", commentContoller.GetAllComments)
		commentRouter.PUT("/:commentId", commentContoller.UpdateComment)
		commentRouter.DELETE("/:commentId", commentContoller.DeleteComment)

	}

	//SocialMediaRepo
	socialMediaRepo := repositories.NewSocialMediaRepo(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	socialMediaRouter := route.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Auth())
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", socialMediaController.GetAllSocialMedias)
		socialMediaRouter.PUT("/:socialMediaId", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", socialMediaController.DeleteSocialMedia)
	}

	route.Run(config.APP_PORT)
}
