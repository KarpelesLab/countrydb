#!/bin/sh

curl https://meetings.icann.org/en/regions | sed -e ':a;N;$!ba;s/\n[[:space:]]*<td/<td/g' | grep 'bgcolor="#FFFFFF"' | sed -e 's#</td>#@#g;s/<[^>]*>//g;s/^[^@]*\@//;s/@$//;s/@/,/;s/ .*//' >icann_regions.csv
