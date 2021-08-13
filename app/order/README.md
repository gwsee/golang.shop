##订单的类型与订单的状态 都是在订单的子表里；<br/>

####order：<br/>
订单编号：order_no  <br/>
订单归属：shop_sn   <br/>

订单用户：user_id   <br/>
订单推荐人：user_referee   <br/>
订单对账人：user_verify <br/>

订单会员等级：order_grade   <br/>
订单实际总额：order_amount   <br/>
订单支付总额：order_pay   <br/>
订单实际支付时间：pay_time   <br/>
订单支付类型：pay_type   <br/>
订单支付流水号：pay_no   <br/>
订单主类型：order_type  (商品/会员/券等..)<br/> 
订单交易完成：finish_time   <br/>
订单返佣：order_reward 如果后续的推荐人返利的话，在订单状态结算之后才行<br/> 
订单是否评价：order_comment <br/>
订单备注：order_remarks <br/>
订单对账：order_verify <br/>
订单对账时间：verify_time <br/>
订单来源：order_from (wx小程序，H5,PC,APP等)<br/>
订单支付时候手机号：pay_mobile <br/>


####order_detail：明细表<br/>
订单编号：order_no  <br/>
订单商品：goods_sn  <br/>
订单商品名称：goods_title  <br/>
订单商品单价：goods_price  <br/>
订单商品数量：goods_num  <br/>
订单实际金额(总额)：detail_price  <br/>
订单支付金额：detail_pay  <br/>
订单明细类型：detail_type（服务，支付，报名，什么类型会员卡，什么类型券等）  <br/>
订单是否需要物流：detail_transport //如果是需要的直接绑定订单物流表的ID  <br/>
订单状态：     <br/>

####order_comment：评论表<br/>
####order_discount：优惠（折扣）表<br/>
####order_service：订单服务表 <br/>
订单编号：order_no  <br/>
订单商品：goods_sn  <br/>
服务客户：user_name  <br/>
服务电话：user_mobile <br/>
服务顾问：service_server  <br/>
服务类型: service_id  <br/>
服务耗时: service_minutes  <br/>
服务开始时间: service_begin  <br/>
服务结束时间: service_end  <br/>

