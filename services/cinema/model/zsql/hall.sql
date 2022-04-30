CREATE TABLE hall(
    hall_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '影厅ID' ,
    row INT    COMMENT '行数' ,
    col INT    COMMENT '列数' ,
    seat TEXT    COMMENT '座位的矩阵【0不可用 1可用】' ,
    cinema_id INT    COMMENT '影院ID' ,
    hall_name VARCHAR(255)    COMMENT '影厅名称' ,
    PRIMARY KEY (hall_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影厅信息';