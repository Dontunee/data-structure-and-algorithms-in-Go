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

func TestKnapSack(t *testing.T) {
	type args struct {
		weights, profits []int
		capacity         int
	}

	tests := []struct {
		name   string
		inputs args
		want   int
	}{
		{
			name: "brute force test",
			inputs: args{
				weights:  []int{1, 2, 3, 5},
				profits:  []int{1, 6, 10, 16},
				capacity: 7,
			},
			want: 22,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if received := RecursionKnapsack(test.inputs.weights, test.inputs.profits, test.inputs.capacity); received != test.want {
				t.Errorf("failed test")
			}
		})
	}
}

func TestKnapsack2dProblem(t *testing.T) {
	type args struct {
		weightsValue [][]int
		capacity     int
	}

	tests := []struct {
		name  string
		input args
		want  []interface{}
	}{
		{
			name: "valid numbers",
			input: args{
				[][]int{[]int{1, 2}, []int{4, 3}, []int{5, 6}, []int{6, 7}},
				10,
			},
			want: []interface{}{10, []int{1, 3}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if received := Knapsack2dProblem(test.input.weightsValue, test.input.capacity); received[1] != test.want[1] {
				t.Errorf("failed test ")
			}
		})
	}
}
