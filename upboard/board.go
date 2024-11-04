//go:build linux

// Package upboard implements an Intel based board.
package upboard

// This is experimental
/*
	Datasheet: https://github.com/up-board/up-community/wiki/Pinout_UP4000
	Supported board: UP4000
*/

import (
	"context"

	"github.com/pkg/errors"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/components/board/genericlinux"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

const modelName = "upboard"

// Model for viam supported upboard.
var Model = resource.NewModel("viam", "up", "upboard")

var logger = logging.NewLogger("test")

func init() {
	gpioMappings, err := genericlinux.GetGPIOBoardMappings(modelName, boardInfoMappings)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		logging.Global().Debugw("error getting up board GPIO board mapping", "error", err)
	}

	resource.RegisterComponent(
		board.API,
		Model,
		resource.Registration[board.Board, *genericlinux.Config]{
			Constructor: func(
				ctx context.Context,
				_ resource.Dependencies,
				conf resource.Config,
				logger logging.Logger,
			) (board.Board, error) {
				return genericlinux.NewBoard(ctx, conf, genericlinux.ConstPinDefs(gpioMappings), logger)
			},
		})
}
