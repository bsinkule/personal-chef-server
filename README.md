* Copy, Paste & Adjust for simple CRUD Go code
* createdb - psql [dbname] - add table and columns with constraints and defaults
  * ex. 
  personal-chef=# CREATE TABLE image_adds(
  personal-chef(# id serial PRIMARY KEY,
  personal-chef(# title VARCHAR (50) NOT NULL,
  personal-chef(# description VARCHAR (255),
  personal-chef(# img_url VARCHAR (255) NOT NULL,
  personal-chef(# dimension VARCHAR (50) NOT NULL,
  personal-chef(# recommended VARCHAR (50) NOT NULL
  personal-chef(# );
* test localhost in postman
* go get -u github.com/kardianos/govendor
* govendor init
* govendor add +external
* git add . && git commit -m "alsdjkf"
* heroku create go-personal-chef
* git push heroku master
* heroku addons:create heroku-postgresql:hobby-dev
* 
