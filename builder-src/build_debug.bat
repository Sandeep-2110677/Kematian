@echo off

cd /d %~dp0

go build .

timeout /t 1 /nobreak >nul

builder.exe

echo %errorlevel%

pause

del builder.exe
exit