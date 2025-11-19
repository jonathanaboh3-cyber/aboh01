#!/bin/bash
awk -F',' -v id="170" '$1 == id {print $2; print $3; print $4}'