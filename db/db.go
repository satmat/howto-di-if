// 本パッケージは、何らかのデータストアへのアクセス処理を集約したパッケージを想定している。
// serviceのユニットテスト中ではフェイク化され、本物のデータストアにはアクセスされない。
package db

import (
	"fmt"
)

// SessionInterface データストアへアクセスするためのメソッドを持つinterfaceを想定
type SessionInterface interface {
	Select(string) (string, error)
	Insert(string, string) error
	// Delete はフェイク化対象が関数（interfaceが存在しない）の状況下での説明のため、あえて定義しない
}

// Session テストコード中でSessionInterfaceのメソッドを定義する構造体
// データストアの接続先情報の構造体を想定
type Session struct {
	Host string
	User string
	Pass string
}

// NewSession Session構造体のコンストラクタ
func NewSession(h string, u string, p string) *Session {
	return &Session{
		Host: h,
		User: u,
		Pass: p,
	}
}

//
// 以下、データストアへアクセスするためのメソッドを想定。
//

// Select メソッドの一例。
func (s *Session) Select(k string) (string, error) {
	ret := "this is production code. Select: key=" + k + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass
	fmt.Println(ret)
	return ret, nil
}

// Insert メソッドの一例。
func (s *Session) Insert(k string, v string) error {
	fmt.Println("this is production code. Insert: key=" + k + ", value=" + v + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass)
	return nil
}

// Delete メソッドの一例。
func (s *Session) Delete(k string) error {
	fmt.Println("this is production code. Delete: key=" + k + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass)
	return nil
}
