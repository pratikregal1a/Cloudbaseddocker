#!/bin/bash
file_name="bournvita.png"
MONTH="November" 
echo "Current Month is $MONTH"
echo "Today is $(date)"
echo "Linux version : $(uname -r)"
echo "Memory Information"
free -m
echo "<p><img src=\"$file_name\"><br>$file_name"
