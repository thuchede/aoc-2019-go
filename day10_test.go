package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCompareVectorDirection(t *testing.T) {
	vector0a := Key{X: 1, Y: 1}
	vector0b := Key{X: 2, Y: 2}
	if !CompareVectorDirection(vector0a, vector0b) {
		t.Errorf("expected to have same direction")
	}
	vector1a := Key{X: 1, Y: 1}
	vector1b := Key{X: 3, Y: 3}
	if !CompareVectorDirection(vector1a, vector1b) {
		t.Errorf("expected to have same direction")
	}
	vector2a := Key{X: 1, Y: 1}
	vector2b := Key{X: 3, Y: 2}
	if CompareVectorDirection(vector2a, vector2b) {
		t.Errorf("expected to have different direction")
	}
	vector3a := Key{X: 1, Y: 1}
	vector3b := Key{X: -1, Y: -1}
	if CompareVectorDirection(vector3a, vector3b) {
		t.Errorf("expected to have different direction")
	}
	vector4a := Key{X: -2, Y: -1}
	vector4b := Key{X: -4, Y: -2}
	if !CompareVectorDirection(vector4a, vector4b) {
		t.Errorf("expected to have same direction")
	}
}

func TestComputeVectorFromCoordinates0(t *testing.T) {
	point0a := Key{X: 1, Y: 1}
	point0b := Key{X: 2, Y: 2}
	if res := ComputeVectorFromCoordinates(point0a, point0b); res.X != 1 && res.Y != 1 {
		t.Errorf("expected (1, 1), got %v\n", res)
	}
	point1a := Key{X: 1, Y: 1}
	point1b := Key{X: 3, Y: 3}
	if res := ComputeVectorFromCoordinates(point1a, point1b); res.X != 2 && res.Y != 2 {
		t.Errorf("expected (2, 2), got %v\n", res)
	}
	point2a := Key{X: 1, Y: 1}
	point2b := Key{X: 3, Y: 2}
	if res := ComputeVectorFromCoordinates(point2a, point2b); res.X != 2 && res.Y != 1 {
		t.Errorf("expected (2, 1), got %v\n", res)
	}
	point3a := Key{X: 1, Y: 1}
	point3b := Key{X: -1, Y: -1}
	if res := ComputeVectorFromCoordinates(point3a, point3b); res.X != -2 && res.Y != -2 {
		t.Errorf("expected (-2, -2), got %v\n", res)
	}
	point4a := Key{X: -2, Y: -1}
	point4b := Key{X: -4, Y: -2}
	if res := ComputeVectorFromCoordinates(point4a, point4b); res.X != -2 && res.Y != -1 {
		t.Errorf("expected (-2, -1), got %v\n", res)
	}
}

func TestCountVisibleAsteroid(t *testing.T) {
	input0 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	if res := CountVisibleAsteroid(input0, Key{X: 3, Y: 4}); res != 8 {
		t.Errorf("0: expected 8 got %v\n", res)
	}
	input1 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	if res := CountVisibleAsteroid(input1, Key{X: 1, Y: 0}); res != 7 {
		t.Errorf("1: expected 7 got %v\n", res)
	}
	input2 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	if res := CountVisibleAsteroid(input2, Key{X: 0, Y: 2}); res != 6 {
		t.Errorf("2: expected 7 got %v\n", res)
	}
	input3 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	if res := CountVisibleAsteroid(input3, Key{X: 3, Y: 2}); res != 7 {
		t.Errorf("3: expected 7 got %v\n", res)
	}
	input4 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	if res := CountVisibleAsteroid(input4, Key{X: 4, Y: 2}); res != 5 {
		t.Errorf("4: expected 5 got %v\n", res)
	}
}

func TestMapVisible(t *testing.T) {
	input0 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	expected0 := [][]string{
		[]string{".", "7", ".", ".", "7"},
		[]string{".", ".", ".", ".", "."},
		[]string{"6", "7", "7", "7", "5"},
		[]string{".", ".", ".", ".", "7"},
		[]string{".", ".", ".", "8", "7"},
	}
	if res := MapVisible(input0); !CompareII(res, expected0) {
		t.Errorf("expected\n%v\n%v\n", expected0, res)
	}
	input1 := [][]string{
		[]string{".", "#", ".", ".", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "#"},
		[]string{".", ".", ".", "#", "#"},
	}
	expected1 := [][]string{
		[]string{".", "7", ".", ".", "7"},
		[]string{".", ".", ".", ".", "."},
		[]string{"6", "0", "7", "7", "5"},
		[]string{".", ".", ".", ".", "7"},
		[]string{".", ".", ".", "8", "7"},
	}
	if res := MapVisible(input1); CompareII(res, expected1) {
		t.Errorf("expected %v got %v\n", expected1, res)
	}
}

func CompareII(arr1 [][]string, arr2 [][]string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if len(arr1[i]) != len(arr2[i]) {
			return false
		}
		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestFindHigest(t *testing.T) {
	input0 := [][]string{
		[]string{".", ".", ".", ".", ".", ".", "#", ".", "#", "."},
		[]string{"#", ".", ".", "#", ".", "#", ".", ".", ".", "."},
		[]string{".", ".", "#", "#", "#", "#", "#", "#", "#", "."},
		[]string{".", "#", ".", "#", ".", "#", "#", "#", ".", "."},
		[]string{".", "#", ".", ".", "#", ".", ".", ".", ".", "."},
		[]string{".", ".", "#", ".", ".", ".", ".", "#", ".", "#"},
		[]string{"#", ".", ".", "#", ".", ".", ".", ".", "#", "."},
		[]string{".", "#", "#", ".", "#", ".", ".", "#", "#", "#"},
		[]string{"#", "#", ".", ".", ".", "#", ".", ".", "#", "."},
		[]string{".", "#", ".", ".", ".", ".", "#", "#", "#", "#"},
	}
	if res, visible := FindHigest(input0); res.X != 5 && res.Y != 8 && visible != 33 {
		t.Errorf("expected 5,8 with 33 got %v with %v\n", res, visible)
	}
	input1 := [][]string{
		[]string{"#", ".", "#", ".", ".", ".", "#", ".", "#", "."},
		[]string{".", "#", "#", "#", ".", ".", ".", ".", "#", "."},
		[]string{".", "#", ".", ".", ".", ".", "#", ".", ".", "."},
		[]string{"#", "#", ".", "#", ".", "#", ".", "#", ".", "#"},
		[]string{".", ".", ".", ".", "#", ".", "#", ".", "#", "."},
		[]string{".", "#", "#", ".", ".", "#", "#", "#", ".", "#"},
		[]string{".", ".", "#", ".", ".", ".", "#", "#", ".", "."},
		[]string{".", ".", "#", "#", ".", ".", ".", ".", "#", "#"},
		[]string{".", ".", ".", ".", ".", ".", "#", ".", ".", "."},
		[]string{".", "#", "#", "#", "#", ".", "#", "#", "#", "."},
	}
	if res, visible := FindHigest(input1); res.X != 1 && res.Y != 2 && visible != 35 {
		t.Errorf("expected 1,2 with 35 got %v with %v\n", res, visible)
	}
	input2 := [][]string{
		[]string{".", "#", ".", ".", "#", ".", ".", "#", "#", "#"},
		[]string{"#", "#", "#", "#", ".", "#", "#", "#", ".", "#"},
		[]string{".", ".", ".", ".", "#", "#", "#", ".", "#", "."},
		[]string{".", ".", "#", "#", "#", ".", "#", "#", ".", "#"},
		[]string{"#", "#", ".", "#", "#", ".", "#", ".", "#", "."},
		[]string{".", ".", ".", ".", "#", "#", "#", ".", ".", "#"},
		[]string{".", ".", "#", ".", "#", ".", ".", "#", ".", "#"},
		[]string{"#", ".", ".", "#", ".", "#", ".", "#", "#", "#"},
		[]string{".", "#", "#", ".", ".", ".", "#", "#", ".", "#"},
		[]string{".", ".", ".", ".", ".", "#", ".", "#", ".", "."},
	}
	if res, visible := FindHigest(input2); res.X != 6 && res.Y != 3 && visible != 41 {
		t.Errorf("expected 6,3 with 41 got %v with %v\n", res, visible)
	}
	input3 := [][]string{
		[]string{".", "#", ".", ".", "#", "#", ".", "#", "#", "#", ".", ".", ".", "#", "#", "#", "#", "#", "#", "#"},
		[]string{"#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", "."},
		[]string{".", "#", ".", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#"},
		[]string{".", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", ".", "#", "."},
		[]string{"#", "#", "#", "#", "#", ".", "#", "#", ".", "#", ".", "#", "#", ".", "#", "#", "#", ".", "#", "#"},
		[]string{".", ".", "#", "#", "#", "#", "#", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		[]string{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		[]string{"#", ".", "#", "#", "#", "#", ".", ".", ".", ".", "#", "#", "#", ".", "#", ".", "#", ".", "#", "#"},
		[]string{"#", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		[]string{"#", "#", "#", "#", "#", ".", "#", "#", ".", "#", "#", "#", ".", ".", "#", "#", "#", "#", ".", "."},
		[]string{".", ".", "#", "#", "#", "#", "#", "#", ".", ".", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#"},
		[]string{"#", "#", "#", "#", ".", "#", "#", ".", "#", "#", "#", "#", ".", ".", ".", "#", "#", ".", ".", "#"},
		[]string{".", "#", "#", "#", "#", "#", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#"},
		[]string{"#", "#", ".", ".", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", ".", "."},
		[]string{"#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", "#", "#", "#"},
		[]string{".", "#", "#", "#", "#", ".", "#", ".", "#", "#", "#", ".", "#", "#", "#", ".", "#", ".", "#", "#"},
		[]string{".", ".", ".", ".", "#", "#", ".", "#", "#", ".", "#", "#", "#", ".", ".", "#", "#", "#", "#", "#"},
		[]string{".", "#", ".", "#", ".", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", ".", "#", "#", "#"},
		[]string{"#", ".", "#", ".", "#", ".", "#", "#", "#", "#", "#", ".", "#", "#", "#", "#", ".", "#", "#", "#"},
		[]string{"#", "#", "#", ".", "#", "#", ".", "#", "#", "#", "#", ".", "#", "#", ".", "#", ".", ".", "#", "#"},
	}
	if res, visible := FindHigest(input3); res.X != 11 && res.Y != 13 && visible != 210 {
		t.Errorf("expected 11,13 with 210 got %v with %v\n", res, visible)
	}
}

func TestCosTheta(t *testing.T) {
	if CosTheta(Key{X: 0, Y: -1}, Key{X: 1, Y: 0}) != 0.0 {
		t.Error("no")
	}
	if CosTheta(Key{X: 0, Y: -1}, Key{X: -1, Y: 1}) >= 0.0 {
		t.Error("no")
	}
}

func TestOrientedCosTheta(t *testing.T) {
	if OrientedCosTheta(Key{X: 0, Y: -1}, Key{X: 1, Y: 0}) != 0.0 {
		t.Error("no")
	}
	if OrientedCosTheta(Key{X: 0, Y: -1}, Key{X: -1, Y: 1}) >= 0.0 {
		t.Error("no")
	}
	if OrientedCosTheta(Key{X: 0, Y: -1}, Key{X: 1, Y: -1}) >= 0.0 {
		t.Error("no")
	}
	if OrientedCosTheta(Key{X: 0, Y: -1}, Key{X: 1, Y: 1}) == 0.0 {
		t.Error("no")
	}
}

func TestComputeVectorFromCoordinates(t *testing.T) {
	origin := Key{X: 1, Y: 1}
	vect4 := ComputeVectorFromCoordinates(origin, Key{X: 1, Y: 1})
	fmt.Printf("vec %v", vect4)
	if res := math.Pi/2 - math.Atan2(-float64(vect4.Y), float64(vect4.X)); vect4.X != 0 || vect4.Y != 0 {
		t.Errorf("no ; %v\n", res)
	}
	vect1 := ComputeVectorFromCoordinates(origin, Key{X: 1, Y: 0})
	fmt.Printf("vec %v", vect1)
	if res := math.Pi/2 - math.Atan2(-float64(vect1.Y), float64(vect1.X)); vect1.X != 0 || vect1.Y != -1 {
		t.Errorf("no ; %v\n", res)
	}
	vect2 := ComputeVectorFromCoordinates(origin, Key{X: 2, Y: 0})
	fmt.Printf("vec %v", vect2)
	if res := math.Pi/2 - math.Atan2(-float64(vect2.Y), float64(vect2.X)); vect2.X != 1 || vect2.Y != -1 {
		t.Errorf("no ; %v\n", res)
	}
	vect5 := ComputeVectorFromCoordinates(origin, Key{X: 2, Y: 1})
	fmt.Printf("vec %v", vect5)
	if res := math.Pi/2 - math.Atan2(-float64(vect5.Y), float64(vect5.X)); vect5.X != 1 || vect5.Y != 0 {
		t.Errorf("no ; %v\n", res)
	}
	vect8 := ComputeVectorFromCoordinates(origin, Key{X: 2, Y: 2})
	fmt.Printf("vec %v", vect8)
	if res := math.Pi/2 - math.Atan2(-float64(vect8.Y), float64(vect8.X)); vect8.X != 1 || vect8.Y != 1 {
		t.Errorf("no ; %v\n", res)
	}
	vect7 := ComputeVectorFromCoordinates(origin, Key{X: 1, Y: 2})
	fmt.Printf("vec %v", vect7)
	if res := math.Pi/2 - math.Atan2(-float64(vect7.Y), float64(vect7.X)); vect7.X != 0 || vect7.Y != 1 {
		t.Errorf("no ; %v\n", res)
	}
	vect6 := ComputeVectorFromCoordinates(origin, Key{X: 0, Y: 2})
	fmt.Printf("vec %v", vect6)
	if res := math.Pi/2 - math.Atan2(-float64(vect6.Y), float64(vect6.X)); vect6.X != -1 || vect6.Y != 1 {
		t.Errorf("no ; %v\n", res)
	}
	vect3 := ComputeVectorFromCoordinates(origin, Key{X: 0, Y: 1})
	fmt.Printf("vec %v", vect3)
	if res := math.Pi/2 - math.Atan2(-float64(vect3.Y), float64(vect3.X)); vect3.X != -1 || vect3.Y != 0 {
		t.Errorf("no ; %v\n", res)
	}
	vect0 := ComputeVectorFromCoordinates(origin, Key{X: 0, Y: 0})
	fmt.Printf("vec %v", vect0)
	if res := math.Pi/2 - math.Atan2(-float64(vect0.Y), float64(vect0.X)); vect0.X != -1 || vect0.Y != -1 {
		t.Errorf("no ; %v\n", res+2*math.Pi)
	}
}
