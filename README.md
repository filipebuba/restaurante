# rest-test

docker run --name mysql-db -dp 3306:3306 -v /mylocaldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root mysql:latest

docker run --name myadmin -d --link mysql_db:db -p 8080:80 phpmyadmin/phpmyadmin

## Para parar todos os comandos em execução
docker stop $(docker ps -q)

## Para remover todos os camendos em execução 
docker rm $(docker ps -a -q)