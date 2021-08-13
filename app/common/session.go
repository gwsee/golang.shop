package common

import (
	"sync"
	"time"
)

// 定义一个map 用于存储数据
//var providers = make(map[string]interface{})
// 1：我们要有个session接口可以读取 删除 设置session的值
type Session interface {
	Set(key string, tags, value interface{}, maxTime time.Time) error // 设置Session
	Get(key string) interface{}                                       // 获取Session
	Del(key string) error                                             // 删除Session
	Save() error
	SID() string // 当前Session ID
}

//设置session的时候需要设置有效时长
//func (s sessionMgr)Set(key string,tags,value interface{},maxTime time.Time)  {
//	s.lock.Lock()
//	defer s.lock.Unlock()
//	if exsits, ok := providers[key]; ok {
//		exsits["a"]=key
//	} else {
//		fmt.Println("KEY not exist")
//	}
//}
//获取session的时候需要更新最后操作时间
func (s sessionMgr) Get(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
}

//删除session可以根据key 也可以根据标签进行删除
func (s sessionMgr) Del(key string, tags interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
}

//func (s sessionMgr)SID() string {
//	return s.sessionId
//}
// 2：上述的每个session的值需要存储在哪里  建立一个结构体进行存储
type sessionMgr struct {
	sessionMap map[string]interface{}
	lock       sync.Mutex
}

/**
r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	sessionNames := []string{"a", "b"}
	r.Use(sessions.SessionsMany(sessionNames, store))

	r.GET("/hello", func(c *gin.Context) {
		sessionA := sessions.DefaultMany(c, "a")
		sessionB := sessions.DefaultMany(c, "b")

		if sessionA.Get("hello") != "world!" {
			sessionA.Set("hello", "world!")
			sessionA.Save()
		}

		if sessionB.Get("hello") != "world?" {
			sessionB.Set("hello", "world?")
			sessionB.Save()
		}

		c.JSON(200, gin.H{
			"a": sessionA.Get("hello"),
			"b": sessionB.Get("hello"),
		})
	})
	r.Run(":8000")
*/
func InitSession() {
	//store:=cookie.NewStore([]byte("secret"))
	//sessionNames := []string{"a", "b"}
	//fmt.Println(sessionNames)
	//sessions.SessionsMany(sessionNames, store)
	//r.Use(sessions.SessionsMany(sessionNames, store))
}

/**
r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8000")
*/
func InitSessionRedis() {
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//r.Use(sessions.Sessions("mysession", store))
	//
	//r.GET("/incr", func(c *gin.Context) {
	//	session := sessions.Default(c)
	//	var count int
	//	v := session.Get("count")
	//	if v == nil {
	//		count = 0
	//	} else {
	//		count = v.(int)
	//		count++
	//	}
	//	session.Set("count", count)
	//	session.Save()
	//	c.JSON(200, gin.H{"count": count})
	//})
}

//func EnableGinCookieSession() gin.HandlerFunc{
//	store:=cookie.NewStore([]byte("MY_KEY"))
//	return sessions.Sessions("SESSIONID",store)
//}
