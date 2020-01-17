// 本パッケージは何らかのデータストアへのアクセス処理をまとめたパッケージを想定している。
// ユニットテスト中ではフェイクにして、本物のデータストアにアクセスせず、任意の結果を上位に返してテストする。
package db

import (
	"fmt"
)

// SessionInterface 何らかのデータストアへアクセスするための層を想定
type SessionInterface interface {
	Select(string) (string, error)
	Insert(string, string) error
	// Delete は説明のため除外
}

// Session データストアの接続先情報の構造体を想定
type Session struct {
	Host string
	User string
	Pass string
}

// NewSession Session のコンストラクタ
func NewSession(h string, u string, p string) *Session {
	return &Session{
		Host: h,
		User: u,
		Pass: p,
	}
}

// (1) Select は Interfaceをレシーバに持つメソッドがサービスロジックに実装されている場合を想定
func (s *Session) Select(k string) (string, error) {
	ret := "this is production code. Select: key=" + k + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass
	fmt.Println(ret)
	return ret, nil
}

// (2) Insert は Interfaceをメンバに持つ構造体をレシーバに持つメソッドがサービスロジックに実装されている場合を想定
func (s *Session) Insert(k string, v string) error {
	fmt.Println("this is production code. Insert: key=" + k + ", value=" + v + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass)
	return nil
}

// (3) Delete は外部ライブラリ等でinterfaceが用意されていない場合を想定
func (s *Session) Delete(k string) error {
	fmt.Println("this is production code. Delete: key=" + k + ", host=" + s.Host + ", user=" + s.User + ", pass=" + s.Pass)
	return nil
}
