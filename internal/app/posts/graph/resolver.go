package graph

import (
	"sync"

	"github.com/evseevbl/posts/internal/app/posts/graph/model"
)

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
	postReaders []chan *model.Post
}

