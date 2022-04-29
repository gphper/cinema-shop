CREATE TABLE screen(
    screen_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '排片ID' ,
    cinema_id INT    COMMENT '影院ID' ,
    film_id INT    COMMENT '影片ID' ,
    t_date DATE    COMMENT '排片日期' ,
    price INT(32) UNSIGNED   DEFAULT 0 COMMENT '售价(单位：分)',
    start_time TIME    COMMENT '开场时间' ,
    hall_id INT(32) UNSIGNED    COMMENT '影厅id' ,
    current_seat TEXT    COMMENT '当前座位售卖情况【0表示不存在 1待出售 2已出售】' ,
    PRIMARY KEY (screen_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '排片';
CREATE INDEX union_index_tdate_filmid_cinemaid ON screen(t_date,film_id,cinema_id);