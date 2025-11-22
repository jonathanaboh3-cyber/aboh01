#!/bin/bash

INTERVIEW_NUM=$(find . -path "**/interviews/interview_*" -type f | grep -oE '[0-9]+')
echo "$INTERVIEW_NUM"
cat "$(find . -path "**/interviews/interview_$INTERVIEW_NUM" -type f)"
echo "$MAIN_SUSPECT"
