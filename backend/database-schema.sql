CREATE DATABASE Omiran;
USE Omiran;

CREATE TABLE User (
	uuid VARCHAR(36) PRIMARY KEY,
	username VARCHAR(40) NOT NULL,
	email VARCHAR(255) NOT NULL,
	password TEXT NOT NULL,
	description TEXT,
	profile_picture VARCHAR(2083),
  UNIQUE KEY (username)
);

CREATE TABLE Follows (
	uuid VARCHAR(36) NOT NULL,
	user_following VARCHAR(36) NOT NULL,
	FOREIGN KEY (uuid) REFERENCES User(uuid) ON DELETE CASCADE,
	FOREIGN KEY (user_following) REFERENCES User(uuid) ON DELETE CASCADE,
	PRIMARY KEY (uuid, user_following)
);
