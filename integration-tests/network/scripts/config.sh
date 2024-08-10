#!/bin/bash

set -eu
source scripts/constants.sh

POD_INDEX=${HOSTNAME##*-}
VALIDATOR_INDEX=$((POD_INDEX+1))
VALIDATOR_NAME=val${VALIDATOR_INDEX}

# Redefined to confirm they're set
CHAIN_NAME=${CHAIN_NAME}
CHAIN_HOME=${CHAIN_HOME}
BINARY=${BINARY}
DENOM=${DENOM}
DENOM_DECIMALS=${DENOM_DECIMALS}
NUM_VALIDATORS=${NUM_VALIDATORS}

MICRO_DENOM_ZERO_PAD=$(printf "%${DENOM_DECIMALS}s" | tr ' ' "0")
CHAIN_ID=${CHAIN_NAME}-test-1
BLOCK_TIME=1s
VALIDATOR_BALANCE=10000000${MICRO_DENOM_ZERO_PAD}
VALIDATOR_STAKE=1000000${MICRO_DENOM_ZERO_PAD}
RELAYER_BALANCE=10000000${MICRO_DENOM_ZERO_PAD}
USER_BALANCE=10000000${MICRO_DENOM_ZERO_PAD}

DEPOSIT_PERIOD="30s"
VOTING_PERIOD="30s"
EXPEDITED_VOTING_PERIOD="29s"
UNBONDING_TIME="240s"

STRIDE_DAY_EPOCH_DURATION="140s"
STRIDE_EPOCH_EPOCH_DURATION="35s"