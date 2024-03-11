package calc

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	testData := []struct {
		name      string
		args      args
		want      int
		wantErr   error
		isWantErr bool
	}{
		{
			name: "正常系1",
			args: args{
				x: 1,
				y: 2,
			},
			want:      3,
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "正常系2",
			args: args{
				x: 0,
				y: 0,
			},
			want:      0,
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "正常系3",
			args: args{
				x: 50,
				y: 50,
			},
			want:      100,
			wantErr:   nil,
			isWantErr: false,
		},
		{
			name: "異常系",
			args: args{
				x: -1,
				y: -2,
			},
			want:      0,
			wantErr:   errors.New("x and y must be greater than 0"),
			isWantErr: true,
		},
	}
	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// 並列テスト
			t.Parallel()
			got, err := Add(tt.args.x, tt.args.y)
			if (err != nil) != tt.isWantErr {
				t.Error("isWantErr does't match bool(err != nil)")
			}
			if err != nil {
				if tt.wantErr == nil {
					t.Fatalf("err.Error() = %v, want error is nil", err.Error())
				}
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("err msg got:%v, want:%v", err.Error(), tt.wantErr.Error())
				}
			}
			opts := []cmp.Option{}
			if diff := cmp.Diff(got, tt.want, opts...); diff != "" {
				t.Errorf("Add() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
