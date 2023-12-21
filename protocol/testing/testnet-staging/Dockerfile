FROM furyaprotocol-base

COPY ./testing/testnet-staging/staging.sh /furyaprotocol/staging.sh
COPY ./testing/genesis.sh /furyaprotocol/genesis.sh
COPY ./testing/start.sh /furyaprotocol/start.sh
COPY ./daemons/pricefeed/client/constants/testdata /furyaprotocol/exchange_config
COPY ./testing/delaymsg_config /furyaprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /furyaprotocol/staging.sh

ENV DAEMON_NAME=furyaprotocold
ENTRYPOINT ["cosmovisor", "run"]
