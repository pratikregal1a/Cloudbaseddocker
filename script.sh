#!/bin/bash
sudo apt-get install fim
file_name="/bournvita.png"
MONTH="November" 
echo "Current Month is $MONTH"
echo "Today is $(date)"
echo "Linux version : $(uname -r)"
echo "Memory Information"
free -m
fim *.png
