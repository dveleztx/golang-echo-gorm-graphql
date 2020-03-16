#!/bin/bash
###############################################################################
# Script      : prep.sh
# Author      : David Velez
# Date        : 03/16/2020
# Description : Setup and configure database
###############################################################################

# Move Database CSV to /tmp
cp $GOPATH/src/github.com/dveleztx/golang-echo-gorm-graphql/data/users_db.csv /tmp
cp $GOPATH/src/github.com/dveleztx/golang-echo-gorm-graphql/scripts/sql/sample_db_create.sql /tmp

# Setup and Configure Database
sudo su - postgres -c "createdb sample_db"
sudo su - postgres -c "psql -U postgres -d sample_db -f /tmp/sample_db_create.sql"

# Clean up
rm -rf /tmp/users_db.csv
rm -rf /tmp/sample_db_create.sql
