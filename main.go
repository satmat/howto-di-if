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

	// (1) フェイクにしたいメソッドを定義したInterface型を引数に持つ場合
	service.GetServiceData(s, "key1")

	// (2) フェイクにしたいメソッドを定義したInterfaceをメンバに持つ構造体をレシーバに持つ場合
	d := service.NewService(s)
	d.PutServiceData("key2", "value2")

	// (3) interfaceが用意されていないため自分のコード側にラッパを実装した場合
	sw := service.NewSessionWrapper(s)
	sw.DeleteWrapper("key3")
}
