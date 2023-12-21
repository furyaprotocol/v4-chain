import {
  ComplianceClient,
  getComplianceClient,
} from '@furyaprotocol-indexer/compliance';
import { ComplianceProvider } from '@furyaprotocol-indexer/postgres';

export interface ClientAndProvider {
  client: ComplianceClient;
  provider: ComplianceProvider;
}

export const complianceProvider: ClientAndProvider = {
  client: getComplianceClient(),
  provider: ComplianceProvider.ELLIPTIC,
};
