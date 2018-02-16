#!/bin/bash

# Restore the dependencies of the function.
if [ -e /openbrisk/$MODULE_NAME.deps.sh ]
then
    # Make deps script executable and run it.
    chmod +x /openbrisk/$MODULE_NAME.deps.sh
    /openbrisk/$MODULE_NAME.deps.sh
fi

# Start the server.
./server