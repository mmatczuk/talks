package x

import (
	"net/http/httptest"
	"testing"
)

func TestUnpack(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/search?sid=id&l=foo&l=bar&max=100&x=true", nil)
	search(w, r)
}

func BenchmarkUnpack(b *testing.B) {
	debug = false
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/search?sid=id&l=foo&l=bar&max=100&x=true", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search(w, r)
	}
}
