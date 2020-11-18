create database music;
use music;

CREATE TABLE artist (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
 first_name VARCHAR (30) NOT NULL,
 last_name VARCHAR (30) NOT NULL,
 gender VARCHAR (30) NOT NULL,
 salt VARCHAR (30) NOT NULL,
 password VARCHAR (30) NOT NULL,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE songs (
song_id INT (15) PRIMARY KEY AUTO_INCREMENT,
user_id INT,
song_title VARCHAR (30) NOT NULL,
category VARCHAR (30) NOT NULL,
release_date DATETIME,
FOREIGN KEY (user_id) references artist (user_id)
);

CREATE TABLE album (
    album_id INT (15) PRIMARY KEY AUTO_INCREMENT,
    album_title VARCHAR (30) NOT NULL,
    user_id INT (15),
    song_id INT (15),
    category VARCHAR (30),
    FOREIGN KEY (user_id) references artist (user_id),
    FOREIGN KEY (song_id) references songs (song_id),
    FOREIGN KEY (category) references songs (category)
);

CREATE TABLE user_login (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
 first_name VARCHAR (30) NOT NULL,
 last_name VARCHAR (30) NOT NULL,
 email VARCHAR (30) UNIQUE NOT NULL,
 salt VARCHAR (30) NOT NULL,
 password VARCHAR (30) NOT NULL,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);


