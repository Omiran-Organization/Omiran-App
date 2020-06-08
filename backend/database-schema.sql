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
	follower VARCHAR(36) NOT NULL,
	followee VARCHAR(36) NOT NULL,
	FOREIGN KEY (follower) REFERENCES User(uuid) ON DELETE CASCADE,
	FOREIGN KEY (followee) REFERENCES User(uuid) ON DELETE CASCADE,
	PRIMARY KEY (follower, followee),
  INDEX (followee)
);