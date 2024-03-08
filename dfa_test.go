package dfa_check

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFA_Check(t *testing.T) {
	type fields struct {
		badWords []string
	}
	type args struct {
		txt string
	}
	type want struct {
		found []string
		exist bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "no bad words",
			fields: fields{
				badWords: []string{"黑词"},
			},
			args: args{
				txt: "无敏感词文本",
			},
			want: want{
				found: nil,
				exist: false,
			},
		},
		{
			name: "have bad words",
			fields: fields{
				badWords: []string{"黑词", "违禁词"},
			},
			args: args{
				txt: "有敏感词文本，黑#wefwef词xxxx，违禁 ooohhg词，hh$%fppp",
			},
			want: want{
				found: []string{"黑#wefwef词", "违禁 ooohhg词"},
				exist: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := New()
			f.AddInvalidWords(tt.fields.badWords)
			found, exist := f.Check(tt.args.txt)
			assert.Equal(t, tt.want.found, found)
			assert.Equal(t, tt.want.exist, exist)
		})
	}
}
