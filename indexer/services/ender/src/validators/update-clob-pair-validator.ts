import { perpetualMarketRefresher } from '@furyaprotocol-indexer/postgres';
import { IndexerTendermintEvent, UpdateClobPairEventV1 } from '@furyaprotocol-indexer/v4-protos';

import { Handler } from '../handlers/handler';
import { UpdateClobPairHandler } from '../handlers/update-clob-pair-handler';
import { Validator } from './validator';

export class UpdateClobPairValidator extends Validator<UpdateClobPairEventV1> {
  public validate(): void {
    if (perpetualMarketRefresher.getPerpetualMarketFromClobPairId(
      this.event.clobPairId.toString(),
    ) === undefined) {
      return this.logAndThrowParseMessageError(
        'UpdateClobPairEvent.clobPairId must correspond with an existing perpetual_market.clobPairId',
        { event: this.event },
      );
    }
  }

  public createHandlers(
    indexerTendermintEvent: IndexerTendermintEvent,
    txId: number,
  ): Handler<UpdateClobPairEventV1>[] {
    const handler: Handler<UpdateClobPairEventV1> = new UpdateClobPairHandler(
      this.block,
      this.blockEventIndex,
      indexerTendermintEvent,
      txId,
      this.event,
    );

    return [handler];
  }
}
