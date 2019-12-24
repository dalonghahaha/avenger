package tests

import (
	"fmt"
	"testing"
)

func testingReport(t *testing.T, result interface{}, err error) {
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(result)
	}
}
