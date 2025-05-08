CREATE TABLE
  users (
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(320) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp
  );

CREATE TABLE
  accounts (
    account_id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL,
    name VARCHAR(150) NOT NULL,
    color CHAR(7) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
  );

CREATE TABLE
  transactions (
    transaction_id SERIAL PRIMARY KEY,
    account_id SERIAL NOT NULL,
    description VARCHAR(300) NOT NULL,
    value INT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT current_timestamp,
    type
      SMALLINT NOT NULL,
      created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
      CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts (account_id)
  );

CREATE TABLE categories (
  category_id SERIAL,
  user_id SERIAL,
  name VARCHAR (50),
  PRIMARY KEY (category_id),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE
  transaction_category (
    category_id SERIAL,
    transaction_id SERIAL,
    PRIMARY KEY (transaction_id, category_id),
    CONSTRAINT fk_transaction FOREIGN KEY (transaction_id) REFERENCES transactions (transaction_id),
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (category_id)
  );

