package post_handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/edwardkerckhof/goblog/internal/core/ports"
	responses "github.com/edwardkerckhof/goblog/pkg/utils"
)

type PostHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type PostHandlerImpl struct {
	postService ports.PostService
}

func NewHTTPHandler(postService ports.PostService) PostHandler {
	return &PostHandlerImpl{
		postService: postService,
	}
}

func (h *PostHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	post, err := h.postService.Get(uint(id))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}
