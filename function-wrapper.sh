#!/bin/bash

# A wrapper script that starts a new sub-shell to catch
# any error caused by usage of exit() in the child function
# script.

echo $MODULE_NAME
echo $FUNCTION_HANDLER
# TODO: Forward env to subshell

(
    source ./examples/$MODULE_NAME.sh

    # Check if the function exists (bash specific).
    if declare -f "$1" > /dev/null
    then
        # Call arguments verbatim.
        "$@"
    else
        # Show an error message.
        echo "The function handler '$1' is not known."
    fi
)

if [ ! "$?" = 0 ]; then
    echo "The function exited with an error."
fi
