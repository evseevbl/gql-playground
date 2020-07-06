package graph

import (
	"context"
	"sync"
	"time"

	"github.com/evseevbl/posts/internal/app/posts/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type foo interface{}

func NewResolver() *Resolver {
	return &Resolver{postStorage: make(map[int64]*model.Post)}
}

type Resolver struct {
	mu sync.Mutex

	idCounter   int64
	postStorage map[int64]*model.Post
}

func (r *Resolver) createPost(ctx context.Context, req *model.CreatePostReq) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.idCounter++
	post := &model.Post{
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}
	cnt := int(r.idCounter)
	post.ID = &cnt
	r.postStorage[r.idCounter] = post
	return post, nil
}
