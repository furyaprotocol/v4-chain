import { readFileSync } from 'fs';
import path from 'path';

import { logger } from '@furyaprotocol-indexer/base';
import { dbHelpers, storeHelpers } from '@furyaprotocol-indexer/postgres';

export type PostgresFunction = {
  // The name of the script
  readonly name: string;
  // The contents of the script
  readonly script: string;
};

/**
 * Loads a named script from the specified path.
 *
 * @param name The name of the script.
 * @param scriptPath The path to the script.
 * @returns The created script object
 */
function newScript(name: string, scriptPath: string): PostgresFunction {
  const script: string = readFileSync(path.resolve(__dirname, scriptPath)).toString();
  return {
    name,
    script,
  };
}

const HANDLER_SCRIPTS: string[] = [
  'furya_asset_create_handler.sql',
  'furya_block_processor_ordered_handlers.sql',
  'furya_block_processor_unordered_handlers.sql',
  'furya_deleveraging_handler.sql',
  'furya_funding_handler.sql',
  'furya_liquidity_tier_handler.sql',
  'furya_market_create_handler.sql',
  'furya_market_modify_handler.sql',
  'furya_market_price_update_handler.sql',
  'furya_perpetual_market_handler.sql',
  'furya_stateful_order_handler.sql',
  'furya_subaccount_update_handler.sql',
  'furya_transfer_handler.sql',
  'furya_update_clob_pair_handler.sql',
  'furya_update_perpetual_handler.sql',
];

const DB_SETUP_SCRIPTS: string[] = [
  'create_extension_pg_stat_statements.sql',
  'create_extension_uuid_ossp.sql',
];

const HELPER_SCRIPTS: string[] = [
  'furya_clob_pair_status_to_market_status.sql',
  'furya_create_initial_rows_for_tendermint_block.sql',
  'furya_create_tendermint_event.sql',
  'furya_create_transaction.sql',
  'furya_event_id_from_parts.sql',
  'furya_from_jsonlib_long.sql',
  'furya_from_protocol_order_side.sql',
  'furya_from_protocol_time_in_force.sql',
  'furya_from_serializable_int.sql',
  'furya_get_fee_from_liquidity.sql',
  'furya_get_order_status.sql',
  'furya_get_perpetual_market_for_clob_pair.sql',
  'furya_get_total_filled_from_liquidity.sql',
  'furya_get_weighted_average.sql',
  'furya_liquidation_fill_handler_per_order.sql',
  'furya_order_fill_handler_per_order.sql',
  'furya_perpetual_position_and_order_side_matching.sql',
  'furya_protocol_condition_type_to_order_type.sql',
  'furya_tendermint_event_to_transaction_index.sql',
  'furya_trim_scale.sql',
  'furya_update_perpetual_position_aggregate_fields.sql',
  'furya_uuid.sql',
  'furya_uuid_from_asset_position_parts.sql',
  'furya_uuid_from_fill_event_parts.sql',
  'furya_uuid_from_funding_index_update_parts.sql',
  'furya_uuid_from_oracle_price_parts.sql',
  'furya_uuid_from_order_id.sql',
  'furya_uuid_from_order_id_parts.sql',
  'furya_uuid_from_perpetual_position_parts.sql',
  'furya_uuid_from_subaccount_id.sql',
  'furya_uuid_from_subaccount_id_parts.sql',
  'furya_uuid_from_transaction_parts.sql',
  'furya_uuid_from_transfer_parts.sql',
];

const MAIN_SCRIPTS: string[] = [
  'furya_block_processor.sql',
];

const SCRIPTS: string[] = [
  ...HANDLER_SCRIPTS.map((script: string) => `handlers/${script}`),
  ...HELPER_SCRIPTS.map((script: string) => `helpers/${script}`),
  ...DB_SETUP_SCRIPTS.map((script: string) => `setup/${script}`),
  ...MAIN_SCRIPTS,
];

export async function createPostgresFunctions(): Promise<void> {
  await Promise.all([
    dbHelpers.createModelToJsonFunctions(),
    ...SCRIPTS.map((script: string) => storeHelpers.rawQuery(newScript(script, `../../scripts/${script}`).script, {})
      .catch((error: Error) => {
        logger.error({
          at: 'postgres-functions#createPostgresFunctions',
          message: `Failed to create or replace function contained in ${script}`,
          error,
        });
        throw error;
      }),
    ),
  ]);
}
