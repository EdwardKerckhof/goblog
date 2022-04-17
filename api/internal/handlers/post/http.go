package post_handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/edwardkerckhof/goblog/internal/core/ports"
	responses "github.com/edwardkerckhof/goblog/pkg/utils"
)

type PostHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type PostHandlerImpl struct {
	postService ports.PostService
}

// NewHTTPHandler creates a new HTTP Handler related to posts
func NewHTTPHandler(postService ports.PostService) PostHandler {
	return &PostHandlerImpl{
		postService: postService,
	}
}

// Get gets a post using the service
func (h *PostHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, _ := strconv.ParseUint(param, 10, 32)

	arg := GetOnePostParams{
		PostID: uint(id),
	}

	post, err := h.postService.Get(arg.PostID)
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		responses.ERROR(w, http.StatusNotFound, err)
		return
	case err != nil:
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, CreatePostResponse(post))
}
