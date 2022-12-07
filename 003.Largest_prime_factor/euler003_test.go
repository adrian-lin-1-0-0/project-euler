package euler003

import "testing"

func TestLargestPrimeFactor(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "test1", args: args{n: 10}, want: 5},
		{name: "test2", args: args{n: 17}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPrimeFactor(tt.args.n); got != tt.want {
				t.Errorf("LargestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
