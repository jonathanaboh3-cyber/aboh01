#!/bin/bash
    find . -type f -name "*.sh" -print0 | while IFS= read -r -d $'\0' file; do
        basename "$file" .sh
    done | sort -r