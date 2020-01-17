package main

import (
	"github.com/satmat/howto-di-if/db"
	"github.com/satmat/howto-di-if/service"
)

func main() {

	s := db.NewSession("prodhost", "produser", "prodpass")

	// 例えばWebAPIのような実装であれば、本来であればhandlerの層を設けて
	// 受信パラメータのバリデーションや認証などをそこで実施するので、
	// mainではそのハンドラをWebAPIフレームワークに登録する処理で済むはずだが、
	// 説明のため、main()から直接読んでいる。

	service.GetServiceData(s, "key1")

	d := service.NewDB(s)
	d.PutServiceData("key2", "value2")

	sw := service.NewSessionWrapper(s)
	sw.DeleteWrapper("key3")
}
