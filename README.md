docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e "DB_SOURCE=postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable"  simplebank:latest

docker build -t simplebank:latest .

docker network create bank-network


- Create a new migration :
    migrate create -ext sql -dir db/migration -seq <migration_name>