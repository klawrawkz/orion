#!/usr/bin/env bash

URL="$(curl -s https://api.github.com/repos/microsoft/orion/releases/latest | grep browser_download_url | grep linux_amd64 \
| cut -d : -f 2,3 \
| tr -d \")"

echo "Downloading: " $URL
TMP_FILE=$(mktemp)
echo "Saving to file: " "$TMP_FILE"
curl -L $URL -o "$TMP_FILE" #-s

chmod +x "$TMP_FILE"
mv "$TMP_FILE" "/usr/local/bin/orion"
