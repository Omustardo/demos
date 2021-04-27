package solved

import (
	"fmt"
	"testing"
)

// run this to generate test cases
func genTests() {
	answers := map[int]int64{
		4:  906609,
		5:  232792560,
		6:  25164150,
		7:  104743,
		8:  23514624000,
		9:  31875000,
		10: 142913828922,
		11: 70600674,
		12: 76576500,
		13: 5537376230,
		14: 837799,
		15: 137846528820,
		16: 1366,
		17: 21124,
		18: 1074,
		19: 171,
		20: 648,
		21: 31626,
	}

	for i := 4; i < len(answers)+4; i++ {
		fmt.Printf(`func Test%d(t *testing.T) {
	if n := Problem%d(); n != %d {
		t.Errorf("#%d got %%d, expected %%d", n, %d)
	}
}
`, i, i, answers[i], i, answers[i])
	}
}

func Test4(t *testing.T) {
	if n := Problem4(); n != 906609 {
		t.Errorf("#4 got %d, expected %d", n, 906609)
	}
}
func Test5(t *testing.T) {
	if n := Problem5(); n != 232792560 {
		t.Errorf("#5 got %d, expected %d", n, 232792560)
	}
}
func Test6(t *testing.T) {
	if n := Problem6(); n != 25164150 {
		t.Errorf("#6 got %d, expected %d", n, 25164150)
	}
}
func Test7(t *testing.T) {
	if n := Problem7(); n != 104743 {
		t.Errorf("#7 got %d, expected %d", n, 104743)
	}
}
func Test8(t *testing.T) {
	if n := Problem8(); n != 23514624000 {
		t.Errorf("#8 got %d, expected %d", n, 23514624000)
	}
}
func Test9(t *testing.T) {
	if n := Problem9(); n != 31875000 {
		t.Errorf("#9 got %d, expected %d", n, 31875000)
	}
}
func Test10(t *testing.T) {
	if n := Problem10(); n != 142913828922 {
		t.Errorf("#10 got %d, expected %d", n, 142913828922)
	}
}
func Test11(t *testing.T) {
	if n := Problem11(); n != 70600674 {
		t.Errorf("#11 got %d, expected %d", n, 70600674)
	}
}
func Test12(t *testing.T) {
	if n := Problem12(); n != 76576500 {
		t.Errorf("#12 got %d, expected %d", n, 76576500)
	}
}
func Test13(t *testing.T) {
	if n := Problem13(); n != 5537376230 {
		t.Errorf("#13 got %d, expected %d", n, 5537376230)
	}
}
func Test14(t *testing.T) {
	if n := Problem14(); n != 837799 {
		t.Errorf("#14 got %d, expected %d", n, 837799)
	}
}
func Test15(t *testing.T) {
	if n := Problem15(); n != 137846528820 {
		t.Errorf("#15 got %d, expected %d", n, 137846528820)
	}
}
func Test16(t *testing.T) {
	if n := Problem16(); n != 1366 {
		t.Errorf("#16 got %d, expected %d", n, 1366)
	}
}
func Test17(t *testing.T) {
	if n := Problem17(); n != 21124 {
		t.Errorf("#17 got %d, expected %d", n, 21124)
	}
}
func Test18(t *testing.T) {
	if n := Problem18(); n != 1074 {
		t.Errorf("#18 got %d, expected %d", n, 1074)
	}
}
func Test19(t *testing.T) {
	if n := Problem19(); n != 171 {
		t.Errorf("#19 got %d, expected %d", n, 171)
	}
}
func Test20(t *testing.T) {
	if n := Problem20(); n != 648 {
		t.Errorf("#20 got %d, expected %d", n, 648)
	}
}
func Test21(t *testing.T) {
	if n := Problem21(); n != 31626 {
		t.Errorf("#21 got %d, expected %d", n, 31626)
	}
}
