#!/bin/bash
docker run -d -p 5432:5432 --name champi -e POSTGRES_PASSWORD=champipassword -e POSTGRES_DB=champi postgres
