// ファイル名の最後を_test.go
// カレントディレクトリ配下の全テストを実行したい場合、go test ./... として実行する

package calc1

import "testing"

// テストロジックとなる関数名はTest＊＊＊と始める必要がある。
func TestA(t *testing.T) {
	var nums []int = []int{1,2,3,4,5}
	if !(Summarize(nums) == 15) {
		t.Error(`miss`);
	}
}

func TestB(t *testing.T) {
	var nums []int = []int{1, 2, 3, 4, 5, 6}
	if !(Summarize(nums) == 20) {
		t.Error(`6 miss`)
	}
}