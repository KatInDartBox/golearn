package math

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Print("main test run!")
	os.Exit(m.Run())
}
