## 常用抽奖活动服务端 - Golang

### 抽奖种类
> 关键点：奖品类型，数量有限，中奖概率，发奖规则  
> 并发安全性问题：互斥锁，队列，CAS递减（atomic）  
> 优化：通过三列减少单个集合的大小

- [年会抽奖](./_demo/1annualMeeting)
- [彩票刮奖](./_demo/2ticket)
- [微信摇一摇](./_demo/3wechatShake)
- [支付宝集福卡](./_demo/4alipayFu)
- [微博抢红包](./_demo/5weiboRedPacket)
- [大转盘](./_demo/6wheel)

未完待续。。。
