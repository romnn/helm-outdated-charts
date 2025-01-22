#!/bin/sh -e

# Copied w/ love from the excellent hypnoglow/helm-s3

if [ -n "${HELM_OUTDATED_DEPENDENCIES_PLUGIN_NO_INSTALL_HOOK}" ]; then
    echo "Development mode: not downloading versioned release."
    exit 0
fi

VERSION="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Downloading and installing helm-outdated v${VERSION} ..."

REPO="https://github.com/romnn/helm-outdated-charts"
BIN=helm-outdated-charts

URL=""
if [ "$(uname)" = "Darwin" ]; then
    if [ "$(uname -m)" = "x86_64" ]; then
        URL="${REPO}/releases/download/v${VERSION}/${BIN}_${VERSION}_darwin_amd64.tar.gz"
    else
        URL="${REPO}/releases/download/v${VERSION}/${BIN}_${VERSION}_darwin_arm64.tar.gz"
    fi
elif [ "$(uname)" = "Linux" ] ; then
    if [ "$(uname -m)" = "x86_64" ]; then
        URL="${REPO}/releases/download/v${VERSION}/${BIN}_${VERSION}_linux_amd64.tar.gz"
    else
        URL="${REPO}/releases/download/v${VERSION}/${BIN}_${VERSION}_linux_arm64.tar.gz"
    fi
else
    echo "${BIN} does not support windows"
    exit 1
fi

echo $URL

mkdir -p "bin"
mkdir -p "releases/v${VERSION}"

# Download with curl if possible.
if [ -x "$(which curl 2>/dev/null)" ]; then
    curl -sSL "${URL}" -o "releases/v${VERSION}.tar.gz"
else
    wget -q "${URL}" -O "releases/v${VERSION}.tar.gz"
fi
tar xzf "releases/v${VERSION}.tar.gz" -C "releases/v${VERSION}"
mv "releases/v${VERSION}/bin/helm-outdated" "bin/helm-outdated"
