package hgraberworker

import (
	"app/internal/controller/internal/worker"
	"app/internal/domain/hgraber"
	"app/pkg/ctxtool"
	"context"
	"time"
)

func (c *Controller) servePageWorker(ctx context.Context) {
	const (
		interval      = time.Second * 15
		queueSize     = 10000
		handlersCount = 10
	)

	ctx = ctxtool.NewSystemContext(ctx, "worker-page")

	w := worker.New[hgraber.Page](
		queueSize,
		interval,
		c.logger,
		func(ctx context.Context, page hgraber.Page) {
			err := c.hgraberUseCases.LoadPageWithUpdate(ctx, page)
			if err != nil {
				c.logger.ErrorContext(ctx, err.Error())
			}
		},
		c.hgraberUseCases.GetUnsuccessPages,
	)

	c.register("page", w)

	w.Serve(ctx, handlersCount)
}
