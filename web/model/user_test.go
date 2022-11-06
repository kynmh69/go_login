package model

import (
	"reflect"
	"testing"
	"time"
)

func TestSignUp(t *testing.T) {
	type args struct {
		userId   string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "sign up normal",
			args: args{"test_user1", "password"},
			want: &User{
				Id:        "",
				UserId:    "",
				Password:  "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignUp(tt.args.userId, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
