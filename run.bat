:: run: .\run.bat

@echo off
go build -o gin-bin.exe main.go
start "" gin-bin.exe
exit /b 0

