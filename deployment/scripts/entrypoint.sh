#!/bin/sh

set -e
set -x

export SERVICE_NAME="$1"

exec ./"$SERVICE_NAME"