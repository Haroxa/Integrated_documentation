package router

import (
	"github.com/Haroxa/Integrated_documentation/controller"
	"github.com/Haroxa/Integrated_documentation/middleware"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	e.POST("/user/login", controller.Login)
	e.POST("/user/register", controller.Register)
	e.POST("/user/register/reg", controller.Reg)

	user := e.Group("user")
	user.Use(middleware.AuthMiddleware)
	{
		user.GET("/getall", controller.GetAllUser)
		user.GET("/getbyid", controller.GetUserById)
		user.PUT("/update", controller.UpdateUser)
		user.DELETE("/delete", controller.DeleteUser)

		carshare := user.Group("carshare")
		{
			carshare.POST("/add", controller.AddCarShare)
			carshare.GET("/getbyid", controller.GetCarShareById)
			carshare.GET("/getbyuser", controller.GetCarShareByUser)
			carshare.PUT("/update", controller.UpdateCarShare)
			carshare.DELETE("/delete", controller.DeleteCarShare)
		}

		teacher := user.Group("teacher")
		{
			teacher.POST("/add", controller.AddTeacher)
			teacher.GET("/getbyid", controller.GetTeacherById)
			teacher.PUT("/update", controller.UpdateTeacher)
			teacher.DELETE("delete", controller.DeleteTeacher)
		}
	}
	e.GET("/carshare/getall", controller.GetAllCarShare)
	e.GET("/carshare/getbydestination", controller.GetCarShareByDestination)

	e.GET("/teacher/getall", controller.GetAllTeacher)
	e.GET("/teacher/getbynameandcourse", controller.GetTeacherByNameAndCourse)

	e.Run(":9090")
}
