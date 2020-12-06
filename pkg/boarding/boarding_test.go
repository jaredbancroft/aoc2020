package boarding

import (
	"reflect"
	"testing"
)

func TestNewPass(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want Pass
	}{
		{
			name: "example1",
			args: args{code: "BFFFBBFRRR"},
			want: Pass{Row: 70, Column: 7, SeatID: 567},
		},
		{
			name: "example2",
			args: args{code: "FFFBBBFRRR"},
			want: Pass{Row: 14, Column: 7, SeatID: 119},
		},
		{
			name: "example3",
			args: args{code: "BBFFBBFRLL"},
			want: Pass{Row: 102, Column: 4, SeatID: 820},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPass(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPass() = %v, want %v", got, tt.want)
			}
		})
	}
}
