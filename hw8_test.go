package hw8

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	N := 10
	b := 3
	_ , err := Pollard(N,b)
	s := fmt.Sprint(err)
	if s != "Test Failed" {
		t.Errorf("The test should fail with N = %d, b = %d",N,b)
	}
}

func Test2(t *testing.T) {
	N := 10
	b := 4
	got , err := Pollard(N,b)
	want := 5 
	if err != nil {
		t.Errorf("The test should work but you returned the error: %q",err)
	}
	if got != want {
		t.Errorf("You return %d as the common prime but it is %d.",got,want)
	}
}

func Test3(t *testing.T) {
	N := 11
	b := 100
	_ , err := Pollard(N,b)
	s := fmt.Sprint(err)
	if s != "Test Failed" {
		t.Errorf("The test should fail with N = %d, b = %d",N,b)
	}
}

func Test4(t *testing.T) {
	N := 61685147
	b := 45
	got , err := Pollard(N,b)
	want := 7841 
	if err != nil {
		t.Errorf("The test should work but you returned the error: %q",err)
	}
	if got != want {
		t.Errorf("You return %d as the common prime but it is %d.",got,want)
	}
}

func Test5(t *testing.T) {
	N := 58972723
	b := 100
	got , err := Pollard(N,b)
	want := 7481 
	if err != nil {
		t.Errorf("The test should work but you returned the error: %q",err)
	}
	if got != want {
		t.Errorf("You return %d as the common prime but it is %d.",got,want)
	}
}
