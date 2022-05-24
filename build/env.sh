#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"

echo "$root" "$workspace"

phoenixchaindir="$workspace/src/github.com/PhoenixGlobal"
if [ ! -L "$phoenixchaindir/Phoenix-Chain-Core" ]; then
    mkdir -p "$phoenixchaindir"
    cd "$phoenixchaindir"
    ln -s ../../../../../. Phoenix-Chain-Core
    cd "$root"
fi

echo "ln -s success."

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$phoenixchaindir/Phoenix-Chain-Core"
PWD="$phoenixchaindir/Phoenix-Chain-Core"

# Launch the arguments with the configured environment.
exec "$@"
