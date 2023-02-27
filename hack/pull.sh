#!/usr/bin/env bash

#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.ist>
#     .\_/.
#

clone_and_update_repo() {
  local REPO="$1"
  echo "$(pwd)/$REPO"
  echo ""
  if [ -d "$REPO" ] && [ ! -L "$REPO" ]; then
    echo "$REPO folder already exists. Not going to clone it."
  else
    git clone "https://github.com/zerotohero-dev/$REPO.git"
  fi
  cd "$REPO" || exit
  git stash
  git checkout main
  git pull
  cd ..
  echo ""
}

cd ..
echo "This will stash all your uncommitted changes EVERYWHERE."
read -p "Are you sure? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
  clone_and_update_repo "aegis-spire"
  clone_and_update_repo "aegis-core"
  clone_and_update_repo "aegis-sdk-go"
  clone_and_update_repo "aegis-sentinel"
  clone_and_update_repo "aegis-sidecar"
  clone_and_update_repo "aegis-init-container"
  clone_and_update_repo "aegis-safe"
  clone_and_update_repo "aegis-web"
  clone_and_update_repo "aegis-workload-demo-using-sidecar"
  clone_and_update_repo "aegis-workload-demo-using-sdk"
  clone_and_update_repo "aegis-workload-demo-using-init-container"

  echo "Everything is awesome!"
fi