FROM postgres:14.2

WORKDIR /docker-entrypoint-initdb.d

COPY ./postgres.sql .