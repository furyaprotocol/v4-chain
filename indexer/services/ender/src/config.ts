/**
 * Environment variables required by Ender.
 */

import {
  parseSchema,
  baseConfigSchema,
  parseBoolean,
} from '@furyaprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@furyaprotocol-indexer/kafka';
import {
  postgresConfigSchema,
} from '@furyaprotocol-indexer/postgres';
import { redisConfigSchema } from '@furyaprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...redisConfigSchema,
  ...kafkaConfigSchema,
  SEND_WEBSOCKET_MESSAGES: parseBoolean({
    default: true,
  }),
};

export default parseSchema(configSchema);
