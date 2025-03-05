#!/bin/bash

printenv | grep -E 'POSTGRES_' > /etc/environment
cron -f
