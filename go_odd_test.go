package is_odd

import "testing"

func TestIsOdd(t *testing.T) {
	r, _ := IsOdd(0)
	expect := false
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(1)
	expect = true
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(2)
	expect = false
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(3)
	expect = true
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(-1)
	expect = true
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(-2)
	expect = false
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(-3)
	expect = true
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}
	r, _ = IsOdd("-3")
	expect = true
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	// _, err := IsOdd("hello")
	// if err != nil {
	// 	t.Errorf("%v", err)
	// }

	r, _ = IsOdd(1.0e0)
	expect = false
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	r, _ = IsOdd(900719925474099200)
	expect = false
	if expect != r {
		t.Errorf("result should have %v "+
			";but had %v", expect, r)
	}

	// _, err := IsOdd("www")
	// if err != nil {
	// 	t.Errorf("err value is:%s", err.Error())
	// }

	// _, err = IsOdd(-1.1)
	// if err != nil {
	// 	t.Errorf("err value is:%s", err.Error())
	// }

	// _, err = IsOdd("-1.1")
	// if err != nil {
	// 	t.Errorf("err value is:%s", err.Error())
	// }
}
