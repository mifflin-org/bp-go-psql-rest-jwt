#!/bin/bash

export PG_HOST=my-database
export PG_PORT=5432
export PG_USER=postgres
export PG_PASSWORD=password
export PG_DB=postgres

docker-compose -f ./docker-compose-install.yml up -d
