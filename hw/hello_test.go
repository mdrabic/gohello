package hello 

import "testing"

func TestHelloArrays(t *testing.T)  {
	count := 3
	result := helloArrays(count)
	if (len(result) != 3) {
		t.Errorf("Expected to find array of len 5")
	}
}


