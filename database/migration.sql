CREATE TABLE `checkList`
(
    id INTEGER AUTO_INCREMENT,
    taskName varchar(255) NOT NULL,
    taskDescription TEXT,
    taskUntil datetime,
    taskStatus bool,
    taskEnd datetime,
    PRIMARY KEY (id)
) COMMENT='this is main check-list table';

INSERT INTO `checkList` (`taskName`, `taskDescription`,`taskUntil`,`taskStatus`,`taskEnd`)
VALUES ('Wake up', 'I need to wake up at 9 am', '20211214090000', true, '20211213090331'),
       ('Do RIP labs', 'I need to do RIP labs', '20211215110000', false, '20211213111130');

--select * from checkList;