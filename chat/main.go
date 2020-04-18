// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Type string
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "refresh/chat/home.html")
}

type logger interface {
	Success(msg interface{}, args ...interface{})
	Error(msg interface{}, args ...interface{})
	Print(msg interface{}, args ...interface{})
}

type Web struct {
	logger
	cancelFunc func()
	hub        *Hub
}

func (w *Web) Success(msg interface{}, args ...interface{}) {
	w.hub.broadcast <- []byte("[success] " + fmt.Sprintf(fmt.Sprint(msg), args...))
	w.logger.Success(msg, args...)
}

func (w *Web) Error(msg interface{}, args ...interface{}) {
	w.hub.broadcast <- []byte("[error] " + fmt.Sprintf(fmt.Sprint(msg), args...))
	w.logger.Error(msg, args...)
}

func (w *Web) Print(msg interface{}, args ...interface{}) {
	w.hub.broadcast <- []byte("[print] " + fmt.Sprintf(fmt.Sprint(msg), args...))
	w.logger.Print(msg, args...)
}

func Start(ctx context.Context, cancelFunc func(), logger logger) *Web {
	hub := newHub(cancelFunc)
	go hub.run()

	web := &Web{logger: logger, cancelFunc: cancelFunc, hub: hub}

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	go func() {
		err := http.ListenAndServe(":8082", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	return web
}
