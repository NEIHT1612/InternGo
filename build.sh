#!/bin/bash
APP_NAME=app
BUILD_DIR=build

#Build app for Linux
GOOS=linux GOARCH=amd64 go build -o ${BUILD_DIR}/${APP_NAME}-linux main.go

#Build app for Windows
GOOS=windows GOARCH=amd64 go build -o ${BUILD_DIR}/${APP_NAME}-win.exe main.go

#Command to run production environment
# $env:APP_ENV="production"; .\my-app-window.exe
# APP_ENV=production ./my-app-linux