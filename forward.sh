#!/bin/bash

function forward() {
    jo result=$(echo "$1" | base64 -w 0) forward=$(jo to="$2")
}