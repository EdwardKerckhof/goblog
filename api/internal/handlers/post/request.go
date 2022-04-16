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

// GetOneParams arameters
// swagger:parameters GetOneParams
type GetOneParams struct {
	// in: path
	// required: true
	PostID uint `json:"id"`
}

// CreateParams parameters
// swagger:parameters CreateParams
type CreateParams struct {
	// in: body
	// required: true
	Body struct {
		Title string `json:"title" validate:"required,min=3,max=100"`
		Body  string `json:"body" validate:"required,min=10"`
	}
}

// DeleteParams arameters
// swagger:parameters DeleteParams
type DeleteParams struct {
	// in: path
	// required: true
	PostID uint `json:"id"`
}

// UpdateParams parameters
// swagger:parameters UpdateParams
type UpdateParams struct {
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
