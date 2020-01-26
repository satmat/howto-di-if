// 本パッケージは、サービスの仕様に基づいた処理が記述されることを想定している。
// 本パッケージをユニットテストする際、dbパッケージの処理は実際にはDBへアクセスさせたくないので、フェイクにしたい。
package service

import (
	"fmt"
	"github.com/satmat/howto-di-if/db"
	"unicode"
)

//
// (1) GetServiceData フェイクにしたいメソッドを定義したInterface型を引数に持つ場合
//
func GetServiceData(s db.SessionInterface, k string) string {

	// プロダクションコード中では
	// s の実態は db.Session である。
	// dbパッケージの Select が呼ばれる。

	// テストコード中では、
	// s の実態は service_test.go の FakeSession である。
	// service_test.go の Select が呼ばれる。

	result, err := s.Select(k)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return ""
	}

	// 結果の1文字目が大文字なら、参照された文字を大文字にし、
	// そうでなければそのままにする、という仕様だとして、
	// それをテストしたいと仮定する。
	if result != "" && unicode.IsUpper([]rune(k)[0]) {
		ur := ""
		for _, v := range []rune(result) {
			ur += string(unicode.ToUpper(v))
		}
		result = ur
	}

	return result
}

//
// (2) フェイクにしたいメソッドを定義したInterfaceをメンバに持つ構造体をレシーバに持つ場合
//

// Service SessionInterfaceをメンバに持つ構造体
// サービス仕様で必要な変数メンバを持った構造体を想定しているが
// 良い例えが思いつかなかったのでsessionだけ持っている
type Service struct {
	session db.SessionInterface
}

// NewService Service構造体のコンストラクタ
func NewService(s db.SessionInterface) *Service {
	return &Service{
		session: s,
	}
}

// PutServiceData サービス仕様に基づいてInsertを実行するメソッド
func (s *Service) PutServiceData(k string, v string) {

	// 正常ケースでは何もしないが、
	// 異常ケースではレシーバにエラー文字列

	err := s.session.Insert(k, v)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
}

// (3) SessionInterfaceWrapper 例えばdbが外部ライブラリだったとして、interfaceが用意されていない場合、
// 自分のコード側にラッパとなるinterfaceを実装する方法。
// 大抵の場合、外部ライブラリを実行する一連の処理が1パッケージにまとまっていると思うので、
// 以下のうちDeleteServiceData以外は、サービスロジックの記述されたパッケージではなく、そのようなパッケージに実装したほうが良いと思う。
type SessionInterfaceWrapper interface {
	DeleteWrapper(k string) error
}

// (3) SessionWrapperで宣言したメソッドを、実際にレシーバとして持たせて実装させる構造体
type SessionWrapper struct {
	session *db.Session
}

// (3) SessionWrapperの構造体
func NewSessionWrapper(s *db.Session) SessionInterfaceWrapper {
	return &SessionWrapper{
		session: s,
	}
}

// (3) (*Session) Delete(string)のラッパ
func (s *SessionWrapper) DeleteWrapper(k string) error {
	return s.session.Delete(k)
}

// DeleteServiceData 実際にサービスロジックが記述されたメソッド。
// s は製品コードでは外部ライブラリのdb.Session
func DeleteServiceData(s SessionInterfaceWrapper, k string) {
	s.DeleteWrapper(k)
}
