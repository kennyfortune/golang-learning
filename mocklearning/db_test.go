package mocklearning

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	type args struct {
		db  DB
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{nil, "me"}, want: 1}, // TODO: Add test cases.
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("me")).Return(1, errors.New("not exist"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromDB(m, tt.args.key); got != tt.want {
				t.Errorf("GetFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
