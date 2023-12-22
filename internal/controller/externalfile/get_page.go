package externalfile

import (
	"app/internal/dto"
	"app/pkg/webtool"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (c *Controller) getPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		bookID, err := strconv.Atoi(r.Header.Get(dto.ExternalFileBookID))
		if err != nil {
			webtool.WritePlain(ctx, w, http.StatusBadRequest, err.Error())

			return
		}

		page, err := strconv.Atoi(r.Header.Get(dto.ExternalFilePageNumber))
		if err != nil {
			webtool.WritePlain(ctx, w, http.StatusBadRequest, err.Error())

			return
		}

		ext := r.Header.Get(dto.ExternalFilePageExtension)

		rawFile, err := c.fileStorage.OpenPageFile(ctx, bookID, page, ext)
		if err != nil {
			webtool.WritePlain(ctx, w, http.StatusBadRequest, err.Error())

			return
		}

		defer c.logger.IfErrFunc(ctx, rawFile.Close)

		rawData, err := io.ReadAll(rawFile)
		if err != nil {
			webtool.WritePlain(ctx, w, http.StatusBadRequest, err.Error())

			return
		}

		buff := bytes.NewReader(rawData)

		http.ServeContent(w, r, fmt.Sprintf("%d.%s", page, ext), time.Now(), buff)
	})
}