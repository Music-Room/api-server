#!/bin/sh

export MIGRATE='1'

export PGHOST='postgres'
export PGPORT='5432'
export PGNAME='music_room'
export PGUSER='alex'
export PGPASS='12345'

export RDSHOST='redis'
export RDSPORT='6379'
export RDSNAME='1'

echo ${PGHOST} ${PGPORT} ${PGNAME} ${PGUSER} ${PGPASS} ${RDSHOST} ${RDSPORT} ${MIGRATE}

#git pull origin server


# Delete
docker-compose stop
docker-compose rm -f
docker rm $(docker ps -a -f status=exited -q)
docker rmi $(docker images -a -q)
#docker network rm mr_network

# Create
docker network create my_app
docker-compose up --build
docker-compose logs