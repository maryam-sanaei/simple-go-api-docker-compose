CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;
CREATE TABLE IF NOT EXISTS person (
  	id int(11) NOT NULL auto_increment,
  	firstname varchar(255) default NULL,
  	lastname varchar(255) default NULL,
  	PRIMARY KEY (id)
  	) charset=utf8; 
  	
ALTER TABLE person ADD INDEX (id);


/* Add some sample data */
INSERT INTO person(ID, FirstName, LastName)
VALUES
	(NULL, 'User1', 'Test1'),
	(NULL, 'User2', 'Test2'),
	(NULL, 'User3', 'Test3');