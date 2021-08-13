package common

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
	"gwsee.com.api/config"
	"math"
	"strconv"
)

var db *sqlx.DB
var tx *sql.Tx
var dbConfig = new(config.DbMgr)

// 所有的数据库表都有的这些字段
type DbColumn struct {
	AddUser  uint64 `db:"add_user" form:"adduser" json:"-"`
	AddTime  uint64 `db:"add_time" form:"addtime" json:"addtime"`
	EditUser uint64 `db:"edit_user" form:"edituser" json:"-"`
	EditTime uint64 `db:"edit_time" form:"edittime" json:"edittime"`
	IsDel    uint64 `db:"is_del" form:"isdel" json:"-"`
	State    string `db:"state" form:"state" json:"state"`
}

// 定义公用的结构体用于列表数据的组装与返回
type Data struct {
	Code      int         `form:"code" json:"code"`
	Msg       string      `form:"msg" json:"msg"`
	List      interface{} `form:"data" json:"data"`
	PageNo    int         `form:"current" json:"current"`
	PageTotal int         `form:"pageTotal" json:"pageTotal"`
	Total     int         `form:"total" json:"total"`
	PageSize  int         `form:"pageSize" json:"pageSize"`
}

func BuildCount(data *Data) (limit string) {
	if data.PageSize < 1 {
		return ""
	}
	data.PageTotal = int(math.Ceil(float64(data.Total) / float64(data.PageSize)))
	if data.PageNo < 1 || data.PageNo > data.PageTotal {
		data.PageNo = 1
	}
	limit = " limit " + strconv.Itoa((data.PageNo-1)*data.PageSize) + "," + strconv.Itoa(data.PageSize)
	return
}

//func init() {
//	initDB()
//}
func InitDB() {
	//加载配置
	err := ini.MapTo(dbConfig, "./config/db.ini")
	if err != nil {
		//	c.JSON(200, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	//获取链接 --这里后面配置链接其他数据库的方式
	dsn := dbConfig.MysqlConfig.Username + ":" + dbConfig.MysqlConfig.Password
	dsn = dsn + "@tcp(" + dbConfig.MysqlConfig.Hostname + ":" + strconv.Itoa(dbConfig.MysqlConfig.Hostport) + ")"
	dsn = dsn + "/" + dbConfig.MysqlConfig.Database
	//链接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		//	c.JSON(200, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	// 不手动关的话，可能在你程序结束的时候，或者这个连接对象出了作用域的时候，
	// golang会帮你自动关闭。。。不过还是支持手动关闭。。。这是好习惯。。。
	// defer  db.Close()
	// 自己手动关闭的话 又可能会导致多线程的时候 关闭过早导致没法进行数据操作
	db.SetMaxOpenConns(dbConfig.MysqlConfig.Maxopen)
	db.SetMaxIdleConns(dbConfig.MysqlConfig.Maxidle)
	// c.Next() //先跑外面的 跑结束了 再执行return
	//判断是否有线程在跑
	return
}

//用于执行分页的查询
/**
var userList []*user_info
err = Db.Select(&userList,"SELECT username,email FROM user_info WHERE user_id>5")
if err != nil{
    fmt.Println(err)
    return
}
fmt.Println(userList)
for _,v:= range userList{
    fmt.Println(v)
}
*/
// 要自己组装sql 才能进行赋值
func ListTable(list interface{}, sqlstr string) (err error) {
	logSql("list", sqlstr)
	err = db.Select(list, sqlstr)
	return
}
func CountTable(sqlstr string) (count int, err error) {
	logSql("count", sqlstr)
	db.QueryRow(sqlstr).Scan(&count)
	return
}
func TransTable(flag int) (err error) {
	if flag == 1 {
		//提交事务
		tx.Commit()
		tx = nil
	} else if flag == 2 {
		//回滚事务
		tx.Rollback()
		tx = nil
	} else {
		//开启事务
		tx, err = db.Begin()
	}
	return
}

//用于执行原生sql查询
/**
rows,err := Db.Query("SELECT email FROM user_info WHERE user_id>=5")
if err != nil{
    fmt.Println("select db failed,err:",err)
    return
}
// 这里获取的rows是从数据库查的满足user_id>=5的所有行的email信息，rows.Next(),用于循环获取所有
for rows.Next(){
    var s string
    err = rows.Scan(&s)
    if err != nil{
        fmt.Println(err)
        return
    }
    fmt.Println(s)
}
rows.Close()
// 事务处理
*/

func QueryTable(info interface{}, sqlstr string, p ...interface{}) (err error) {
	rows, err := db.Query(sqlstr, p)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(info)
		if err != nil {
			return
		}
		//data=append(data,fmt.Sprintf("s%d",i));
	}
	return
}

//用于执行单个查询 -- i用于接收多个参数进行单个查询
func FindTable(info interface{}, sqlstr string, p ...interface{}) (err error) {
	logSql("find", sqlstr, p)
	err = db.Get(info, sqlstr, p...)
	if err != nil {
		if err != sql.ErrNoRows {
			return
		}
		err = nil
	}
	return
}

//用于执行新增 -- 返回新增的id
/**
  result,err := Db.Exec("INSERT INTO user_info(username,sex,email)VALUES (?,?,?)","user01","男","8989@qq.com")
  if err != nil{
      fmt.Println("insert failed,",err)
  }
  // 通过LastInsertId可以获取插入数据的id
  userId,err:= result.LastInsertId()
  // 通过RowsAffected可以获取受影响的行数
  rowCount,err:=result.RowsAffected()
  fmt.Println("user_id:",userId)
  fmt.Println("rowCount:",rowCount)
*/
func InsertTable(sqlstr string, p ...interface{}) (id int64, err error) {
	logSql("insert", sqlstr, p...)
	var res sql.Result
	if tx != nil {
		res, err = tx.Exec(sqlstr, p...)
	} else {
		res, err = db.Exec(sqlstr, p...)
	}
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	if id < 1 {
		err = errors.New("添加失败")
	}
	return
}

//用于执行修改 -- 也可以用于删除 --返回受到影响的行数
/**
results,err := Db.Exec("UPDATE user_info SET username=? where user_id=?","golang",5)
if err != nil{
    fmt.Println("update data fail,err:",err)
    return
}
fmt.Println(results.RowsAffected())
*/
func UpdateTable(sqlstr string, p ...interface{}) (line int64, err error) {
	logSql("update", sqlstr, p...)
	var res sql.Result
	if tx != nil {
		res, err = tx.Exec(sqlstr, p...)
	} else {
		res, err = db.Exec(sqlstr, p...)
	}
	if err != nil {
		return
	}
	line, err = res.RowsAffected()
	if err != nil {
		return
	}
	if line < 1 {
		err = errors.New("操作失败")
	}
	return
}

//用于执行删除
func DeleteTable() (err error) {
	return nil
}

type dbSql struct {
	Type string
	Sql  string
	Data interface{}
}

var log []string

func logSql(t, sql string, p ...interface{}) {
	var l dbSql
	l.Type = t
	l.Sql = sql
	l.Data = p
	str, _ := json.Marshal(l)
	log = append(log, string(str))
}

func GetLog() []string {
	defer func() {
		log = nil
	}()
	return log
}
