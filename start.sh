#!/usr/bin/env bash


function checkIfEnvVariableSet() {

    local env_var="$2"
    if [[ -z "$env_var" ]]; then
        echo "Error: Env variable $1 is not set"
        exit 1
    else
        echo "Env variable $1 is set"
    fi
}

echo "GO workspace is set to ${GOPATH}"

if [[ $1 == test ]]; then
    echo "Running all tests";
    checkIfEnvVariableSet TEST_DB_USER $TEST_DB_USER
    checkIfEnvVariableSet TEST_DB_PASS $TEST_DB_PASS
    checkIfEnvVariableSet TEST_DB_NAME $TEST_DB_NAME
    checkIfEnvVariableSet TEST_DB_HOST $TEST_DB_HOST
    checkIfEnvVariableSet TEST_DB_PORT $TEST_DB_PORT
    go test github.com/nihlaakram/go-microservice/pkg/test -v

else
    echo "Building executable"
    checkIfEnvVariableSet DB_USER $DB_USER
    checkIfEnvVariableSet DB_PASS $DB_PASS
    checkIfEnvVariableSet DB_NAME $DB_NAME
    checkIfEnvVariableSet DB_HOST $DB_HOST
    checkIfEnvVariableSet DB_PORT $DB_PORT

    rootPath=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
    buildDir="build/server"
    buildPath="$rootPath/${buildDir}"

    echo "Cleaning build path ${buildDir}..."
    rm -rf "build"
    cd cmd/
    go build -o $buildPath
fi


