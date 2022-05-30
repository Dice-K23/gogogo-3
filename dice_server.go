package main
import(
	"fmt"
	"math/rand"
	"net/http"
)
func main(){
	// サーバーの起動 --- (*1)
	http.HandleFunc("/", DiceHandler)
	http.ListenAndS
erve(":8888", nil)
}
// DiceHandler サーバーの処理内容
func DiceHandler(w http.ResponseWriter, r *http.Request){
	// サイコロの目と出力メッセージを生成 --- (*2)
	v := rand.Intn(6) + 1
	s := fmt.Sprintf("サイコロの目: %d", v)
	/*
	fmt.Sprintfについて
	https://qiita.com/Sekky0905/items/5a65602dce83551184b3
	*/
	// 書き出す --- (*3)
	w.Write([]byte(s))
}