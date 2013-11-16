package main

import (
	"compress/gzip"
	"github.com/gorilla/pat"
	"io"
	"net/http"
	"strings"
)

func initRouter() {
	a := pat.New()

	a.Get("/sales/list", Zhhf(listSales))
	a.Get("/sale/compose", Zhhf(composeSale))
	a.Post("/sale/create", Zhhf(createSale))
	a.PathPrefix("/assets/").Handler(Zfh(Mah(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))))
	a.PathPrefix("/public/").Handler(Mah(http.StripPrefix("/public/", http.FileServer(http.Dir("public/")))))

	//a.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	a.Get("/", Zhhf(getHome))

	//http.Handle("/", zipHandler(a))
	http.Handle("/", a)
}

func serve() {
	initRouter()
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
	//sniffDone bool
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {

	/*if !w.sniffDone {
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", http.DetectContentType(b))
		}
		w.sniffDone = true
	}*/

	return w.Writer.Write(b)
}

//ZipFileHandler
func Zfh(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//w.Header().Set("Access-Control-Allow-Origin", "*")

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		/*
			if r.Method == "HEAD" {
				h.ServeHTTP(w, r)
				return
			}
		*/

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		h.ServeHTTP(&gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	})
}

//ZipHtmlHandlerFunc
func Zhhf(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "text/html")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		fn(&gzipResponseWriter{Writer: gz, ResponseWriter: w}, r)
	}
}

//maxAgeHandler
func Mah(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=31536000")
		h.ServeHTTP(w, r)
	})
}
