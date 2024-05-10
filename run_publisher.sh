#!/bin/bash

source .env
# FINALIZE_SNAPSHOT_URL="http://localhost:3000" TARGET_CHAIN_URL="http://localhost:3000"
go run ./cmd/pull-update-publisher
