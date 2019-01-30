#this script install mysql and do a minimal set of customization
exec &> logfile.txt

##### Variables
db_password=Tito2016

echo
echo "####### BIND VARIABLES #######"
echo "db_password=$db_password" > tito.log 

service mariadb start

#Conf MySql
mysql -u root <<EOF
FLUSH PRIVILEGES;
CREATE USER 'root'@'%' IDENTIFIED BY '$db_password';
SET PASSWORD for root@'%' = PASSWORD('$db_password');
SET PASSWORD for root@'localhost' = PASSWORD('$db_password');
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
create database if not exists TitoDB;
use TitoDB;
CREATE TABLE TitoTable (id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, home VARCHAR(50) NOT NULL, work VARCHAR(50) NOT NULL, hour_home_departure VARCHAR(50) NOT NULL, hour_work_departure VARCHAR(50) NOT NULL);
EOF

