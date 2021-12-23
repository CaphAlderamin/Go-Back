CREATE TABLE `checkList`
(
    id INTEGER AUTO_INCREMENT,
    taskName varchar(255) NOT NULL,
    taskDescription TEXT,
    taskUntil timestamp,
    taskStatus bool,
    taskEnd timestamp,
    PRIMARY KEY (id)
) COMMENT='this is main check-list table';

SET time_zone='+00:00';

INSERT INTO `checkList` (`taskName`, `taskDescription`,`taskUntil`,`taskStatus`,`taskEnd`)
VALUES ('Wake up', 'I need to wake up at 9 am', '2021-12-14 09:00:00', true, '20211213090331'),
       ('Do RIP labs', 'I need to do RIP labs', '2021-12-15 11:00:00', false, '20211213111130'),
       ('Now i work', 'time is now', current_time, false, current_time);

#select * from checkList;