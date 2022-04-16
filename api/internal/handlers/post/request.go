package post_handler

import "time"

// GetAllParams  parameters
// swagger:parameters GetAllParams
type GetAllParams struct {
	// in: query
	Limit uint `json:"limit"`
	// in: query
	Offset uint `json:"offset"`
}

// GetOnePostParams arameters
// swagger:parameters GetOnePostParams
type GetOnePostParams struct {
	// in: path
	// required: true
	PostID uint `json:"id"`
}

// CreatePostParams parameters
// swagger:parameters CreatePostParams
type CreatePostParams struct {
	// in: body
	// required: true
	Body struct {
		Title string `json:"title" validate:"required,min=3,max=100"`
		Body  string `json:"body" validate:"required,min=10"`
	}
}

// DeletePostParams arameters
// swagger:parameters DeletePostParams
type DeletePostParams struct {
	// in: path
	// required: true
	PostID uint `json:"id"`
}

// UpdatePostParams parameters
// swagger:parameters UpdatePostParams
type UpdatePostParams struct {
	// in: path
	// required: true
	PostId uint `json:"id"`

	// in: body
	Body struct {
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		UpdatedAt time.Time `json:"updated_at"`
	}
}
