package main

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Config struct {
	Mysql struct {
		DataSource string
	}
}

var configFile = flag.String("f", "data.yaml", "the config file")

func main() {
	var c Config
	conf.MustLoad(*configFile, &c)

	conn := sqlx.NewMysql(c.Mysql.DataSource)

	conn.Exec("INSERT INTO `cinema_shop`.`hall`(`hall_id`, `row`, `col`, `seat`, `cinema_id`, `hall_name`) VALUES (1, 4, 5, '17', 1, '测试影厅');")
	conn.Exec("INSERT INTO `user`(`id`, `name`, `email`, `salt`, `password`, `create_time`, `update_time`, `reftoken`) VALUES (6, 'gphper', '570165887@qq.com', 'RPLFZFAYOP', '12d3d691dcacf2bbeedb1b5091d08ae7', '2022-04-08 22:45:01', '2022-05-28 17:32:03', '3342c0a7-95d3-437f-823d-c3b740196de4');")
	conn.Exec("INSERT INTO `cinema`(`cinema_name`, `cinema_id`, `place`, `province`, `city`, `area`, `score`, `tags`) VALUES ('万达影城', 1, 'XX省XX市XX区华人街38号', '22', '33', '44', 100, 'asdas');")
	conn.Exec("INSERT INTO `film`(`film_id`, `film_name`, `film_desc`, `duration`, `cover_pic`, `type`, `cate`) VALUES (1, '变形金刚', '该片主要讲述了一个名叫山姆的青年，因机缘巧合而卷入了变形金刚中汽车人与霸天虎两派抢夺“火种”的战争之中，并与其他众人与汽车人联手打碎霸天虎征服世界野心的故事', 120, '', 1, 2),(2, '绣春刀Ⅱ：修罗战场', '该片讲述了明天启七年，北镇抚司锦衣卫沈炼在追查案件中身陷阴谋，为了证明清白，与少女北斋，同僚裴纶协力查明真相的故事 ', 130, '', 1, 1);")
	conn.Exec("INSERT INTO `screen`(`screen_id`, `cinema_id`, `film_id`, `t_date`, `price`, `start_time`, `hall_id`, `current_seat`, `seat_num`) VALUES (1, 1, 1, '2022-04-26', 10, '21:45:24', 1, '[[0,0,1,1,0],[1,1,1,1,1],[1,1,1,1,1],[1,1,1,1,1]]', 14);")
}
