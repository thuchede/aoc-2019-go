package main

import (
	"testing"
)

func TestComputeLength(t *testing.T) {
	com := make(map[string][]string)
	com["COM"] = []string{"ABC"}
	if res := ComputeLength(com); res != 1 {
		t.Errorf("expected 1 got %v", res)
	}
	com1 := make(map[string][]string)
	com1["COM"] = []string{"ABC"}
	com1["COM"] = append(com1["COM"], "ABB")
	if res := ComputeLength(com1); res != 2 {
		t.Errorf("expected 2 got %v", res)
	}
	com2 := make(map[string][]string)
	com2["COM"] = []string{"ABC"}
	com2["ABC"] = []string{"CDD"}
	com2["ABC"] = append(com2["ABC"], "CDE")
	if res := ComputeLength(com2); res != 5 {
		t.Errorf("expected 5 got %v", res)
	}
	com3 := make(map[string][]string)
	com3["COM"] = []string{"ABC"}
	com3["ABC"] = []string{"CDD"}
	com3["ABC"] = append(com3["ABC"], "CDE")
	com3["ABC"] = append(com3["ABC"], "CDF")
	if res := ComputeLength(com3); res != 7 {
		t.Errorf("expected 7 got %v", res)
	}
	com4 := make(map[string][]string)
	com4["COM"] = []string{"ABC"}
	com4["COM"] = append(com4["COM"], "ABB")
	com4["ABC"] = []string{"CDD"}
	com4["ABC"] = append(com4["ABC"], "CDE")
	com4["ABC"] = append(com4["ABC"], "CDF")
	if res := ComputeLength(com4); res != 8 {
		t.Errorf("expected 58got %v", res)
	}
	com5 := make(map[string][]string)
	com5["COM"] = []string{"ABC"}
	com5["COM"] = append(com5["COM"], "ABB")
	com5["ABC"] = []string{"CDD"}
	com5["ABC"] = append(com5["ABC"], "CDE")
	com5["ABC"] = append(com5["ABC"], "CDF")
	com5["CDE"] = []string{"EFG"}
	com5["CDE"] = append(com5["CDE"], "EFH")
	if res := ComputeLength(com5); res != 14 {
		t.Errorf("expected 14 got %v", res)
	}
}

func TestComputeLengthToAnotherOrbit(t *testing.T) {
	com := make(map[string]string)
	com["ABC"] = "COM"
	com["DEF"] = "COM"
	if res := ComputeLengthToAnotherOrbit(com, "ABC", "DEF"); res != 0 {
		t.Errorf("expected 0 got %v", res)
	}
	com1 := make(map[string]string)
	com1["ABC"] = "COM"
	com1["DEF"] = "ABC"
	com1["HIJ"] = "DEF"
	com1["KLM"] = "ABC"
	if res := ComputeLengthToAnotherOrbit(com1, "KLM", "HIJ"); res != 1 {
		t.Errorf("expected 1 got %v", res)
	}
	com2 := make(map[string]string)
	com2["ABC"] = "COM"
	com2["DEF"] = "ABC"
	com2["HIJ"] = "DEF"
	com2["KLM"] = "ABC"
	com2["NOP"] = "KLM"
	com2["TUV"] = "NOP"
	com2["QRS"] = "KLM"
	if res := ComputeLengthToAnotherOrbit(com2, "TUV", "HIJ"); res != 3 {
		t.Errorf("expected 3 got %v", res)
	}
}

// COM - ABC - DEF - HIJ
// 					- KLM - NOP - TUV
// 								- QRS