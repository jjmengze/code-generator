APIS=(apis/foo/v1 apis/bar/v1alpha1)
function join() {
 #: local IFS="$1"

  echo "$1"
  echo "$2"
  echo "========"

  echo "$#"
  echo "========"
  echo "$*"
}

echo "$(join , "${APIS[@]}")"
