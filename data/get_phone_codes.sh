#!/bin/sh
curl 'https://raw.githubusercontent.com/google/libphonenumber/master/resources/PhoneNumberMetadata.xml' | grep '<territory' | sed -e 's/.*id="//;s/".*countryCode="/,/;s/".*//' >phone_codes.csv
