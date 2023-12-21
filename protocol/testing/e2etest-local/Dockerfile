FROM furyaprotocol-base

COPY ./testing/e2etest-local/local.sh /furyaprotocol/local.sh
COPY ./testing/genesis.sh /furyaprotocol/genesis.sh
COPY ./testing/start.sh /furyaprotocol/start.sh
COPY ./daemons/pricefeed/client/constants/testdata /furyaprotocol/exchange_config
COPY ./testing/delaymsg_config /furyaprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.4.0

RUN /furyaprotocol/local.sh

ENV DAEMON_NAME=furyaprotocold
ENTRYPOINT ["cosmovisor", "run"]
