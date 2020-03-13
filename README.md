# Golang Stack 

## Description

This project creates a SQL API using GraphQL that returns data in JSON format to the web browser via Echo. Queries can be made with GORM to retrieve data for the front-end. GraphQL is for retrieving data via URLs (like a REST API). Unlike REST, you can send *one* query for *multiple* tables, limiting the GET request to just one to retrieve data. I use the Echo Web Framework and Realize to run my web server. Realize keeps track of the logs of the Echo Web Framework.

## Getting Started

For the data, use the CSV under `data/users_db.csv`. The data could be better, feel free to change it. I opted to use PostgreSQL, the code will reflect that I'm using libraries in order to read a PSQL database. To setup your database is simple.

### Install packages
*For initial setup (in Debian/Ubuntu only)*

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install -y git vim golang postgresql-11 
```

### Setup and Configure Go

First, create your `go` directory first:

```
mkdir -p $HOME/go/src
```

Next, you need to set environment variables for go, for $GOPATH and $GOROOT inside `.bashrc`. It should have the following in it:

```bash
export GOROOT=/usr/lib/go
export GOPATH=$HOME/go
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

Once complete, simply source it: `source ~/.bashrc`

Then, git clone the project and go get dependencies:

```
cd ~/go/src/
git clone https://github.com/dveleztx/golang-echo-gorm-graphql.git
go get github.com/oxequa/realize
go get github.com/labstack/echo
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get github.com/graphql-go/graphql
go get github.com/graphql-go/handler
```


### PostgreSQL Setup and Configuration

Move CSV to `/tmp` directory and then enter PSQL shell

```
cd ~/go/src/
cp golang-echo-gorm-graphql/data/users_db.csv /tmp
sudo su - postgres -c psql
```

In PSQL Shell:

```
ALTER USER postgres WITH PASSWORD 'postgres';  -- use better passwords obviously, this is only a demo
CREATE DATABASE sample_db;
\c sample_db
CREATE TABLE users (
	id		int primary key,
	name		varchar(50) not null,
	age		int,
	created_at	date,
	updated_at	date,
	deleted_at	date
);
COPY users(id,name,age,created_at,updated_at,deleted_at) FROM '/tmp/users_db.csv' DELIMITER ',' CSV HEADER;
```

### Using Realize

In the project root of the golang project, run `realize start`

## References

Link: https://medium.com/@manakuro/go-api-with-echo-gorm-and-graphql-1565ad921626 
