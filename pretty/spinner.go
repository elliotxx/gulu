package pretty

import (
	"time"

	"github.com/pterm/pterm"
)

// Pretty spinner style.
//
// Usage:
//   sp, _ := pretty.NewPrettySpinner().Start("Starting ...")
//   time.Sleep(time.Second * 3)
//   sp.Success("Done")
func NewPrettySpinner() *pterm.SpinnerPrinter {
	return &pterm.SpinnerPrinter{
		Sequence:            []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "},
		Style:               &pterm.ThemeDefault.SpinnerStyle,
		Delay:               time.Millisecond * 100,
		ShowTimer:           true,
		TimerRoundingFactor: time.Second,
		TimerStyle:          &pterm.ThemeDefault.TimerStyle,
		MessageStyle:        &pterm.ThemeDefault.InfoMessageStyle,
		SuccessPrinter:      &Success,
		FailPrinter:         &Error,
		WarningPrinter:      &Warning,
	}
}
