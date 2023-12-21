package app_test

import (
	"github.com/furyanprotocol/v4-chain/protocol/testutil/app"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExportAppStateAndValidators_Panics(t *testing.T) {
	furyaApp := app.DefaultTestApp(nil)
	require.Panics(t, func() { furyaApp.ExportAppStateAndValidators(false, nil, nil) }) // nolint:errcheck
}
