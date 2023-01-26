#!/usr/bin/env bash

OLD=aegis-sentinel-58f6478b79-6g242
NEW=aegis-sentinel-58f6478b79-6g242

sed -i "s/$OLD/$NEW/" ./*.sh
