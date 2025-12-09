package handlers

import (
	"errors"

	"L4.5/internal/model"
	"github.com/wb-go/wbf/ginext"
)

type statisticService interface {
	GetStats(nums model.Numbers) *model.Response
}

type Router struct {
	Engine     *ginext.Engine
	StatGetter statisticService
}

func New(e *ginext.Engine, gtr statisticService) (*Router, error) {
	if e == nil || gtr == nil {
		return nil, errors.New("[handlers] invalid Router parametrs")
	}
	return &Router{
		Engine:     e,
		StatGetter: gtr,
	}, nil
}

func (r *Router) Routes() {
	r.Engine.POST("/stats", r.StatsHandler)
}
