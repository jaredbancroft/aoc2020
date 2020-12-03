package toboggan

import "testing"

func TestTreeMap_Traverse(t *testing.T) {
	type fields struct {
		depth int
		width int
		lines []string
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "expect7",
			fields: fields{
				depth: 11,
				width: 11,
				lines: []string{
					"..##.......",
					"#...#...#..",
					".#....#..#.",
					"..#.#...#.#",
					".#...##..#.",
					"..#.##.....",
					".#.#.#....#",
					".#........#",
					"#.##...#...",
					"#...##....#",
					".#..#...#.#",
				},
			},
			args: args{
				x: 3,
				y: 1,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := &TreeMap{
				depth: tt.fields.depth,
				width: tt.fields.width,
				lines: tt.fields.lines,
			}
			if got := tm.Traverse(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("TreeMap.Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
