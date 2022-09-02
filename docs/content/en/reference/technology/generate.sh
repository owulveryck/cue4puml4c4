egrep '^#' ../../../../../technology/*.cue | sed 's/.*#\([^:]*\):.*/\1/' | while read line
do
cat > $line.md <<EOF
---
title: "#$line"
linkTitle: "#$line"
type: docs
description: >-
     This is a description of the #$line object
---

## Definition

\`\`\`cue
$(cue def ../../../../../technology -e "#$line")
\`\`\`

## Usage
$(cue def ../../../../../technology -e "#$line" | egrep '^\s|^=' | sed 's/^	\([a-zA-Z#]\)/- \1/g' | gsed 's/#\([A-Za-z]*\)/\[#\1\](..\/\L\1)/')

EOF
done
