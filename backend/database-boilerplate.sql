CREATE DATABASE Omiran;
USE DATABASE Omiran;

CREATE TABLE User ( uuid VARCHAR(36) PRIMARY KEY, username string TEXT, email VARCHAR(255), password text, profile_picture VARCHAR(2083) );

CREATE TABLE Follows ( uuid VARCHAR(36) NOT NULL, user_following VARCHAR(36) NOT NULL );