package controller

import (
	"fmt"
	"github.com/Haroxa/Integrated_documentation/common"
	"github.com/Haroxa/Integrated_documentation/helper"
	"github.com/Haroxa/Integrated_documentation/model"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"

	"net/http"

	"strconv"
)

// 添加

func AddApply(c *gin.Context) {
	Userid := c.MustGet("user_id").(int)
	// 获取数据
	Apply := &model.Apply{}
	if err := c.ShouldBindBodyWith(Apply, binding.JSON); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}

	Apply.Userid = Userid
	// 检验存在性
	carshare, err := model.GetCarShareById(Apply.Carshareid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if carshare.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单不存在", nil))
		return
	}
	if carshare.Num == carshare.Maxnum {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单人数已达上限", nil))
		return
	}
	// 更新订单
	cmp := map[string]interface{}{
		"Pending": carshare.Pending + 1,
	}
	if err = model.UpdateCarShare(carshare, cmp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单更新失败", err))
		return
	}

	Apply.Status = "待审核"
	Apply.Createdtime = time.Now().In(common.ChinaTime).Format("2006-01-02 15:04")
	fmt.Printf("time: %v \n", Apply.Createdtime)
	// 创建
	if err = model.CreateApply(Apply); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "创建失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "创建成功", nil))
}

// 获取，通过 id

func GetApplyById(c *gin.Context) {
	Id := c.Query("Applyid")
	Applyid, _ := strconv.Atoi(Id)
	// 检验存在性
	Apply, err := model.GetApplyById(Applyid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if Apply.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "申请不存在", nil))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取成功", Apply))
}

// 获取，通过时间

func GetApplyByTime(c *gin.Context) {

	time := c.Query("time")
	// 获取
	Applys, count, err := model.GetApplyByTime(time)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", nil))
		return
	} //fmt.Println(Applys, "\n", count) //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}

	msg := fmt.Sprintf("获取成功,已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Applys))
}

// 获取，通过订单

func GetApplyByCarShare(c *gin.Context) {
	Carshareid := c.Query("carshareid")
	id, _ := strconv.Atoi(Carshareid)
	// 获取数据
	Applys, count, err := model.GetApplyByCarShare(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}

	msg := fmt.Sprintf("获取成功,已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Applys))
}

// 获取，通过用户

func GetApplyByUser(c *gin.Context) {
	Userid := c.MustGet("user_id").(int)
	// 获取数据
	Applys, count, err := model.GetApplyByUser(Userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}

	msg := fmt.Sprintf("获取成功,已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Applys))
}

// 获取所有

func GetAllApply(c *gin.Context) {
	Applys, count, err := model.GetAllApply()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	} //if count == 0 { //	c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeSuccess, "目前无订单", nil)) //	return //}

	msg := fmt.Sprintf("获取成功,已获取个数为:%d", count)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, Applys))
}

// 更新

func UpdateApply(c *gin.Context) {
	Id := c.Query("Applyid")
	Applyid, _ := strconv.Atoi(Id)
	// 检验存在性
	Apply, err := model.GetApplyById(Applyid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if Apply.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "申请不存在", nil))
		return
	}
	// 获取数据
	s := &model.Apply{}
	if err = c.ShouldBindJSON(s); err != nil {
		log.Errorf("Invalid Param %+v", errors.WithStack(err))
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "数据绑定失败", err))
		return
	}

	mp := structs.Map(s)
	mp["Status"] = ""
	for k, v := range mp {
		if v == "" || v == 0 {
			delete(mp, k)
		}
	}

	if Apply.Status != "申请已取消" && s.Status != "" {
		// 检验存在性
		carshare, err := model.GetCarShareById(Apply.Carshareid)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
			return
		}
		if carshare.Id == 0 {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单不存在", nil))
			return
		}
		cmp := make(map[string]interface{})
		if Apply.Status == "待审核" {
			cmp["Pending"] = carshare.Pending - 1
		}
		if s.Status == "通过" {
			if carshare.Num == carshare.Maxnum {
				c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单人数已达上限", nil))
				return
			}
			if Apply.Status != "申请成功" {
				cmp["Num"] = carshare.Num + 1
			}
			mp["Status"] = "申请成功"
		} else if s.Status == "不通过" {
			if Apply.Status == "申请成功" {
				cmp["Num"] = carshare.Num - 1
			}
			mp["Status"] = "申请失败"
		} else if s.Status == "取消" {
			if Apply.Status == "申请成功" {
				cmp["Num"] = carshare.Num - 1
			}
			mp["Status"] = "申请已取消"
		}
		// 更新订单
		if err = model.UpdateCarShare(carshare, cmp); err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单更新失败", err))
			return
		}
	}

	fmt.Printf("%v\n", mp)

	// 更新
	if err = model.UpdateApply(Apply, mp); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "更新失败", err))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "更新成功", nil))
}

// 删除

func DeleteApply(c *gin.Context) {
	Id := c.Query("Applyid")
	Applyid, _ := strconv.Atoi(Id)
	// 获取
	Apply, err := model.GetApplyById(Applyid)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if Apply.Id == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单不存在", nil))
		return
	}
	// 更新
	carshare, err1 := model.GetCarShareById(Apply.Carshareid)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err1))
		return
	}
	if carshare.Id != 0 {
		cmp := map[string]interface{}{
			"Pending": 0,
		}
		// 更新订单
		if err = model.UpdateCarShare(carshare, cmp); err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单更新失败", err))
			return
		}
	}

	// 删除
	if err = model.DeleteApply(Apply); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "删除成功", nil))
}

// 删除所有

func DeleteAllApply(c *gin.Context) {
	Applys, count, err := model.GetAllApply()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "目前无订单", nil))
		return
	}
	ex := make(map[int]int)
	for _, v := range Applys {
		// 检验存在性
		if ex[v.Carshareid] != 0 {
			continue
		}
		ex[v.Carshareid] = 1
		carshare, err := model.GetCarShareById(v.Carshareid)
		if err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "获取失败", err))
			return
		}
		if carshare.Id == 0 {
			continue
		}
		cmp := map[string]interface{}{
			"Pending": 0,
		}
		// 更新订单
		if err = model.UpdateCarShare(carshare, cmp); err != nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "订单更新失败", err))
			return
		}
	}
	// 删除
	if err = model.DeleteAllApply(Applys); err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiReturn(common.CodeError, "删除失败", err))
		return
	}
	msg := fmt.Sprintf("删除成功,已删除数为:%d", count)

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, msg, nil))
}
