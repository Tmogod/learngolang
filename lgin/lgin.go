package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strconv"
)

type Student struct {
	Name   string
	Age    int
	Height float64
}

type Request struct {
	StudentId string `json:"student_id"`
}

// 从redis上根据studentId获取Student实体
func GetStudentInfo(studentId string) Student {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.164.132:6379",
		Password: "ll1995ll",
		DB:       0,
	})
	ctx := context.TODO()
	stu := Student{}
	result := client.HGetAll(ctx, studentId).Val()

	for field, value := range result {
		if field == "Name" {
			stu.Name = value
		} else if field == "Age" {
			stu.Age, _ = strconv.Atoi(value)
		} else if field == "Height" {
			stu.Height, _ = strconv.ParseFloat(value, 10)
		}
	}
	return stu
}

func GetName(ctx *gin.Context) {
	param := ctx.Query("student_id") // 从Get请求中获取参数
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "please indidate student_id") // StatusBadRequest即400
		return
	}
	stu := GetStudentInfo(param)
	// ctx.String(http.StatusOK, stu.Name) // StatusOK即200
	ctx.JSON(http.StatusOK, stu)
	return
}

func GetAge(ctx *gin.Context) {
	param := ctx.PostForm("student_id") // 从post form获取参数
	if len(param) == 0 {                // 没有指定student_id参数
		ctx.String(http.StatusBadRequest, "please indidate student_id")
		return
	}
	stu := GetStudentInfo(param)
	// ctx.String(http.StatusOK, strconv.Itoa(stu.Age))
	ctx.JSON(http.StatusOK, stu)
	return
}

func GetHeight(ctx *gin.Context) {
	var request Request
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.String(http.StatusBadRequest, "please indidate student_id")
		return
	}
	stu := GetStudentInfo(request.StudentId)
	//ctx.String(http.StatusOK, strconv.FormatFloat(float64(stu.Height), 'f', -1, 64))
	ctx.JSON(http.StatusOK, stu)
	return
}

func main() {
	engine := gin.Default()
	engine.GET("/get_name", GetName)
	engine.POST("/get_age", GetAge)
	engine.POST("/get_height", GetHeight)

	err := engine.Run("127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
}
