#!/bin/bash

KEY_INTERVIEW=$(grep -o '[0-9]\+' path/to/the/file/that/reveals/the/key_interview)
export KEY_INTERVIEW

echo "$KEY_INTERVIEW"

cat path/to/the/file/that/reveals/the/key_interview

echo "$MAIN_SUSPECT"
