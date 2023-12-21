package types_test

import (
	"github.com/furyanprotocol/v4-chain/protocol/x/clob/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsuranceFundModuleAddress(t *testing.T) {
	require.Equal(t, "furya1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq", types.InsuranceFundModuleAddress.String())
}
