* Copy, Paste & Adjust for simple CRUD Go code
* createdb - psql [dbname] - add table and columns with constraints and defaults
  * ex. 
  personal-chef=# CREATE TABLE image_add(
  personal-chef(# id serial PRIMARY KEY,
  personal-chef(# title VARCHAR (50) NOT NULL,
  personal-chef(# description VARCHAR (255),
  personal-chef(# img_url VARCHAR (255) NOT NULL,
  personal-chef(# dimension VARCHAR (50) NOT NULL,
  personal-chef(# recommended VARCHAR (50) NOT NULL
  personal-chef(# );
* 