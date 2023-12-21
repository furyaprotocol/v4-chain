FROM furyaprotocol-base

COPY ./testing/containertest/containertest.sh /furyaprotocol/containertest.sh
COPY ./testing/containertest/preupgrade_entrypoint.sh /furyaprotocol/preupgrade_entrypoint.sh
COPY ./testing/containertest/preupgrade_genesis.json /furyaprotocol/preupgrade_genesis.json
COPY ./testing/genesis.sh /furyaprotocol/genesis.sh
COPY ./daemons/pricefeed/client/constants/testdata /furyaprotocol/exchange_config
COPY ./testing/delaymsg_config /furyaprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /furyaprotocol/containertest.sh

COPY ./testing/containertest/config/pricefeed_exchange_config.toml /furyaprotocol/chain/.alice/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /furyaprotocol/chain/.bob/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /furyaprotocol/chain/.carl/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /furyaprotocol/chain/.dave/config/pricefeed_exchange_config.toml

ENTRYPOINT ["furyaprotocold"]
