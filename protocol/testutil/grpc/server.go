package grpc

import pricetypes "github.com/furyanprotocol/v4-chain/protocol/x/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
