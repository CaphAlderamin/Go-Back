# start with base image
FROM mysql:8.0.27

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./database/migrations/*.sql /docker-entrypoint-initdb.d/