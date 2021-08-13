####命名规范
所有的地方都是强制性的 动词+表名的方式，比如添加是add 查询是find 列表是list 修改是edit 删除是del<br/>
比如修改user_token是 EditToken（因为user是这个user模块的基本）如果是修改user_token_log 就应该EditTokenLog，
如果是修改其中某个字段 比如状态就应该是eidtTokenState,同时每个地方需要加备注方便人认识。
####model层=entity层。`数据层`
存放我们的实体类，与数据库中的属性值基本保持一致。<br/>
####service层。`业务逻辑`
存放业务逻辑处理，也是一些关于数据库处理的操作，但不是直接和数据库打交道，
他有接口还有接口的实现方法，在接口的实现方法中需要导入mapper层，mapper层
是直接跟数据库打交道的，他也是个接口，只有方法名字，具体实现在mapper.xml
文件里，service是供我们使用的方法。在实际开发中的Service层可能被处理为实
体Service层，而不是接口，业务逻辑直接写在Service（Class，不是Interface）
层中，Controller直接调用Service，Service调用Mapper。当然了，Service之间
也是可以互相调用！<br/>
####mapper层=dao层 `这个目前不用`
现在用mybatis逆向工程生成的mapper层，其实就是dao层。对数据库进行数据持久
化操作，他的方法语句是直接针对数据库操作的，而service层是针对我们controller，
也就是针对我们使用者。service的impl是把mapper和service进行整合的文件。<br/>
####controller层。`请求层`
控制器，导入service层，因为service中的方法是我们使用到的，controller通过
接收前端传过来的参数进行业务操作，在返回一个指定的路径或者数据表。<br/>
####整体参数规划。`参数`
登陆类的：<br/>
_token:表示登陆获取的token 存储在header中（后续做token次数管控等，有效时长1个月）<br/>
分享类的：<br/>
_shareUserid:主要用于分享者分享的人id，用于分享分销系统<br/>
_shareLogid:主要用于此次分享，分享者带来的人点击了哪些页面<br/>
_share类的（（即最后一次访问结束之后，无访问1小时后失效），连续有效时间是1个小时）<br/>

session类的:<br/>
为了避免session或者缓存滥用，可能性的发生命名重复：建议用以下规则<br/>
1：缓存与session均需要包含 命名与标签 外加 值与有效期<br/>
2：命名以当前主要取自表名开头然后，标签命名以涉及表名加,（逗号）的方式进行--表名均需完整
3：用表名是为了方便查找，标签是为了方便将有关单位都直接更新或者替换


