cat << EOF
---
title: "gcp"
linkTitle: "technology/gcp"
type: docs
description: >-
     This is a description of the gcp library
---

## Import
EOF

echo '`import "github.com/owulveryck/technology/gcp"`'
echo ""

echo "## Content"

echo "| Technology | Usage |"
echo "|------------|-------|"
cat ../../../../../../technology/gcp/*.cue| while read line
do
  echo $line | grep -q '^[A-Z]'
  if [ $? -eq 0 ]
  then
    current=$(echo $line | sed 's/\(\):.*/\1/')
  fi
  echo $line | grep -q '}$'
  if [ $? -eq 0 ]
  then
    current=""
  fi
  if [ "_$current" != "_" ]
  then
    echo $line | grep -q "name:"
    if [ $? -eq 0 ]
    then
      name=$(echo $line |sed 's/name: \"\(.*\)\"/\1/')
      echo "|$name | \`technology: dev.$current\`|"
    fi
  fi
done
