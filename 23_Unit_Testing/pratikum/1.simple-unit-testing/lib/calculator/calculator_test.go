package lib

import (
	"testing"
)

func TestAddition(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "[Success] 1 + 1 = 2",
			args: args{
				a: 1,
				b: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "[Success] 20 / 10 = 2",
			args: args{
				a: 20,
				b: 10,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "[Error] 20 / 0 = Division by zero",
			args: args{
				a: 20,
				b: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := Division(tt.args.a, tt.args.b)

			if (err != nil) != tt.wantErr {
				t.Errorf("Division() error = %v, wantErr %v", err, tt.wantErr)
			}
			
			if !tt.wantErr {
				if got != tt.want {
					t.Errorf("Division() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}

func TestSubtraction(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "[Success] 1 - 1 = 0",
			args: args{
				a: 1,
				b: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtraction(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "[Success] 1 * 1 = 1",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplication(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}
