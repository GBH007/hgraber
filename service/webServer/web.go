package webServer

import (
	"app/service/webServer/base"
	"app/service/webServer/static"
	"app/system"
	"context"
	"errors"
	"net/http"
	"time"
)

// Start запускает веб сервер
func Start(parentCtx context.Context, addr string, staticDir string) {
	ctx := system.NewSystemContext(parentCtx, "Web-srv")
	mux := http.NewServeMux()

	// обработчик статики
	if staticDir != "" {
		mux.Handle("/", http.FileServer(http.Dir(staticDir)))
	} else {
		mux.Handle("/", http.FileServer(http.FS(static.StaticDir)))
	}

	// обработчик файлов
	mux.Handle("/file/", http.StripPrefix("/file/", http.FileServer(http.Dir(system.GetFileStoragePath(ctx)))))

	// API
	base.AddHandler(mux, "/info", MainInfo())
	base.AddHandler(mux, "/new", NewTitle())
	base.AddHandler(mux, "/title/list", TitleList())
	base.AddHandler(mux, "/title/details", TitleInfo())
	base.AddHandler(mux, "/title/page", TitlePage())
	base.AddHandler(mux, "/to-zip", SaveToZIP())
	base.AddHandler(mux, "/app/info", AppInfo())
	base.AddHandler(mux, "/title/rate", SetTitleRate())
	base.AddHandler(mux, "/title/page/rate", SetPageRate())

	server := http.Server{
		Addr:        addr,
		Handler:     mux,
		ErrorLog:    system.StdErrorLogger(ctx),
		BaseContext: base.NewBaseContext(ctx),
	}

	system.AddWaiting(ctx)
	go func() {
		defer system.DoneWaiting(ctx)
		system.Info(ctx, "Запущен веб сервер")
		defer system.Info(ctx, "Веб сервер остановлен")

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			system.Error(ctx, err)
		}

	}()

	go func() {
		<-ctx.Done()
		system.Info(ctx, "Остановка веб сервера")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		system.IfErr(ctx, server.Shutdown(shutdownCtx))
	}()
}
