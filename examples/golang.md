# Go言語入門

## 目次

* [Go言語とは](#Go言語とは)
* [Hello Go](#hello-go)
* [変数](#変数)
* [型](#型)
* [関数](#関数)
* [ループ](#ループ)
* [構造体とメソッド](#構造体とメソッド)
* [インタフェース](#インタフェース)
* [型変換と型アサーション](#型変換と型アサーション)
* [配列とスライス](#配列とスライス)
* [マップ](#マップ)
* [エラー処理](#エラー処理)
* [ゴルーチン](#ゴルーチン)
* [パッケージ](#パッケージ)
* [デバッグ](#デバッグ)
* [テスト](#テスト)

## Go言語とは

* [The Go Programming Language](http://golang.org/)

Googleによって開発されたオープンソースのプログラミング言語

### 特徴

* シンプル
* コンパイル・実行速度が早い
* 充実した標準パッケージ
* 並行処理が容易
* ポータビリティ
* 強力なツール
* ダックタイピング
* イントロスペクション
* 型推論
* GC
* [ʕ◔ϖ◔ʔ](http://goo.gl/V7BEqn)

### 注目度

[Google トレンド](http://www.google.co.jp/trends/explore#q=%2Fm%2F09gbxjr%2C%20golang&date=11%2F2009%2053m&cmpt=q)

### 実装やサービス

* Google App Engine
* Docker, Packer, Serf
* Vitess
* CloudFlare

[etc.](https://code.google.com/p/go-wiki/wiki/Projects)

## Hello Go

http://play.golang.org/p/AzJcxtdp5A

* セミコロン不要
* 実行にはmain()を含むmainパッケージが必要
* importで外部パッケージをインポート

## 変数

http://play.golang.org/p/wRzPMqQA-o

* 型は最後に書く
* := で型推論
* ポインタあり（ただしポインタ演算なし[^3]）
* new(T)は初期化（ゼロ初期化[^4]）してポインタを返す
* makeは参照型（スライス、マップ、チャネル）を初期化して値を返す

[^3]: #変数 "実際にはunsafeパッケージをつかえばポインタ演算可能ですが忘れましょう"
[^4]: #変数 "新たな変数を宣言した場合やnew(T)、makeで値を生成した場合は対応する型のゼロ値に初期化されます。bool型ならfalse、int型なら0、string型なら空文字、ポインタやスライス、チャネルならnilとなります。"

## 型

* 論理値型
* 数値型
* 文字列型
* 配列型
* スライス型
* 構造体型
* ポインタ型
* 関数型
* インタフェース型
* マップ型
* チャネル型

### 宣言済みの型

```
bool       : 論理値（true, false）
uint8      : 符号なし 8ビット 整数
uint16     : 符号なし 16ビット 整数
uint32     : 符号なし 32ビット 整数
uint64     : 符号なし 64ビット 整数
int8       : 符号あり 8ビット 整数
int16      : 符号あり 16ビット 整数
int32      : 符号あり 32ビット 整数
int64      : 符号あり 64ビット 整数
float32    : IEEE-754 32-ビット 浮動小数値
float64    : IEEE-754 64-ビット 浮動小数値
complex64  : float32の実数部と虚数部を持つ複素数
complex128 : float64の実数部と虚数部を持つ複素数
byte       : uint8の別名
rune       : int32の別名
string     : 文字列（値は不変）
```

### 実装に依存する宣言済みの型

```
uint    : 32または64ビット
int     : uintと同じサイズ
uintptr : ポインタの値を格納するのに充分な大きさの符号なし整数
```

## 関数

http://play.golang.org/p/l7omoy0-n1

* funcで関数を宣言
* ...Tで可変個引数（引数は[]Tとなる）
* interface{}（空インタフェース型）で任意の型を受け取る
* スライスを可変個引数として渡す場合は、末尾に...をつける
* 関数は複数の値を返すことが可能
* 関数の返り値は全て無視かすべて受け取る。一部を無視したい場合は\_（ブランク識別子）を使用する
* クロージャ

## ループ

http://play.golang.org/p/-8jQIbBH3c

* ループはforのみ
* ループの外に抜けるにはcontinue, break
* ネストされたループを抜けるにはラベルを使用する
* ++（--）は後置のみ。また式ではないため値へ評価されないので a := b++ などはNG

## 構造体とメソッド

http://play.golang.org/p/LZQaLcVV6P

* 型レベルでのアクセス指定子はないが、下記のルールでパッケージレベルでアクセス制御される
  - [大文字](http://www.fileformat.info/info/unicode/category/Lu/list.htm)で始まるトップレベルの型、メソッド、変数名、構造体のフィールドはエクスポートされる(public)
  - それ以外はエクスポートされない(private)
* 継承はないのでprotectedなフィールドはない
* フィールドにはタグを指定でき、reflectパッケージから参照可能
* 既存の型に名前付けをして別の型（具象型）を定義できる
* メソッドのレシーバは値型でもポインタ型でも定義できる
* レシーバが値型の場合は、値・ポインタ両方に対して呼び出し可能
* [レシーバをポインタとする場合の指針](http://golang.org/doc/faq#methods_on_values_or_pointers)
  - レシーバのフィールドに対して変更を行う場合
  - レシーバが巨大な場合
  - メソッドの一貫性を持たせる場合

### 型埋め込み

http://play.golang.org/p/xwulpHHKZV

* 型埋め込みによる暗黙の委譲（継承ではない）
* レシーバはフィールドであり、フィールドを含む構造体ではないので注意
* 同名のメソッドの場合は呼び出すフィールドを明示する

## インタフェース

http://play.golang.org/p/NWPW97YdVv

* インタフェースはメソッドの集合
* インタフェースでダックタイピングを実現
* 型埋め込み同様、インタフェースを埋め込むことが可能。インタフェースの継承と等価 [^1]
* インタフェースが定義しているメソッドをすべて実装している限り、インタフェース型を持つ変数へどんな型でも代入可能
* 直接のフィールドへのアクセスが必要ない場合は、構造体自身ではなくパブリックのメソッドを定義したインタフェースだけを公開することで実装を隠蔽する [^2]

[^1]: #インタフェース "インタフェースコンポジション"
[^2]: #インタフェース "サンプルは同一パッケージ(main)なのでアクセス可能"

## 型変換と型アサーション

http://play.golang.org/p/FJULgGKYZY

**※play.golang.orgではエラーになってしまう。コピーして手元で実行しよう！**

* 型変換はCのキャストのような感じ。細かいルールは[こちら](http://golang.jp/go_spec#Conversions)
* 型アサーションは実行時にチェックされる
  - 特定のインタフェースを実装した型なのか
  - 期待した型なのか
* インタフェース型じゃない値はinterface{}にキャストしたあとで型アサーションを行う
* 型switch文で型アサーション

## 配列とスライス

* 配列

  http://play.golang.org/p/CKhq1wMFb8

  - [...]で要素数を推論
  - 配列は通常の値
  - 別の配列への代入は要素のコピーとなる（ポインタではない）

* スライス

  http://play.golang.org/p/jkHphPU2uc

  - スライスは配列内の連続した領域への参照（配列に対するビュー）
  - スライスは配列へのポインタ(Data)、アクセスできる要素の数(Len)、最大のスライスの大きさ(Cap)をもつ

### スライスのイメージ

![](https://gist.githubusercontent.com/hayajo/9559874/raw/8c2b040450a17a0bf4e0d30fb35f6a67e8113355/go_slice.png)

### スライスの拡張と縮小

http://play.golang.org/p/TT5N-rRwlV

* スライスに要素を追加するにはappendを使用する
* スライスのCapが十分な大きさを持たない場合は、追加する値を含んだ新しい配列へのポインタをDataとするので注意が必要
* スライスにスライスをappendする場合は...を使用する
* スライスの縮小は、縮小したCapの新しいスライスを作り、既存のスライスから新しいスライスへ値をコピーする

### 配列とスライスのループ

http://play.golang.org/p/L9JWrwhBf1

* 配列やスライスのループにはrangeを使用する

## マップ

http://play.golang.org/p/ivM5DvEyed

* 他言語でいうところの連想配列や辞書
* make()で作成する。宣言やnew()で作ると初期化されないので注意
* キーに使用できる型は等価演算子が定義されていなければならない。詳しくは[こちら](http://golang.jp/go_spec#Comparison_operators)を参照
* マップのイテレートにもrangeを使用する

## エラー処理

http://play.golang.org/p/3esCtUQT4r

* （一般的には）戻り値としてのエラー値を検査することでエラー処理を行う

**余裕があれば解説**

* [エラー委譲パターン](http://play.golang.org/p/qyxAWlQxYx)

### リソースのクリーンアップ

http://play.golang.org/p/UpOjMtfUJJ

* deferでfinallyと似たようなことができ、関数の実行を遅延させることができる
* deferに指定された関数は、それが書かれている関数が終了する時点で実行される

### パニックとリカバー

http://play.golang.org/p/bVOyhYt2zN

* panic(), recover()でtry-catch-finally的な実装が可能
* deferで呼ばれる関数内であればrecover()は機能するので、呼び出す関数にrecover()を記述することで例外処理をまとめることもできる
* が、panic()は明らかに回復手段がない場合にだけ使用し、また、回復するのが絶対的に安全で無い限りrecover()は使用しない

## ゴルーチン

http://play.golang.org/p/y4Vzz1nybW

* ゴルーチンは軽量スレッドのようなもの
* go 関数呼び出し で実行

### ゴルーチンの同期とバックグランド実行

http://play.golang.org/p/MAaDXrUqat

* ゴルーチンの同期にはミューテックスや条件変数([syncパッケージ](http://golang.org/pkg/sync/))を使用する方法もあるが、チャネルを使ったほうが自然

http://play.golang.org/p/bF5PBNwOO6

* バックグランドでの処理実行の終了待ちには[sync.WaitGroup](http://golang.org/pkg/sync/#WaitGroup)を使用する

### チャネル

http://play.golang.org/p/NJ0lFwH8UA

* チャネルはUNIXパイプのようなもので、ゴルーチン間の通信に使用する。型付けされる
* make()で作成する

http://play.golang.org/p/dBQ4guoIu2

* チャネルはバッファリング可能

http://play.golang.org/p/X2DD-P_Ub1

* 複数チャネルを扱う場合はselectを使用する

### 並列化

```go
runtime.GOMAXPROCS(runtime.NumCPU())
```

* 同時に利用可能なCPU数はデフォルトで１。**並列**度を上げる場合は環境変数GOMAXPROCSもしくはruntime.GOMAXPROCS()で設定する

## パッケージ

```go
// user/user.go
package user

type User interface {
  Name() string
}

func NewUser(name string) User {
  if name != "" {
    return &user{name}
  }
  return new(guest)
}
```

```go
// user/types.go
package user

type user struct {
  name string
}

func (u user) Name() string {
  return u.name
}

type guest struct {}

func (_ guest) Name() string {
  return "Guest"
}
```

* [大文字](http://www.fileformat.info/info/unicode/category/Lu/list.htm)で始まるトップレベルの型、メソッド、変数名、構造体のフィールドはエクスポートされ、他のパッケージからアクセスできる(public)
* それ以外はエクスポートされず、同一パッケージからしかアクセス出来ない（private)
* パッケージは複数ファイルに分割できる

### 環境変数GOROOTとGOPATH

```shell
$ export GOPATH=$HOME/go
$ export PATH="$GOPATH/bin:$PATH"
$ go env GOROOT GOPATH
...
```

* GOROOT

  - 標準ライブラリを探すためのベースパスが設定される。JAVA_HOMEのようなもの。
  - 基本的には設定不要だが、Goをバイナリからインストールした際には設定が必要になる場合もある

* GOPATH

  - ワーキングディレクトリを設定する
  - go getやgo installなど、goツールのベースパスとなる
  - $GOPATH配下のディレクトリ構成
    * bin コンパイル後に生成される実行ファイルの格納先
    * pkg コンパイル後に生成されるパッケージの格納先
    * src ソースコードの保存先
      パッケージごとにサブディレクトリを作成
  - importの参照先は$GOROOT/pkgまたは$GOPATH/pkgの、アーキテクチャ配下のパスとなる

### パッケージのインストール

```shell
$ go get github.com/codegangsta/martini
$ ls $GOPATH/pkg/darwin_amd64/github.com/codegangsta
inject.a martini.a
```

* go get でサードパーティパッケージの取得とインストールを行う

### パッケージのビルド

```
$GOPATH
└── src
    ├── user
    │   ├── types.go
    │   └── user.go
    └── userapp
        └── main.go
```

---

* go install でパッケージのビルド(go build)と$GOPATHへのインストールを行う

  ```shell
  $ go install user
  ```

  ```
  $GOPATH
  ├── pkg
  │   └── darwin_amd64
  │        └── user.a
  └── src
      ├── user
      │   ├── types.go
      │   └── user.go
      └── userapp
          └── main.go
  ```

* ビルド対象がmainパッケージの場合は実行可能なバイナリが生成される

  ```shell
  $ go install userapp
  ```

  ```
  $GOPATH
  ├── bin
  │   └── userapp
  ├── pkg
  │   └── darwin_amd64
  │        └── user.a
  └── src
      ├── user
      │   ├── types.go
      │   └── user.go
      └── userapp
          └── main.go
  ```

  - ファイル名や出力先を変えたい場合は go build -o <OUTPUT_PATH> を実行する

    ```
    $ go build -o /tmp/fooapp userapp
    $ ls /tmp
    fooapp
    ```

* クロスコンパイルは環境変数GOOSとGOARCHを指定する（クロスコンパイル用の環境構築が必要）

  ```shell
  $ GOOS=windows GOARCH=386 go install userapp
  ```

  ```
  $GOPATH
  ├── bin
  │   ├── windows_386
  │   │   └── userapp.exe
  │   └── userapp
  ├── pkg
  │   ├── darwin_amd64
  │   │   └── user.a
  │   └── windows_amd64
  │       └── user.a
  └── src
      ├── user
      │   ├── types.go
      │   └── user.go
      └── userapp
          └── main.go
  ```

  - go buildでも同様

    ```
    $ GOOS=linux GOARCH=arm go build userapp
    ```

## デバッグ

* デバッグにはgdb(>= 7.1)を使う
* Goの最適化を無視するために -gcflags '-N -l' をつけてビルドする

  ```shell
  $ go build -gcflags '-N -l'
  ```

* runtime-gdb.pyをロードする（~/.gdbinitに書いてもOK）

  ```shell
  (gdb) source <$GOROOT/pkg/runtime/runtime-gdb.py>
  ```

### デバッグの簡単な流れ

```go
// myapp.go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a[9]) // panic
}
```

1. -gcflags '-N -l' をつけてビルドする

  ```shell
  $ go build -gcflags '-N -l' myapp.go
  ```

2. デバッグ開始

  ```shell
  $ gdb myapp
  ```

3. runtime-gdb.pyをロードする（必要な場合）

  ```shell
  (gdb) source <PATH_TO_runtime-gdb.py>
  ```

4. runtime.panicindex にブレークポイントを設定

  ```shell
  (gdb) break runtime.panicindex
  ```

5. 実行

  ```shell
  (gdb) run
  ...
  Breakpoint 1, runtime.panicindex () at ...
  ```

6. upして呼び出し元に戻る

  ```shell
  (gdb) up
  #1  0xXXXXXXXXXXXXXXXX in main.main () at /.../myapp.go:7
  7       fmt.Println(a[9])
  ```

7. ローカル変数の確認

  ```shell
  (gdb) info locals
  a =  []int = {0, 0, 0, 0, 0}
  ```

8. lenとcapの確認

  ```shell
  (gdb) p $len(a)
  $1 = 5
  (gdb) p $cap(a)
  $2 = 5
  ```

### テスト

```go
// user_test.go
package user

import(
	"testing"
	"reflect"
	"runtime"
)

func isType(t *testing.T, got interface{}, expected interface{}) {
	gotT := reflect.TypeOf(got)
	expectedT := reflect.TypeOf(expected)
	if gotT != expectedT {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("\nLocation: %s:%d\nError: got %s, expected %s", file, line, gotT, expectedT)
	}
}

func TestNewUser(t *testing.T) {
	g := NewUser("")
	isType(t, g, new(guest))

	u := NewUser("hayajo")
	isType(t, u, new(guest)) // fail
}
```

```shell
$ ls
types.go  user.go  user_test.go
$ go test
--- FAIL: TestNewUser (0.00 seconds)
        user_test.go:14:
                Location: /<PWD>/user_test.go:23
                Error: got *user.user, expected *user.guest
FAIL
exit status 1
FAIL    _/<PWD>      0.006s
```

* testingパッケージとgo testを使用してテストを行う
* ファイル名は\_test.goで終わるようにする
* Testで始まり(t \*testing.T)のシグニチャをもつ関数を順番に実行する

**テスト対象は "[パッケージ](#パッケージ)" のuserパッケージ**

