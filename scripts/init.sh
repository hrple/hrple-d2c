#!/bin/bash

golangci_content=$(golangci-lint --version)
golangci_res=$?
if [ $golangci_res -gt "0" ]
then
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin v1.23.7	    
else
    echo "$golangci_content"
fi

safesql_content=$(safesql)
safesql_res=$?

if [ $safesql_res -gt "0" ]
then
    go get github.com/stripe/safesql
else
    echo "$safesql_content"
fi

git config core.hooksPath .githooks