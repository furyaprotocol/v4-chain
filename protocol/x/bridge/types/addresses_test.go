package types_test

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/bridge/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "furya1zlefkpe3g0vvm9a4h0jf9000lmqutlh9jwjnsv", types.ModuleAddress.String())
}
