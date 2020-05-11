# 
#  This file starts up and initialises a PostreSQL DB for dev/test purposes
# 
FROM postgres:alpine
COPY init.sql /docker-entrypoint-initdb.d/