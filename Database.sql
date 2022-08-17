CREATE DATABASE TODOLIST; 

CREATE TABLE TODOS(
    `TodoId` INT AUTO_INCREMENT,
    `Text` VARCHAR(255) NOT NULL DEFAULT "",
    `DONE` BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY(`TodoId`));
