package lib

import "testing"

func TestAddition(t *testing.T) {
	result := Addition(1, 1)
	if result != 2 {
		t.Errorf("Addition(1, 1) = %d; want 2", result)
	}
	result = Addition(-1, -1)
	if result != -2 {
		t.Errorf("Addition(-1, -1) = %d; want -2", result)
	}
	result = Addition(-1, 1)
	if result != 0 {
		t.Errorf("Addition(-1, 1) = %d; want 0", result)
	}
	result = Addition(1, -1)
	if result != 0 {
		t.Errorf("Addition(1, -1) = %d; want 0", result)
	}

}

func TestSubtraction(t *testing.T) {
	result := Subtraction(1, 1)
	if result != 0 {
		t.Errorf("Subtraction(1, 1) = %d; want 0", result)
	}
	result = Subtraction(-1, -1)
	if result != 0 {
		t.Errorf("Subtraction(-1, -1) = %d; want -2", result)
	}
	result = Subtraction(-1, 1)
	if result != -2 {
		t.Errorf("Subtraction(-1, 1) = %d; want -2", result)
	}
	result = Subtraction(1, -1)
	if result != 2 {
		t.Errorf("Subtraction(1, -1) = %d; want 2", result)
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(2, 2)
	if result != 4 {
		t.Errorf("Multiplication(2, 2) = %d; want 4", result)
	}
	result = Multiplication(-2, -2)
	if result != 4 {
		t.Errorf("Multiplication(-2, -2) = %d; want 4", result)
	}
	result = Multiplication(-2, 2)
	if result != -4 {
		t.Errorf("Multiplication(-2, 2) = %d; want -4", result)
	}
	result = Multiplication(0, 2)
	if result != 0 {
		t.Errorf("Multiplication(0, 2) = %d; want 0", result)
	}
}

func TestDivision(t *testing.T) {
	result := Division(4, 2)
	if result != 2 {
		t.Errorf("Division(4, 2) = %d; want 2", result)
	}
	result = Division(-4, -2)
	if result != 2 {
		t.Errorf("Division(-4, -2) = %d; want 2", result)
	}
	result = Division(-4, 2)
	if result != -2 {
		t.Errorf("Division(-4, 2) = %d; want -2", result)
	}
	result = Division(4, -2)
	if result != -2 {
		t.Errorf("Division(4, -2) = %d; want -2", result)
	}
	result = Division(0, 2)
	if result != 0 {
		t.Errorf("Division(0, 2) = %d; want 0", result)
	}
}