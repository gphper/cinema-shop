CREATE TABLE cinema(
    cinema_name VARCHAR(255)    COMMENT '影院名称' ,
    cinema_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '影院ID' ,
    place VARCHAR(255)    COMMENT '地点' ,
    city VARCHAR(255)    COMMENT '市级编号' ,
    area VARCHAR(255)    COMMENT '县区编号' ,
    score INT    COMMENT '评分' ,
    tags VARCHAR(900)    COMMENT '标签，多个使用,分割' ,
    PRIMARY KEY (cinema_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影院管理';
CREATE INDEX union_index_ciyt_area ON cinema(city,area);