package reverse_string

import (
	"fmt"
	"testing"

	"github.com/kyokomi/emoji"
)

func TestReverseString(t *testing.T) {

	t.Run("One", func(t *testing.T) {

		var result = emoji.Sprint("Hello :world_map:!")
		var input = ReverseString(result)

		realResult := ReverseString(input)

		if realResult != result {
			t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", result, realResult)
		}

	})

	t.Run("Two", func(t *testing.T) {

		var input = "Señor"
		var result = "roñeS"

		realResult := ReverseString(input)

		if realResult != result {
			t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", result, realResult)
		}
	})

	t.Run("Three", func(t *testing.T) {

		var input = `Test
test`
		var result = `tset
tseT`

		// bytes1 := []byte(input)
		// fmt.Println(bytes1)
		// bytes2 := []byte(result)
		// fmt.Println(bytes2)
		// fmt.Println(string(bytes2))

		realResult := ReverseString(input)

		fmt.Println(realResult)

		if realResult != result {
			t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", result, realResult)
		}
	})

}
