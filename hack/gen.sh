#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

#cat <<EOF
#Examples:
# Usage: $(basename "$0") <generators> <output-package> <apis-package> <groups-versions> ...
# $(basename "$0") deepcopy,client github.com/example/project/pkg/client github.com/example/project/pkg/apis "foo:v1 bar:v1alpha1,v1beta1"
#EOF
#  exit 0
#
#
#
#GENS="$1"
##generators
##  gen-client
#OUTPUT_PKG="$2"
##output-package
##  github.com/example/project/pkg/client
#APIS_PKG="$3"
##apis-package
##  github.com/example/project/pkg/apis
#GROUPS_WITH_VERSIONS="$4"
##groups-versions
##  foo:v1
#
#
#echo "Generating clientset for ${GROUPS_WITH_VERSIONS} at ${OUTPUT_PKG}/${CLIENTSET_PKG_NAME:-clientset}"

APIS_PKG="github.com/example/project/pkg/apis"
FQ_APIS=()
GROUPS_WITH_VERSIONS=(foo:v1 bar:v1alpha1,v1beta1)
for GVs in ${GROUPS_WITH_VERSIONS}; do
  IFS=: read -r G Vs <<<"${GVs}"
  echo "${G}"
  echo "${Vs}"
  # enumerate versions
  for V in ${Vs//,/ }; do
    FQ_APIS+=("${APIS_PKG}/${G}/${V}")
  done
done

echo "${FQ_APIS}"
FQ_APIS=(github.com/example/project/pkg/apis/foo/v1 github.com/example/project/pkg/apis/bar/v1alpha1)
function codegen::join() {
  local IFS="$1"
  echo "$0"
  echo "$1"
  echo "$2"

#  shift
#  echo "$*"

  echo "$*"
}

echo "$(codegen::join , "${FQ_APIS[@]}")"
