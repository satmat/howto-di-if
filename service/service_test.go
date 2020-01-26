package service

import (
	"fmt"
	"testing"
)

// FakeSession テストコード中でSessionInterfaceのメソッドを定義する構造体
type FakeSession struct {
	FakeSelect func(string) (string, error)
	FakeInsert func(string, string) error
}

// テストコード中では、service.goでのSelect実行時、
// dbパッケージではなく、こちらのSelectが実行される。
// 最終的に、テストケースで定義した関数変数が実行されるため、
// テスト用の振る舞いをさせることができる。
func (s FakeSession) Select(k string) (string, error) {
	return s.FakeSelect(k)
}

//
// (1) TestGetServiceData フェイクにしたいメソッドを定義したInterface型を引数に持つ場合
//
func TestGetServiceData(t *testing.T) {

	// 結果の1文字目が小文字のケース

	// テスト用の文字列を定義
	teststr := "this is test code. Select: key="

	// テスト用の振る舞いを定義
	fakesession := &FakeSession{}
	fakesession.FakeSelect = func(k string) (string, error) {
		ret := teststr + k
		fmt.Println(ret)
		return ret, nil
	}

	k := "testkey1"
	expect := teststr + k

	// 製品コード中のSession構造体の代わりに、
	// テストコード中のFakeSessionを渡す。
	actual := GetServiceData(fakesession, k)
	if expect != actual {
		t.Errorf("expect isn't same to actual.: expect=%v, actual=%v", expect, actual)
	}

	// 結果の1文字目が大文字のケース

	// テスト用の文字列を定義
	teststr = "This is test code. Select: key="

	// テスト用の振る舞いを定義
	fakesession.FakeSelect = func(k string) (string, error) {
		ret := teststr + k
		fmt.Println(ret)
		return ret, nil
	}

	k = "testkey2"
	// テスト用の期待値は愚直に大文字の文字列を用意することにした
	expect = "THIS IS TEST CODE. SELECT: KEY=TESTKEY2"
	actual = GetServiceData(fakesession, k)
	if expect != actual {
		t.Errorf("expect isn't same to actual.: expect=%v, actual=%v", expect, actual)
	}
}

//
// (2) フェイクにしたいメソッドを定義したInterfaceをメンバに持つ構造体をレシーバに持つ場合
//

// テストコード中では、service.goでのInsert実行時、
// dbパッケージではなく、こちらのInsertが実行される。
// 最終的に、テストケースで定義した関数変数が実行されるため、
// テスト用の振る舞いをさせることができる。
func (s *FakeSession) Insert(k string, v string) error {
	return s.FakeInsert(k, v)
}

// (2) SessionInterfaceをメンバに持つ構造体
func TestPutServiceData(t *testing.T) {

	// テスト用の振る舞いを定義
	fakesession := &FakeSession{}

	// kの1文字目が小文字のケース
	k := "testkey1"
	v := "testvalue1"
	expect := v
	fakesession.FakeInsert = func(fk string, fv string) error {
		actual := fv
		fmt.Println("This is test code. Select: key=" + fk + ", value=" + fv)
		if expect != actual {
			err := fmt.Errorf("expect isn't same to actual.: expect=%v, actual=%v", expect, actual)
			t.Errorf("%v", err)
			return err
		}
		return nil
	}

	// 製品コードと同様にService構造体を用いる。
	// 但しメンバsessionの実態はテストコード中のFakeSessionを渡す。
	s := &Service{session: fakesession}
	s.PutServiceData(k, v)

	// kの1文字目が大文字のケース
	k = "Testkey2"
	v = "testvalue2"
	expect = "TESTVALUE2"
}

func TestDeleteServiceData(t *testing.T) {
}
