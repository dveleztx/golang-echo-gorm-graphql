# Golang Stack 

## Description

This project creates a SQL API using GraphQL that returns data in JSON format to the web browser via Echo. Queries can be made with GORM to retrieve data for the front-end. GraphQL is for retrieving data via URLs (like a REST API). Unlike REST, you can send *one* query for *multiple* tables, limiting the GET request to just one to retrieve data. I use the Echo Web Framework and Realize to run my web server. Realize keeps track of the logs of the Echo Web Framework.

## Getting Started

For the data, use the CSV under `data/users_db.csv`. The data could be better, feel free to change it. I opted to use PostgreSQL, the code will reflect that I'm using libraries in order to read a PSQL database. To setup your database is simple.

For initial setup (in Debian/Ubuntu):

```
sudo apt update && sudo apt upgrade -y
sudo apt install -y postgresql-11
sudo su - postgres -c psql
```

**NOTE**: Before proceding to the next step, move your db csv file to the `/tmp` directory

In PSQL Shell:

```
CREATE DATABASE sample_db;
\c sample_db
CREATE TABLE users;
COPY users(id,name,age,created_at,updated_at,deleted_at) FROM '/tmp/users_db.csv' DELIMITER ',' CSV HEADER;
```

Once created, you should be able to use my program.


## References

Link: https://medium.com/@manakuro/go-api-with-echo-gorm-and-graphql-1565ad921626 
