package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo_app/app/models"
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

func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	// route_auth.goのNamaで指定した値からクッキーを取得する
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainserver() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	// 第二引数のnilはデフォルトのマルチプレックスを使用することの宣言、基本はデフォルトのマルチプレっクスを使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
