// 本パッケージはサービスロジックが記述されることを想定している。
// 本パッケージをユニットテストしたい。dbパッケージはフェイクにしたい。
package service

import (
	"fmt"
	"github.com/satmat/howto-di-if/db"
)

// (1) GetServiceData Interfaceを引数に持つ場合
func GetServiceData(s db.SessionInterface, k string) {
	_, err := s.Select(k)
	if err != nil {
		fmt.Errorf("err=%v\n", err)
	}
}

// (2) DB SessionInterfaceをメンバに持つ構造体
type DB struct {
	session db.SessionInterface
}

// (2) DBのコンストラクタ
func NewDB(s db.SessionInterface) *DB {
	return &DB{
		session: s,
	}
}

// (2) PutServiceData Interfaceをメンバに持つ構造体をレシーバに持つ場合
func (d *DB) PutServiceData(k string, v string) {
	err := d.session.Insert(k, v)
	if err != nil {
		fmt.Errorf("err=%v\n, err")
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
