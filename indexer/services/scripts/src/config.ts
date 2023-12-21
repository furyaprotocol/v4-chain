import { baseConfigSchema, parseSchema } from '@furyaprotocol-indexer/base';
import { kafkaConfigSchema } from '@furyaprotocol-indexer/kafka';
import { postgresConfigSchema } from '@furyaprotocol-indexer/postgres';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
};

export default parseSchema(configSchema);
