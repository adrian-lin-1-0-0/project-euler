package euler001

import "testing"

func TestMultiplesOf3And5(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "test1", args: args{n: 10}, want: 23},
		{name: "test1", args: args{n: 100}, want: 2318},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplesOf3And5(tt.args.n); got != tt.want {
				t.Errorf("MultiplesOf3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
