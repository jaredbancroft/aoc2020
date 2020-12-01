package expense

import "testing"

func TestReport_FindEntries(t *testing.T) {
	type fields struct {
		expenses []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "example",
			fields: fields{
				expenses: []int{1721, 979, 366, 299, 675, 1456},
			},
			args: args{
				target: 2020,
			},
			want: 514579,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Report{
				expenses: tt.fields.expenses,
			}
			if got := r.FindEntries(tt.args.target); got != tt.want {
				t.Errorf("Report.FindEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
