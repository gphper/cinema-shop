CREATE TABLE tickets(
    ticket_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '门票ID' ,
    order_id INT    COMMENT '订单ID' ,
    ticket_sn VARCHAR(255)    COMMENT '门票编号' ,
    created_at datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    screen_id INT(32) UNSIGNED    COMMENT '排片ID' ,
    status TINYINT UNSIGNED    COMMENT '状态【1待支付 2支付完成 3检票成功 4已退票 5自动取消 6已过期】' ,
    check_time DATETIME    COMMENT '检票时间' ,
    seat VARCHAR(255)    COMMENT '座位编号【1#2 1排2座】' ,
    PRIMARY KEY (ticket_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影票';
CREATE INDEX order_id_index ON tickets(order_id);