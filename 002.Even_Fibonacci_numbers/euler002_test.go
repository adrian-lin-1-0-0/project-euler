package euler002

import "testing"

func TestFibEven(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "test1", args: args{n: 10}, want: 10},
		{name: "test2", args: args{n: 100}, want: 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FibEven(tt.args.n); got != tt.want {
				t.Errorf("FibEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
