package test1

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "123",
			args: args{
				x: 1,
				y: 20,
			},
			want: 3,
		},
	}
	t.Log(tests)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
