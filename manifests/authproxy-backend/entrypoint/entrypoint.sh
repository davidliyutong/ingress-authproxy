#!/bin/bash
CONFIG_DIR=/config
if [ ! -f $CONFIG_DIR/config.yaml ]; then
    echo "creating config.yaml"
    ./authproxy init --print > "$CONFIG_DIR/config.yaml"
    cat "$CONFIG_DIR/config.yaml"
fi

./authproxy serve --config "$CONFIG_DIR/config.yaml"
