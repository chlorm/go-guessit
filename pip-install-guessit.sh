#!/usr/bin/env sh

# WARNING: this is does not securely download get-pip.py

set -e

if ! type pip >/dev/null; then
  curl -O 'https://bootstrap.pypa.io/get-pip.py'
  python 'get-pip.py' '--user'
  PATH="$PATH:$HOME/.local/bin"
fi

pip 'install' '--user' 'guessit'

echo "Add \$HOME/.local/bin to your PATH"

exit 0
