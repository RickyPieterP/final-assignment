package middleware

import (
	"fmt"
	uc "mygram/app/usecase"
	"mygram/app/usecase/request"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (m *middleware) AuthJwt(ctx *gin.Context) {

	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		resp := map[string]interface{}{
			"message": "Need Sign In",
		}
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	bearer := strings.Split(auth, "Bearer ")

	tokStr := bearer[1]

	tok, err := uc.ValidateToken(tokStr)
	fmt.Println(err, "errr")
	if err != nil {
		resp := map[string]interface{}{
			"message": "Need Sign In",
		}
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	ctx.Set("jwt", tok)
	ctx.Next()
}

func (m *middleware) RegisterUser(c *gin.Context) {
	var req request.RegisterUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	c.Set("body", req)
	c.Next()
}

func (m *middleware) LoginUser(c *gin.Context) {
	var req request.LoginUserReq

	err := c.ShouldBind(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	c.Set("body", req)
	c.Next()
}

func (m *middleware) UpdateUser(c *gin.Context) {
	var req request.UpdateUserReq

	err := c.ShouldBind(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	req.Id = jwt.Id

	c.Set("body", req)
	c.Next()
}

func (m *middleware) DeleteUser(c *gin.Context) {

	jwt := c.MustGet("jwt").(*uc.Token)

	c.Set("body", jwt)
	c.Next()
}

func (m *middleware) CreatePhoto(c *gin.Context) {
	var req request.CreatePhotoReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	fmt.Println(jwt.Id, "jwt id")
	req.UserId = jwt.Id

	c.Set("body", req)
	c.Next()
}

func (m *middleware) FindPhoto(c *gin.Context) {
	var req request.FindReq

	jwt := c.MustGet("jwt").(*uc.Token)
	req.UserId = jwt.Id

	c.Set("user_id", req)
	c.Next()
}

func (m *middleware) UpdatePhoto(c *gin.Context) {
	var req request.UpdatePhoto

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	req.Id = photoId
	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	req.UserId = jwt.Id

	c.Set("body", req)
	c.Next()
}

func (m *middleware) DeletePhoto(c *gin.Context) {
	jwt := c.MustGet("jwt").(*uc.Token)
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	c.Set("user_id", jwt.Id)
	c.Set("photo_id", photoId)
	c.Next()
}

func (m *middleware) CreateComment(c *gin.Context) {
	var req request.CreateCommentReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	fmt.Println(jwt.Id, "jwt id")
	req.UserIdComment = jwt.Id

	c.Set("body", req)
	c.Next()
}

func (m *middleware) FindComment(c *gin.Context) {
	var req request.FindCommentReq

	jwt := c.MustGet("jwt").(*uc.Token)
	req.UserId = jwt.Id

	c.Set("user_id", req)
	c.Next()
}

func (m *middleware) UpdateComment(c *gin.Context) {
	var req request.UpdateComment

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	req.Id = commentId
	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	req.UserId = jwt.Id

	c.Set("body", req)
	c.Next()
}

func (m *middleware) DeleteComment(c *gin.Context) {
	jwt := c.MustGet("jwt").(*uc.Token)
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	c.Set("user_id", jwt.Id)
	c.Set("comment_id", commentId)
	c.Next()
}

func (m *middleware) FindSocialMedia(c *gin.Context) {
	jwt := c.MustGet("jwt").(*uc.Token)

	c.Set("user_id", jwt.Id)
	c.Next()
}

func (m *middleware) CreateSocialMedia(c *gin.Context) {
	var req request.CreateSocialMediaReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)

	c.Set("body", req)
	c.Set("user_id", jwt.Id)
	c.Next()
}

func (m *middleware) UpdateSocialMedia(c *gin.Context) {
	var req request.UpdateSocialMediaReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	jwt := c.MustGet("jwt").(*uc.Token)
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	c.Set("body", req)
	c.Set("user_id", jwt.Id)
	c.Set("social_media_id", socialMediaId)
	c.Next()
}

func (m *middleware) DeleteSocialMedia(c *gin.Context) {
	jwt := c.MustGet("jwt").(*uc.Token)
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	c.Set("user_id", jwt.Id)
	c.Set("social_media_id", socialMediaId)
	c.Next()
}
