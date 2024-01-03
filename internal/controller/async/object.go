package async

import (
	"context"
	"fmt"
)

type logger interface {
	Error(ctx context.Context, err error)
}

type Runner interface {
	Start(context.Context) (chan struct{}, error)
	Name() string
}

type AfterStopHandler func()

func New(logger logger) *Controller {
	return &Controller{
		logger: logger,
	}
}

type Controller struct {
	logger logger

	runnerChannels []chan struct{}
	runners        []Runner
	after          []AfterStopHandler
}

func (c *Controller) RegisterRunner(ctx context.Context, runner Runner) {
	c.runners = append(c.runners, runner)
}

func (c *Controller) RegisterAfterStop(ctx context.Context, handler AfterStopHandler) {
	c.after = append(c.after, handler)
}

func (c *Controller) Serve(parentCtx context.Context) error {
	ctx, cnl := context.WithCancel(parentCtx)
	defer cnl()

	for _, r := range c.runners {
		exitCh, err := r.Start(ctx)
		if err != nil {
			err = fmt.Errorf("start %s: %w", r.Name(), err)

			c.logger.Error(ctx, err)

			return err
		}

		c.runnerChannels = append(c.runnerChannels, exitCh)
	}

	// Дожидаемся завершения потоков
	for _, exitCh := range c.runnerChannels {
		<-exitCh
	}

	// Проходим по всем послеостановочным функциям
	for _, handler := range c.after {
		handler()
	}

	return nil
}
