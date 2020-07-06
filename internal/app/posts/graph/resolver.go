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

func NewResolver() *Resolver {
	return &Resolver{
		postStorage: make([]*model.Post, 0),
		postChan: make(chan *model.Post, 1),
	}
}

type Resolver struct {
	mu sync.Mutex

	idCounter   int64
	postStorage []*model.Post
	postChan    chan *model.Post
}

func (r *Resolver) createPost(ctx context.Context, req *model.CreatePostReq) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post := &model.Post{
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}
	cnt := int(r.idCounter)
	post.ID = &cnt
	r.postStorage = append(r.postStorage, post)
	r.idCounter++
	r.postChan <- post
	return post, nil
}
