package ahttp

import (
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestParseBody(t *testing.T) {
	type args struct {
		r *http.Request
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Simple Request Body",
			args: args{
				r: &http.Request{
					Body: io.NopCloser(strings.NewReader(`{"msg": "Hello, World!"}`)),
				},
				v: &HelloWorld{},
			},
			want: &HelloWorld{
				Msg: "Hello, World!",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseBody(tt.args.r, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("ParseBody() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(tt.want, tt.args.v) {
				t.Errorf("ParseBody() = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}

// Test Interfaces
type HelloWorld struct {
	Msg string `json:"msg"`
}
