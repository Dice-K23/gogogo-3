// 学習教材: https://news.mynavi.jp/techplus/article/gogogo-3/
package main
import ("net/http")
func main(){
	// ハンドラを登録 --- (*1)
	/*
		webサーバーにアクセスがあった場合、どんな処理を行うかを指定
		ここでは、アクセスがあったら、HelloHandlerという関数を実行することを指定
	*/
	http.HandleFunc("/", HelloHandler) 
	
	// サーバーを起動 --- (*2)
	/*
		ポート番号8888でサーバを起動。
		http.ListenAndServe --- 第一引数にTCPのアドレス, 第二引数にhttp.Handler.
	*/
	http.ListenAndServe(":8888", nil) 
}
// HelloHandler サーバーの処理内容を記述 --- (*3)
/* 関数の書き方について
	func 関数名 (引数 型)　{ 処理 }
	func HelloHandler(w http.ResponseWriter, r *http.Request){ 処理 }
		HelloHandler --- 関数名
		w --- 第一引数
		http.ResponseWriter --- 型（"net/http"パッケージ内の構造体？）
		r --- 第二引数
		*http.Request --- 型(http.Requestはポインタ型？)
			→　参考
				https://jp.quora.com/Go%E8%A8%80%E8%AA%9E%E3%81%A7func-handle-w-http-ResponseWriter-r-http-Request-%E3%81%AEhttp-Request%E3%81%8C%E3%83%9D%E3%82%A4%E3%83%B3%E3%82%BF%E5%9E%8B%E3%81%A7%E3%81%82%E3%82%8B%E7%90%86%E7%94%B1%E3%81%AF
				https://stackoverflow.com/questions/13255907/in-go-http-handlers-why-is-the-responsewriter-a-value-but-the-request-a-pointer/56875204#56875204
	
*/
func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello,World!")) // --- (*4)
}

/* http.HandleFuncなどをきちんと理解する。
	https://journal.lampetty.net/entry/understanding-http-handler-in-go
	https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65
*/