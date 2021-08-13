始于2020年5月,至今没写完，请见谅




golang写的店铺gin，mysql，redis，nsq阿里云oss，message
微信小程序等

#好记性 不如烂笔头，每个接口需要写好备注说明
m/开头的是提供给会员访问的接口 对应的文件也以m_开头命名；其他则是商家或者公用的<br/>
#如何部署到linux服务器
0：服务器要安装golang  yum install golang
1: 在windows中编译好项目 然后丢到服务器运行即可
编译（如果编译不通过设置些相关信息即可）
SET CGO_ENABLED=0

SET GOOS=linux

SET GOARCH=amd64

go build main.go
运行
./main (不能后台运行)
nohup ./main & （能够后台运行）
或者sudo nohup ./main>log.log 2>&1  &
上述最后那个是能够输出重定向就退出命令了
ps -aux | grep main| grep -v grep  //查询是那个在运行 
然后kill -9 掉这个端口
cd /data/html/demo.gwsee.com/gwsee.com.api
chmod 755 main//文件可执行

nsq是：
下载并且安装在tools里面
wget 'https://s3.amazonaws.com/bitly-downloads/nsq/nsq-1.2.0.linux-amd64.go1.12.9.tar.gz'
启动nsqlookup
nohup ./nsqlookupd &

启动nsqd
nohup ./nsqd --lookupd-tcp-address=127.0.0.1:4160 &

运行 nsqadmin 管理，注意更换IP为服务器地址
nohup ./nsqadmin --lookupd-http-address=0.0.0.0:4161 --http-address=IP:8761 &


# 目标：
做个项目，能够用户注册，发布商品，然后用户购买的系统<br/>
##-1:定义表内常见
排序 数值大的一般排在前面<br/>
数据的不限量 用-1表示 不要用0 <br/>

##0：定义接口输出基本规则
-1：必须登陆才能获取数据，此时弹出登陆框或者跳转登陆页面<br/>
0：数据操作失败的时候的错误提示，此时需要报错（一般情况下 在正式环境下不能报错，
   在测试和开发环境中报错，在正式环境中应该配置日志监控系统 出现 code=0的时候,需要突出显示<br/>
1：数据正常，结果不需要提醒<br/>
2：数据正常，需要进行消息提醒<br/>
3：数据正常，需要自定义消息提醒方式的。
##1：平台管理--主要是system
用于数据监控，服务注册等<br/>
1：提供给商家的商品买卖系统<br/>
2：提供给商家的会员系统<br/>
3：提供给商家的任务系统<br/>
4：比如用户访问（行为）记录分析analysis<br/>
5：系统基本信息与数据管理，各种服务开启与管理；
##2：用户中心--user
用于用户注册，处理关系的地方<br/>
1：用户注册，到user表<br/>
2：用户与商户shop的关系在shop_user(表示这个shop有那些用户)<br/>
3：默认第一个shop的第一个user管理员为系统管理员；用于system的管理（且不需要审核）；
其他shop的申请注册者就必须通过这个shop的管理；
同事这个shop的管理者 还能管理网站的域名等各种基本信息设置；（这个主要用于管理上面的system）<br/>

