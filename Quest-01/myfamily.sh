#!/bin/bash

	curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq --arg h_id "$HERO_ID" '.[] | select(.id == ($h_id | tonumber)) | .connections.relatives' | tr -d "\""