create table transactions (
  id MEDIUMINT NOT NULL AUTO_INCREMENT,
  user_id MEDIUMINT,
  account_id MEDIUMINT,
  bank_name VARCHAR(256) NOT NULL
  bank_type VARCHAR(32) NOT NULL
  transaction_type VARCHAR(20) NOT NULL,
  amount FLOAT(10, 2) NOT NULL, 
  date VARCHAR(11) NOT NULL,
  description VARCHAR(255),
  PRIMARY KEY(id)
);
