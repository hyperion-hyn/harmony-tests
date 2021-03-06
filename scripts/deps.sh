#!/usr/bin/env bash
declare -A repos

repos["harmony-one/harmony-tf"]="master"

echo "Updating dependencies..."

for repo in "${!repos[@]}"; do
  branch=${repos[$repo]}
  commit=$(curl -Ls https://api.github.com/repos/${repo}/commits/${branch} | jq '.sha' | tr -d '"')
  echo "Updating dependency ${repo} to latest ${branch} commit ${commit}"
  go get -u github.com/${repo}@${commit}
done

echo "Tidying up go.mod ..."
go mod tidy
