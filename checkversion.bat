@echo off
:loop
curl 192.168.99.100:32000/version
echo.
sleep .5
goto loop