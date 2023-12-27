package api

import (
	"api/service"
	"api/service_chain"
	"api/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	var transferLog service.TransferLog
	var todayNum float64
	var todayCount float64
	logFlag, err := transferLog.IsExistByTime()
	if logFlag {
		todayNum, err = transferLog.GetOutNumInNow()
		if err != nil {
			res := &utils.Response{Code: 1007, Msg: "查询todayNum出错"}
			res.Json(c)
			return
		}
		todayCount, err = transferLog.GetOutCountInNow()
		if err != nil {
			res := &utils.Response{Code: 1008, Msg: "查询todayCount出错"}
			res.Json(c)
			return
		}
	}
	balance, err := service_chain.GetSourceBalance()
	if err != nil {
		res := &utils.Response{Code: 1009, Msg: "查询SourceBalance出错"}
		res.Json(c)
		return
	}
	blockNumber, err := service_chain.GetBlockNumber()
	if err != nil {
		res := &utils.Response{Code: 1010, Msg: "查询BlockNumber出错"}
		res.Json(c)
		return
	}
	c.HTML(http.StatusOK, "public/Faucet.html", gin.H{
		"blockNumber": blockNumber,
		"totalEth":    balance,
		"funded":      todayCount,
		"todayNum":    todayNum,
	})
	res := &utils.Response{Code: 0, Msg: "", Date: "服务正常"}
	res.Json(c)
}

func Faucets(c *gin.Context) {
	var transferInfo service.TransferInfo
	var transferLog service.TransferLog
	award, err := strconv.ParseFloat(c.PostForm("Award"), 64)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据结构出错"}
		res.Json(c)
		return
	}
	err = c.ShouldBind(&transferInfo)
	if err != nil {
		res := &utils.Response{Code: 1000, Msg: "数据结构出错"}
		res.Json(c)
		return
	}
	id, err := transferInfo.IsExist()
	if err != nil {
		res := &utils.Response{Code: 1001, Msg: "查询sender出错"}
		res.Json(c)
		return
	}
	logFlag, err := transferLog.IsExistByTime()
	if err != nil {
		fmt.Println(err)
		res := &utils.Response{Code: 1002, Msg: "查询time出错"}
		res.Json(c)
		return
	}
	fmt.Println("id", id)
	if id != -1 {
		transferInfo.ID = id
		err := transferInfo.UpdateInfo()
		if err != nil {
			res := &utils.Response{Code: 1003, Msg: "更新transferInfo出错"}
			res.Json(c)
			return
		}
	} else {
		err := transferInfo.AddTransferInfo()
		if err != nil {
			res := &utils.Response{Code: 1004, Msg: "添加transferInfo出错"}
			res.Json(c)
			return
		}
	}
	if logFlag {
		transferLog.OutCount += 1
		transferLog.OutNum += award
		err := transferLog.UpdateLog()
		if err != nil {
			res := &utils.Response{Code: 1005, Msg: "修改transferLog出错"}
			res.Json(c)
			return
		}
	} else {
		transferLog.OutCount = 1
		transferLog.OutNum = award
		err := transferLog.AddLog()
		if err != nil {
			res := &utils.Response{Code: 1006, Msg: "添加transferLog出错"}
			res.Json(c)
			return
		}
	}
	res := &utils.Response{Code: 0, Msg: "", Date: "服务正常"}
	res.Json(c)
}
