create table chase (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  ref VARCHAR(255),
  transtype VARCHAR(20),
  description VARCHAR(255),
  amount FLOAT(10, 2), 
  date VARCHAR(11),

  PRIMARY KEY(id)
);
