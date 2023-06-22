package controllers

import (
	"net/http"
	"todo_app/config"
)

func StartMainserver() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	// 第二引数のnilはデフォルトのマルチプレックスを使用することの宣言、基本はデフォルトのマルチプレっクスを使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
