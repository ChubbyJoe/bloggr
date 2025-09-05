package utility

import (
	"time"

	"github.com/ChubbyJoe/bloggr/models"
)

type SignUpInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint          `json:"id"`
	Username  string        `json:"username"`
	Blogs     []models.Blog `json:"blogs"`
	CreatedAt time.Time     `json:"createdAt"`
}
