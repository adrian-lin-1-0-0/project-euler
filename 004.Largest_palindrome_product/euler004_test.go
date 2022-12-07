package euler004

import (
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{n: 12321}, want: true},
		{name: "test2", args: args{n: 12331}, want: false},
		{name: "test3", args: args{n: 123455}, want: false},
		{name: "test4", args: args{n: 321123}, want: true},
		{name: "test5", args: args{n: 101108}, want: false},
		{name: "test6", args: args{n: 800008}, want: true},
		{name: "test6", args: args{n: 321223}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromic(tt.args.n); got != tt.want {
				t.Errorf("IsPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestPalindromeProduct(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "test1", args: args{n: 101110}, want: 101101},
		{name: "test2", args: args{n: 800000}, want: 793397},
		{name: "test3", args: args{n: 999999}, want: 906609},
		{name: "test4", args: args{n: 101111}, want: 101101},
		{name: "test5", args: args{n: 900000}, want: 888888},
		{name: "test6", args: args{n: 793397}, want: 780087},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPalindromeProduct(tt.args.n); got != tt.want {
				t.Errorf("LargestPalindromeProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
