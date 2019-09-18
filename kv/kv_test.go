package kv

import (
	"strconv"
	"testing"
)

var h *hashListKV
var m map[string]string

func init() {
	h = newHashListKV(1024000)
	m = make(map[string]string)
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		h.Set(s, s)
	}
}

func BenchmarkMapSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}
}

func Test_hashListKV_Set(t *testing.T) {

	type args struct {
		k string
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{}

	for i := 0; i < 10000; i++ {
		s := strconv.Itoa(i)
		tests = append(tests, struct {
			name string
			args args
			want string
		}{
			name: s,
			args: args{
				k: s,
				v: s,
			}, want: s})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.Set(tt.args.k, tt.args.v)
			if h.Get(tt.args.v) != tt.want {
				t.Errorf("unexpect result: want:%v, get: %v\n", tt.want, h.Get(tt.args.v))
			}
		})
	}
}
