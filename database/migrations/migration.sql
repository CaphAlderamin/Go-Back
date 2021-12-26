CREATE TABLE `checkList`
(
    id INTEGER AUTO_INCREMENT,
    taskName VARCHAR(255) NOT NULL,
    taskDescription TEXT NOT NULL,
    taskStatus BOOL,
    PRIMARY KEY (id)
) COMMENT='this is main check-list table';

INSERT INTO `checkList` (`taskName`, `taskDescription`,`taskStatus`)
VALUES ('Wake up', 'I need to wake up at 9 am', true),
       ('Do RIP labs', 'I need to do RIP labs', false),
       ('Now i work', 'time is now', false);

#select * from checkList;