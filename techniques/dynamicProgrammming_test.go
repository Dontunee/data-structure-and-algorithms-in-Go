package techniques

import (
	"testing"
)

func Test_NSum_ShouldReturnNumber_WhenValidInputIsGiven(testing *testing.T) {
	//Arrange
	//Act
	result := NSum(5)

	//Assert
	if result != 15 {
		testing.Error("valid number not being returned")
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"simple test case",
			args{5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TabulationFibonacci(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
