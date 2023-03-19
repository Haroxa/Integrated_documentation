package router

import (
	"github.com/Haroxa/Integrated_documentation/controller"
	"github.com/Haroxa/Integrated_documentation/middleware"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	//e.GET("/mail", controller.Mail)
	e.POST("/user/login", controller.Login)
	e.POST("/user/register", controller.Register)
	e.POST("/user/register/reg", controller.Reg)

	//e.POST("/user/register/reg", controller.Reg)

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
			carshare.DELETE("/deleteall", controller.DeleteAllCarShare)

			apply := carshare.Group("apply")
			{
				apply.POST("/add", controller.AddApply)
				apply.GET("/getbyid", controller.GetApplyById)
				apply.GET("/getbyuser", controller.GetApplyByUser)
				apply.PUT("/update", controller.UpdateApply)
				apply.DELETE("/delete", controller.DeleteApply)
				apply.DELETE("/deleteall", controller.DeleteAllApply)
			}
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
	e.GET("/carshare/getbytime", controller.GetCarShareByTime)

	e.GET("/carshare/apply/getall", controller.GetAllApply)
	e.GET("/carshare/apply/getbytime", controller.GetApplyByTime)
	e.GET("/carshare/apply/getbycarshare", controller.GetApplyByCarShare)

	e.GET("/teacher/getall", controller.GetAllTeacher)
	e.GET("/teacher/getbynameandcourse", controller.GetTeacherByNameAndCourse)

	e.Run(":9090")
}
