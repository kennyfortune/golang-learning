package testlearning

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// // handlerError 展示了t.Helper()来帮助封装error的检测，防止大量重复代码
// func handlerError(t testing.T, err error) {
// 	t.Helper()
// 	if err != nil {
// 		t.Fatalf("failed", err)
// 	}
// }

// Test_httpHandler 可以创建真实的网络连接进行测试，尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景。
func Test_httpHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	httpHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}

}
