package types_test

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/rewards/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTreasuryModuleAddress(t *testing.T) {
	require.Equal(t, "furya16wrau2x4tsg033xfrrdpae6kxfn9kyuerr5jjp", types.TreasuryModuleAddress.String())
}
