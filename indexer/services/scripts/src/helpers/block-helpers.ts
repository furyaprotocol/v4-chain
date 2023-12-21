import { logger } from '@furyaprotocol-indexer/base';
import {
  AssetCreateEventV1,
  FundingEventV1,
  IndexerTendermintEvent,
  LiquidityTierUpsertEventV1,
  MarketEventV1,
  OrderFillEventV1,
  PerpetualMarketCreateEventV1,
  StatefulOrderEventV1,
  SubaccountUpdateEventV1,
  TransferEventV1,
  UpdateClobPairEventV1,
  UpdatePerpetualEventV1,
} from '@furyaprotocol-indexer/v4-protos';

import { AnnotatedIndexerTendermintEvent, FuryaIndexerSubtypes } from './types';

export function annotateIndexerTendermintEvent(
  event: IndexerTendermintEvent,
): AnnotatedIndexerTendermintEvent | undefined {
  const eventDataBinary: Uint8Array = event.dataBytes;
  switch (event.subtype) {
    case (FuryaIndexerSubtypes.ORDER_FILL.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(OrderFillEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.SUBACCOUNT_UPDATE.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(SubaccountUpdateEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.TRANSFER.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(TransferEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.MARKET.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(MarketEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.STATEFUL_ORDER.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(StatefulOrderEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.FUNDING.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(FundingEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.ASSET.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(AssetCreateEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.PERPETUAL_MARKET.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(PerpetualMarketCreateEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.LIQUIDITY_TIER.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(LiquidityTierUpsertEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.UPDATE_PERPETUAL.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(UpdatePerpetualEventV1.decode(eventDataBinary)),
      };
    }
    case (FuryaIndexerSubtypes.UPDATE_CLOB_PAIR.toString()): {
      return {
        ...event,
        dataBytes: new Uint8Array(),
        data: JSON.stringify(UpdateClobPairEventV1.decode(eventDataBinary)),
      };
    }
    default: {
      const message: string = `Unable to parse event subtype: ${event.subtype}`;
      logger.error({
        at: 'block-helpers#annotateIndexerTendermintEvent',
        message,
      });
      return undefined;
    }
  }
}
