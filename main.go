package main

import (
	"io"
	"net/http"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type writeCloser struct {
	io.Writer
}

func (wc writeCloser) Close() error {
	return nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		if url == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		qr, err := qrcode.New(url)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "image/jpeg")

		writer := standard.NewWithWriter(writeCloser{Writer: w})
		qr.Save(writer)
	})

	http.ListenAndServe(":8080", nil)
}
