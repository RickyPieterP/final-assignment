package transport

import (
	"fmt"
	"mygram/app/usecase"
	"mygram/app/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	usecase  usecase.Usecase
	response string
}

func NewTransport(usecase usecase.Usecase, response string) *Transport {
	return &Transport{
		usecase:  usecase,
		response: response,
	}
}

func (t *Transport) RegisterUser(c *gin.Context) {
	body := c.MustGet("body").(request.RegisterUserReq)

	res, httpStatus := t.usecase.RegisterUser(body)

	c.JSON(httpStatus, res)
}

func (t *Transport) LoginUser(c *gin.Context) {
	body := c.MustGet("body").(request.LoginUserReq)

	res, httpStatus, err := t.usecase.LoginUser(body)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.JSON(httpStatus, resp)
		return
	} else {
		c.JSON(httpStatus, res)
		return
	}
}

func (t *Transport) UpdateUser(c *gin.Context) {
	body := c.MustGet("body").(request.UpdateUserReq)

	res, httpStatus, err := t.usecase.UpdateUser(body)
	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.JSON(httpStatus, resp)
		return
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}

func (t *Transport) DeleteUser(c *gin.Context) {
	body := c.MustGet("body").(*usecase.Token)

	res, httpStatus, err := t.usecase.DeleteUser(*body)

	if err != nil {
		resp := map[string]interface{}{
			"message": err.Error(),
		}
		c.JSON(httpStatus, resp)
		return
	} else {
		c.JSON(http.StatusOK, res)
		return
	}
}

func (t *Transport) CreatePhoto(c *gin.Context) {
	body := c.MustGet("body").(request.CreatePhotoReq)

	res, httpStatus, err := t.usecase.CreatePhoto(&body)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) FindPhoto(c *gin.Context) {
	req := c.MustGet("user_id").(request.FindReq)
	res, httpStatus, err := t.usecase.FindPhoto(&req)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) UpdatePhoto(c *gin.Context) {
	req := c.MustGet("body").(request.UpdatePhoto)
	res, err := t.usecase.UpdatePhoto(&req)
	fmt.Println(err, "error")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (t *Transport) DeletePhoto(c *gin.Context) {
	user_id := c.MustGet("user_id").(int)
	photo_id := c.MustGet("photo_id").(int)
	res, err := t.usecase.DeletePhoto(photo_id, user_id)
	fmt.Println(err, "error")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (t *Transport) CreateComment(c *gin.Context) {
	body := c.MustGet("body").(request.CreateCommentReq)

	res, httpStatus, err := t.usecase.CreateComment(&body)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) FindComment(c *gin.Context) {
	req := c.MustGet("user_id").(request.FindCommentReq)
	res, httpStatus, err := t.usecase.FindComment(&req)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) UpdateComment(c *gin.Context) {
	req := c.MustGet("body").(request.UpdateComment)
	res, err := t.usecase.UpdateComment(&req)
	fmt.Println(res, "RES TP")
	fmt.Println(err, "error TP")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (t *Transport) DeleteComment(c *gin.Context) {
	user_id := c.MustGet("user_id").(int)
	comment_id := c.MustGet("comment_id").(int)
	res, err := t.usecase.DeleteComment(comment_id, user_id)
	fmt.Println(err, "error")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusCreated, res)
	}
}

func (t *Transport) CreateSocialMedia(c *gin.Context) {
	body := c.MustGet("body").(request.CreateSocialMediaReq)
	userID := c.MustGet("user_id")
	res, httpStatus, err := t.usecase.CreateSocialMedia(body, userID)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) FindSocialMedia(c *gin.Context) {
	userID := c.MustGet("user_id")
	res, httpStatus, err := t.usecase.FindSocialMedia(userID)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) UpdateSocialMedia(c *gin.Context) {
	body := c.MustGet("body").(request.UpdateSocialMediaReq)
	userID := c.MustGet("user_id")
	socialMediaID := c.MustGet("social_media_id")
	res, httpStatus, err := t.usecase.UpdateSocialMedia(body, userID, socialMediaID)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}

func (t *Transport) DeleteSocialMedia(c *gin.Context) {
	socialMediaID := c.MustGet("social_media_id")
	res, httpStatus, err := t.usecase.DeleteSocialMedia(socialMediaID)
	if err != nil {
		c.JSON(httpStatus, err)
		return
	} else {
		c.JSON(httpStatus, res)
	}
}
