package v1

import (
	"encoding/xml"
	"github.com/labstack/echo/v4"
	"myapp/internal/constants"
	"myapp/internal/helpers"
	"myapp/internal/models"
	"myapp/internal/shared/payloads"
	"net/http"
	"strconv"
)

func (h *Handler) getAllPosts(c echo.Context) error {
	responseType := c.QueryParam("type")
	if responseType != constants.JSON && responseType != constants.XML {
		return c.JSON(http.StatusBadRequest, helpers.Res("unknown type of response"))
	}

	postsFromDB, err := h.services.Post.GetAllPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not get posts from database"))
	}

	if responseType == constants.XML {
		type allPosts struct {
			XMLName xml.Name      `xml:"posts"`
			Posts   []models.Post `xml:"post"`
		}
		return c.XML(http.StatusOK, &allPosts{Posts: postsFromDB})
	}
	return c.JSON(http.StatusOK, &postsFromDB)
}

func (h *Handler) getPost(c echo.Context) error {
	responseType := c.QueryParam("type")
	if responseType != constants.JSON && responseType != constants.XML {
		return c.JSON(http.StatusBadRequest, helpers.Res("unknown type of response"))
	}

	postIdFromReq, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}

	postFromDB, err := h.services.Post.GetPost(postIdFromReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("could not find post for the specified id"))
	}

	if responseType == constants.XML {
		return c.XML(http.StatusOK, postFromDB)
	}
	return c.JSON(http.StatusOK, postFromDB)
}

func (h *Handler) createPost(c echo.Context) error {
	post := new(models.Post)
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(post); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	post.UserID = c.Get("user_id").(int) // from the authorization middleware;
	err := h.services.Post.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.Res("post was created"))
}

// updates post in database
func (h *Handler) updatePost(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}
	postFromReq := new(payloads.UpdatePostPayload)
	if err := c.Bind(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}
	if err := c.Validate(postFromReq); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	userIdFromToken := c.Get("user_id").(int)
	err = h.services.Post.UpdatePost(userIdFromToken, postId, postFromReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusAccepted, helpers.Res("post was updated"))
}

// deletes post in database
func (h *Handler) deletePost(c echo.Context) error {
	postIdFromReq, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res("postID was not in the url"))
	}

	userIdFromToken := c.Get("user_id").(int)
	err = h.services.Post.DeletePost(userIdFromToken, postIdFromReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Res(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.Res("post was deleted"))
}