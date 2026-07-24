#!/usr/bin/env bash
set -e

if adb devices | grep -q "emulator-.*device"; then # skip if an emulator is already running
  echo "emulator already running"
  exit 0
fi
emulator -avd "$ANDROID_EMULATOR_DEVICE" -no-snapshot-load >/dev/null 2>&1 &
adb wait-for-device

until [ "$(adb shell getprop sys.boot_completed 2>/dev/null | tr -d '\r')" = "1" ]; do # wait for full boot, not just device visibility
  sleep 2
done
echo "emulator booted"