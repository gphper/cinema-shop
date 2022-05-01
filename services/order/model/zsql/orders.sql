CREATE TABLE order(
    order_id INT(32) UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '订单ID' ,
    order_sn VARCHAR(255)    COMMENT '订单编号' ,
    screen_id INT    COMMENT '排片ID' ,
    created_at DATETIME    COMMENT '创建时间' ,
    uid INT(32) UNSIGNED    COMMENT '用户ID' ,
    updated_at DATETIME    COMMENT '更新时间' ,
    amount INT(32) UNSIGNED    COMMENT '订单金额' ,
    pay_time DATETIME    COMMENT '支付时间' ,
    status TINYINT UNSIGNED    COMMENT '订单状态【1待支付 2支付完成 3检票成功 4已退票 5自动取消 6已过期】' ,
    PRIMARY KEY (order_id)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '订单';
