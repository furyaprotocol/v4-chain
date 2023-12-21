package types_test

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/delaymsg/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModuleAddress(t *testing.T) {
	require.Equal(t, "furya1mkkvp26dngu6n8rmalaxyp3gwkjuzztq5zx6tr", types.ModuleAddress.String())
}
