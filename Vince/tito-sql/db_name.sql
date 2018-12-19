create database if not exists Tito4DB;
use Tito4DB;
CREATE TABLE TitoTable (name VARCHAR(50) PRIMARY KEY, picture VARCHAR(2000) NOT NULL, capa INT(3), location VARCHAR(3) NOT NULL, available INT(1));
INSERT INTO TitoTable VALUES("Audi Q7","https://s3-eu-west-1.amazonaws.com/tito4picture/2018-audi-q7-20-tfsi-quattro-komfort.png",70,0.2,1);
INSERT INTO TitoTable VALUES("Tesla Model S"," https://s3-eu-west-1.amazonaws.com/tito4picture/tesla2.png",88,0.1,0);
