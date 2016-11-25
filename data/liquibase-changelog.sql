--liquibase formatted sql

--changeset jpeterson:1
create table chase (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  ref VARCHAR(255),
  transtype VARCHAR(20),
  description VARCHAR(255),
  amount FLOAT(10, 2), 
  date VARCHAR(11),

  PRIMARY KEY(id)
);
create table fcu (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  ref VARCHAR(255),
  transtype VARCHAR(20),
  description VARCHAR(255),
  amount FLOAT(10, 2), 
  date VARCHAR(11),

  PRIMARY KEY(id)
);
create table simple (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  ref VARCHAR(255),
  transtype VARCHAR(20),
  description VARCHAR(255),
  amount FLOAT(10, 2), 
  date VARCHAR(11),

  PRIMARY KEY(id)
);
create table bofa (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  ref VARCHAR(255),
  transtype VARCHAR(20),
  description VARCHAR(255),
  amount FLOAT(10, 2), 
  date VARCHAR(11),

  PRIMARY KEY(id)
);

--changeset jpeterson:2
ALTER TABLE simple MODIFY COLUMN transtype VARCHAR(255)
