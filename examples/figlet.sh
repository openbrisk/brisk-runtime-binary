#!/bin/bash

function execute() {
    # https://stackoverflow.com/a/19408949
    # Commands inherit their standard input from the process that starts them. 
    # text/plain based input.
    figlet

    # application/json based input using jq (installed by default).
    #jq --raw-output '.text' | figlet

    # Return forward structure.
    #fig=`figlet`
    #target='https://target.url/hooks'
    #forward "$fig" "$target"
}
