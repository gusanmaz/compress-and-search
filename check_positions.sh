#!/bin/bash

FILE=$1
POSITIONS=($2)
NUM_CHARS=${3:-20}

if [ -z "$FILE" ] || [ -z "$POSITIONS" ]; then
    echo "Usage: $0 <file> <positions> [num_chars]"
    exit 1
fi

for POSITION in "${POSITIONS[@]}"; do
    echo "Position $POSITION:"
    # Calculate the line and character within the line for the byte position
    LINE=$(awk -v pos=$POSITION 'BEGIN {FS=""; OFS=""} {count += length($0) + 1; if (count >= pos) {print NR; exit}}' "$FILE")
    CHAR_POS=$(awk -v pos=$POSITION 'BEGIN {FS=""; OFS=""} {count += length($0) + 1; if (count >= pos) {print pos - (count - length($0) - 1); exit}}' "$FILE")

    # Print the context around the position
    awk -v line=$LINE -v char_pos=$CHAR_POS -v num_chars=$NUM_CHARS 'NR == line {
        start = char_pos > 10 ? char_pos - 10 : 1;
        end = char_pos + num_chars - 10 <= length($0) ? char_pos + num_chars - 10 : length($0);
        print substr($0, start, end - start + 1);
    }' "$FILE"
    echo
done
