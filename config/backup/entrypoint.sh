#!/bin/bash

printenv | grep -E 'POSTGRES_USER|POSTGRES_PASSWORD|POSTGRES_DB' > /etc/cron.env

cron

tail -f /var/log/cron.log