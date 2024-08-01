# rest-test

docker run --name mysql-db -dp 3306:3306 -v /mylocaldata:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=secret mysql:latest

docker run --name myadmin -d --link mysql_db:db -p 8080:80 phpmyadmin/phpmyadmin