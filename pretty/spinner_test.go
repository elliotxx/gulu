package pretty

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSpinnerPrinterSuccess(t *testing.T) {
	fmt.Println("SpinnerPrinter:")
	defer fmt.Println("")

	sp, err := NewPrettySpinner().Start("Starting ...")
	assert.Nil(t, err)
	time.Sleep(time.Second * 1)
	sp.Success("Success")

	sp, err = NewPrettySpinner().Start("Starting ...")
	assert.Nil(t, err)
	time.Sleep(time.Second * 1)
	sp.Fail("Fail")

	sp, err = NewPrettySpinner().Start("Starting ...")
	assert.Nil(t, err)
	time.Sleep(time.Second * 1)
	sp.Warning("Warning")
}
