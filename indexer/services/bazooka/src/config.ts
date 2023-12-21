/**
 * Environment variables required by Bazooka.
 */

import {
  parseSchema,
  baseConfigSchema,
  parseInteger,
  parseBoolean,
} from '@furyaprotocol-indexer/base';
import {
  kafkaConfigSchema,
} from '@furyaprotocol-indexer/kafka';
import {
  postgresConfigSchema,
} from '@furyaprotocol-indexer/postgres';
import {
  redisConfigSchema,
} from '@furyaprotocol-indexer/redis';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
  ...redisConfigSchema,
  // Parameters for exponential backoff when retrying clearing Kafka.
  // Max wait time between retries of 3*(2**2) = 12 seconds.
  CLEAR_KAFKA_TOPIC_RETRY_MS: parseInteger({ default: 3000 }),
  CLEAR_KAFKA_TOPIC_MAX_RETRIES: parseInteger({ default: 3 }),
  PREVENT_BREAKING_CHANGES_WITHOUT_FORCE: parseBoolean({
    default: true,
  }),
};

export default parseSchema(configSchema);
