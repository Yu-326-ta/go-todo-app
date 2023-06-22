package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo_app/config"
)

// ...stringは可変長引数
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	// Mustはテンプレートをキャッシュしておくことで効率的に処理できるようにしている、引数のParseFilesがエラーの時、パニックで処理止めれる
	templates := template.Must(template.ParseFiles(files...))
	// defineで定義したテンプレートの読み込み、1引数にレスポンスライター、2にテンプレートネーム
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainserver() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	// 第二引数のnilはデフォルトのマルチプレックスを使用することの宣言、基本はデフォルトのマルチプレっクスを使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
