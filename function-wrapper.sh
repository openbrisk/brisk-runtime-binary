#!/bin/bash

# A wrapper script that starts a new sub-shell to catch
# any error caused by usage of exit() in the child function
# script.

(
    source /openbrisk/$MODULE_NAME.sh

    # Check if the function exists (bash specific).
    if declare -f "$FUNCTION_HANDLER" > /dev/null
    then
        # Call function.
        "$FUNCTION_HANDLER"
    else
        # Show an error message.
        echo "The function handler '$FUNCTION_HANDLER' is not known."
    fi
)

if [ ! "$?" = 0 ]
then
    echo "The function exited with an error."
fi
