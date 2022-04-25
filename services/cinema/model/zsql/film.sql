CREATE TABLE film(
    film_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '影片ID' ,
    film_name VARCHAR(255)   DEFAULT '' COMMENT '影片名称' ,
    film_desc VARCHAR(255) NOT NULL  DEFAULT '' COMMENT '影片描述' ,
    duration INT   DEFAULT 0 COMMENT '影片时长' ,
    cover_pic VARCHAR(255)   DEFAULT '' COMMENT '影片封面图' ,
    type TINYINT UNSIGNED   DEFAULT 1 COMMENT '影片类型 1:2d 2:3d' ,
    cate TINYINT UNSIGNED NOT NULL  DEFAULT 1 COMMENT '影片分类 1古装剧 2动作片 2历史剧' ,
    PRIMARY KEY (film_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片信息表';