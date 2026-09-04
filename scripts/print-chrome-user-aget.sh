#!/bin/bash
SCRIPT_ID="My-Script"
/opt/google/chrome/chrome --headless=new --remote-debugging-port=9222 --user-data-dir="$HOME/.config/google-chrome/${SCRIPT_ID}-TMP/" &>/dev/null &
PID=$!
sleep 0.35
curl --retry-connrefused --retry 10 --retry-delay 1 http://localhost:9222/json/version --silent | jq -r '.["User-Agent"]' | sed s/Headless//
kill -SIGINT $PID
