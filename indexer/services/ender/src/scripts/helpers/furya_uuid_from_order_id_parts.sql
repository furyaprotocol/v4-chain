CREATE OR REPLACE FUNCTION furya_uuid_from_order_id_parts(subaccount_id uuid, client_id text, clob_pair_id text, order_flags text) RETURNS uuid AS $$
/**
  Returns a UUID using the parts of an IndexerOrderId (https://github.com/furyanprotocol/v4-chain/blob/9ed26bd/proto/furyaprotocol/indexer/protocol/v1/clob.proto#L15).

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    return furya_uuid(concat(subaccount_id, '-', client_id, '-', clob_pair_id, '-', order_flags));
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;