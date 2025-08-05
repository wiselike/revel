package revel

import (
	"net/http"
	"testing"
)

func NewTestController(w http.ResponseWriter, r *http.Request) *Controller {
	context := NewGoContext(nil)
	context.Request.SetRequest(r)
	context.Response.SetResponse(w)
	c := NewController(context)
	return c
}

func TestSignWithKey(t *testing.T) {
	key := []byte("abc123")
	msg := "ping"

	want := "064c703825c0f94b8f2f9e4f6914909a0374451e"

	if got := signWithKey(msg, key); got != want {
		t.Fatalf("Sign 结果不符合预期\nwant: %s\ngot : %s", want, got)
	}
}

func TestVerifyAndKeyRotation(t *testing.T) {
	// 设置第一把 key
	SetSecretKey("key1")
	msg := "payload"

	sigOld := Sign(msg)

	// 轮换到第二把 key
	SetSecretKey("key2")

	// 旧签名应继续通过
	if !Verify(msg, sigOld) {
		t.Fatalf("旧签名校验失败，应当通过")
	}

	// 再轮一次到第三把 key
	SetSecretKey("key3")

	// 新旧关系：latest=k3, previous=k2
	latest, prev := GetSecretKey()
	if string(latest) != "key3" || string(prev) != "key2" {
		t.Fatalf("只应保留最近两把 key，得到 latest=%q previous=%q", latest, prev)
	}

	// 用最新 key 重新签名并验证
	sigNew := Sign(msg)
	if !Verify(msg, sigNew) {
		t.Fatalf("最新签名校验失败，应当通过")
	}

	// 再次确认旧旧 key(k1) 已被丢弃，无法通过校验
	if Verify(msg, sigOld) {
		t.Fatalf("key1 已应被丢弃，但旧旧签名仍然通过")
	}
}
