package controllers

import (
	"api_ticketing_web/config"
	"api_ticketing_web/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var userForm model.UserForm
	var user model.User

	if err := c.ShouldBind(&userForm); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result := config.DB.Where("email = ?",userForm.Email).First(&user)

	if result.RowsAffected > 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error",
			"message":"Email sudah ada",
		})
		return
	}


	hashPassword,err := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error",
			"message":"Kesalahan enkripsi ",
		})
		return
	}

	userForm.Password = string(hashPassword)
	user = model.User{
		Name:     userForm.Name,
		Email:    userForm.Email,
		Password: string(hashPassword),
		Role:     userForm.Role,
	}
	if err := config.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"success":"success","message":"Berhasil membuat akun"})
	
}

func Login(c *gin.Context){
	var userForm model.LoginForm
	var user model.User

	if err := c.ShouldBind(&userForm); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	result := config.DB.Where("email = ?", userForm.Email).First(&user)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error", 
			"message":"Email dan Password Salah",
		})
		return 
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userForm.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"error", 
			"message":"Email dan Password Salah",
		})
		return 
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id" : user.ID,
		"name":user.Name,
		"exp":time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret)); 
	if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
			"status":"error", 
			"message":"Token gagal",
		})
		return 
	}

	c.JSON(http.StatusOK,gin.H{
		"status":"success",
		"token":tokenString,
		"message":"Berhasil login",
	})

}