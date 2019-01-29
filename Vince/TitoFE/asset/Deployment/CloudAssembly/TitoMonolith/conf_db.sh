service mariadb start

#Conf MySql
mysql -u root <<EOF
FLUSH PRIVILEGES;
CREATE USER 'root'@'%' IDENTIFIED BY 'Tito2016';
SET PASSWORD for root@'%' = PASSWORD('Tito2016');
SET PASSWORD for root@'localhost' = PASSWORD('Tito2016');
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF

