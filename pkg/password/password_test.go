package password

import (
	"testing"
)

func TestValidator_Validate(t *testing.T) {
	type fields struct {
		min      int
		max      int
		char     string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "example1",
			fields: fields{
				min:      1,
				max:      3,
				char:     "a",
				password: "abcde",
			},
			want: true,
		},
		{
			name: "example2",
			fields: fields{
				min:      1,
				max:      3,
				char:     "b",
				password: "cdefg",
			},
			want: false,
		},
		{
			name: "example3",
			fields: fields{
				min:      2,
				max:      9,
				char:     "c",
				password: "ccccccccc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				min:      tt.fields.min,
				max:      tt.fields.max,
				char:     tt.fields.char,
				password: tt.fields.password,
			}
			if got := v.ValidateNumber(); got != tt.want {
				t.Errorf("Validator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_ValidatePosition(t *testing.T) {
	type fields struct {
		min      int
		max      int
		char     string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "example4",
			fields: fields{
				min:      1,
				max:      3,
				char:     "a",
				password: "abcde",
			},
			want: true,
		},
		{
			name: "example5",
			fields: fields{
				min:      1,
				max:      3,
				char:     "b",
				password: "cdefg",
			},
			want: false,
		},
		{
			name: "example6",
			fields: fields{
				min:      2,
				max:      9,
				char:     "c",
				password: "ccccccccc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Validator{
				min:      tt.fields.min,
				max:      tt.fields.max,
				char:     tt.fields.char,
				password: tt.fields.password,
			}
			if got := v.ValidatePosition(); got != tt.want {
				t.Errorf("Validator.ValidatePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
