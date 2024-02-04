package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given_valid_email_when_check_isValidEmail_then_return_true",
			args: args{
				email: "lionel.messi@gmail.com",
			},
			want: true,
		},
		{
			name: "given_email_without_@_when_check_isValidEmail_then_return_false",
			args: args{
				email: "lionel.messigmail.com",
			},
			want: false,
		},
		{
			name: "given_empty_string_when_check_isValidEmail_then_return_false",
			args: args{
				email: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidEmail(tt.args.email)
			assert.Equal(t, tt.want, got)
		})
	}
}
