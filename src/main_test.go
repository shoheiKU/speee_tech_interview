package main

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		num int64
		n_1 int64
		n_2 int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				num: 1,
				n_1: 1,
				n_2: 0,
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				num: 2,
				n_1: 1,
				n_2: 0,
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				num: 3,
				n_1: 1,
				n_2: 0,
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				num: 4,
				n_1: 1,
				n_2: 0,
			},
			want: 3,
		},
		{
			name: "test5",
			args: args{
				num: 10,
				n_1: 1,
				n_2: 0,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.num, tt.args.n_1, tt.args.n_2); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
