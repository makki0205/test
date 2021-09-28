package test

import (
	"reflect"
	"testing"
)

func TestLesson1(t *testing.T) {
	type args struct {
		names []string
	}
	tests := []struct {
		name string
		args args
		want Users
	}{
		{
			args: args{
				names: []string{"Mike", "Abbie", "Cheryl"},
			},
			want: Users{
				{
					Name:       "Mike",
					NameLength: 4,
				},
				{
					Name:       "Abbie",
					NameLength: 5,
				},
				{
					Name:       "Abbie",
					NameLength: 6,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lesson1(tt.args.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lesson1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLesson2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				n: 0,
			},
		},
		{
			name: "",
			args: args{
				n: 1,
			},
		},
		{
			name: "",
			args: args{
				n: 4,
			},
		},
		{
			name: "",
			args: args{
				n: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestLesson3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want Coin
	}{
		{
			name: "",
			args: args{
				n: 2345,
			},
			want: Coin{
				Thousand:    2,
				FiveHundred: 0,
				hundred:     3,
				Fifty:       0,
				Ten:         5,
				Five:        0,
			},
		},
		{
			name: "",
			args: args{
				n: 2872,
			},
			want: Coin{
				Thousand:    2,
				FiveHundred: 1,
				hundred:     3,
				Fifty:       1,
				Ten:         2,
				Five:        2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lesson3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lesson3() = %v, want %v", got, tt.want)
			}
		})
	}
}
