package expense

import (
	"testing"
)

func TestReport_Find2Entries(t *testing.T) {
	type fields struct {
		expenses []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example2",
			fields: fields{
				expenses: []int{1721, 979, 366, 299, 675, 1456},
			},
			args: args{
				target: 2020,
			},
			want:    514579,
			wantErr: false,
		},
		{
			name: "error2",
			fields: fields{
				expenses: []int{0, 0, 0, 0, 0, 0},
			},
			args: args{
				target: 2020,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Report{
				expenses: tt.fields.expenses,
			}
			got, err := r.Find2Entries(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.FindEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Report.FindEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReport_Find3Entries(t *testing.T) {
	type fields struct {
		expenses []int
	}
	type args struct {
		target int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "example3",
			fields: fields{
				expenses: []int{1721, 979, 366, 299, 675, 1456},
			},
			args: args{
				target: 2020,
			},
			want:    241861950,
			wantErr: false,
		},
		{
			name: "error3",
			fields: fields{
				expenses: []int{0, 0, 0, 0, 0, 0},
			},
			args: args{
				target: 2020,
			},
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Report{
				expenses: tt.fields.expenses,
			}
			got, err := r.Find3Entries(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Report.Find3Entries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Report.Find3Entries() = %v, want %v", got, tt.want)
			}
		})
	}
}
