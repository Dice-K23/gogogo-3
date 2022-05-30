// 学習教材: https://news.mynavi.jp/techplus/article/gogogo-3/
package main
import(
	"fmt"
	"math/rand"
	"net/http"
)
func main(){
	// 公開するディレクトリを指定する --- (*1)
	fs := http.FileServer(http.Dir("pub"))
	// サーバーURLとの対応を指定 --- (*2)
	http.Handle("/", fs)
	// 動的コンテンツの例としてサイコロも実装 --- (*3)
	http.HandleFunc("/dice", DiceHandler)
	// サーバー起動 --- (*4)
	http.ListenAndServe(":8888", nil)
}
// DiceHandler サイコロを返す
func DiceHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(fmt.Sprintf("<h1>サイコロ: %d</h1>", rand.Intn(6)+1)))
}