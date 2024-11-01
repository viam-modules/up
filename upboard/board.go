// Package upboard implements an Intel based board.
package upboard

// This is experimental
/*
	Datasheet: https://github.com/up-board/up-community/wiki/Pinout_UP4000
	Supported board: UP4000
*/

import (
	"github.com/pkg/errors"
	"periph.io/x/host/v3"

	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "upboard"

// Model for viam supported upboard.
var Model = resource.NewModel("viam", "up", "upboard")

func init() {
	if _, err := host.Init(); err != nil {
		logging.Global().Debugw("error initializing host", "error", err)
	}

	gpioMappings, err := genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		logging.Global().Debugw("error getting up board GPIO board mapping", "error", err)
	}

	genericlinux.RegisterBoard(modelName, gpioMappings)
}
