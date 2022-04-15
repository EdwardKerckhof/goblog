package post_handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/edwardkerckhof/goblog/internal/core/ports"
	responses "github.com/edwardkerckhof/goblog/pkg/utils"
)

type HTTPHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	postService ports.PostService
}

func NewHTTPHandler(postService ports.PostService) HTTPHandler {
	return &handler{
		postService: postService,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
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
