#! /bin/bash

VARIABLE_TEMP_DIR="/tmp/"`head /dev/urandom | tr -dc A-Za-z0-9 | head -c 13 ; echo ''`
printf 'switch directory\n' >&2
mkdir -p $VARIABLE_TEMP_DIR
pushd $VARIABLE_TEMP_DIR

# -----------------------------------------------------------------------------

OS=`uname`
ARCH=`uname -m`
VARIABLE_GITHUB_RELEASE_URL="https://api.github.com/repos/avrebarra/filateli/releases/latest"

FILENAME_REGEX=""
if [ "$OS" == "Darwin" ]; then
    FILENAME_REGEX="browser_download_url.*Darwin_x86_64\.tar\.gz"
elif [ "$OS" == "Linux" ]; then
  if [ "$ARCH" == "x86_64" ]; then
    FILENAME_REGEX="browser_download_url.*Linux_x86_64\.tar\.gz"
  fi
  if [ "$ARCH" == "i368" ]; then
    FILENAME_REGEX="browser_download_url.*Linux__i386\.tar\.gz"
  fi
else
  printf '%s\n' "os or architecture not supported" >&2
  exit 1
fi

# -----------------------------------------------------------------------------

printf 'downloading binary for %s %s ...\n' $OS $ARCH>&2
curl -s -H "Accept: application/vnd.github.v3+json" $VARIABLE_GITHUB_RELEASE_URL \
| grep $FILENAME_REGEX \
| cut -d ":" -f 2,3 \
| tr -d \" \
| wget -c -i-


# -----------------------------------------------------------------------------

VARIABLE_SOURCE_TARGET_PATH=filateli
VARIABLE_BINARY_TARGET_PATH=/usr/local/bin/filateli

printf 'extracting tarball...\n'>&2
tar -xzf `find . -name "*.tar.gz"`
chmod +x filateli

printf 'registering to system command...\n'>&2
mv $VARIABLE_SOURCE_TARGET_PATH $VARIABLE_BINARY_TARGET_PATH

$VARIABLE_BINARY_TARGET_PATH
