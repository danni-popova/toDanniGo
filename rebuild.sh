docker-compose down
docker volume rm todannigo_postgres-data
docker-compose build
docker-compose up
