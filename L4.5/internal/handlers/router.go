package handlers

import (
	"errors"
	"net/http/pprof"

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
	//API
	r.Engine.POST("/stats", r.StatsHandler)

	//pprof
	r.Engine.GET("/debug/pprof/", func(c *ginext.Context) { pprof.Index(c.Writer, c.Request) })
	r.Engine.GET("/debug/pprof/cmdline", func(c *ginext.Context) { pprof.Cmdline(c.Writer, c.Request) })
	r.Engine.GET("/debug/pprof/profile", func(c *ginext.Context) { pprof.Profile(c.Writer, c.Request) })
	r.Engine.GET("/debug/pprof/symbol", func(c *ginext.Context) { pprof.Symbol(c.Writer, c.Request) })
	r.Engine.GET("/debug/pprof/trace", func(c *ginext.Context) { pprof.Trace(c.Writer, c.Request) })
}
