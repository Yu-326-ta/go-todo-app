package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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

var validPath = regexp.MustCompile("^/todos/(edit|save|update|delete)/([0-9]+)$")

// リクエストURL内のidを取得するためのメソッド（最初はパターンとして覚える）
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validPathとURLが一致した部分をスライスで取得
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// q[0]がtodos、q[1]がedit、q[2]が1みたいな感じでスライスに入っている
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
		}

		// ここでいうfnはtodoEditのこと（ParseURLが呼ばれその後にtodoEditが呼ばれる）
		fn(w, r, qi)
	}
}

func StartMainserver() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	// パスの最後に/をつけない時、パスが完全に一致しないと処理をハンドラーに渡すことができない
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	// パスの最後に/をつけると、パスが完全に一致しなくても処理をハンドラーに渡すことができる(parseURLもハンドラーとして機能している)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	// 第二引数のnilはデフォルトのマルチプレックスを使用することの宣言、基本はデフォルトのマルチプレっクスを使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
