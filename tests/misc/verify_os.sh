#!/bin/bash

binary_path=$1
target=$2

output_first_line="$($binary_path 2> /dev/null | head -1)"
expected_first_line="Target: $target"

# Compare first line of output with the target OS
if [ "$output_first_line" = "$expected_first_line" ]; then
    exit 0
else
    echo "Expected '$expected_first_line'"
    echo "Received '$output_first_line'"
    echo "Exit(1)"
    exit 1
fi