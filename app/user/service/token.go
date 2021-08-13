package service

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	authU "gwsee.com.api/app/auth/user"
	"gwsee.com.api/app/common"
	modelShop "gwsee.com.api/app/shop/model"
	modelUser "gwsee.com.api/app/user/model"
	"time"
)

//查询登陆信息
func FindToken(info *modelUser.UserToken, token string) (err error) {
	if common.CacheConfig.Name == "redis" {
		//redis方式进行数据读取
		str, err := common.GetByRedis(token)
		if err != nil || str == "" {
			return err
		}
		_ = json.Unmarshal([]byte(str), &info)
		//然后更新redis信息 --不更新吧
	} else {
		sqlstr := "select * from  user_token where token =?"
		err = common.FindTable(info, sqlstr, token)
	}
	return
}

//修改登陆信息
func EditToken(data *modelUser.UserToken, token string, auth *authU.GlobalConfig) (err error) {
	if common.CacheConfig.Name == "redis" {
		//REDIS里面去修改登陆次数？
		//修个token
		tokenStr, _ := json.Marshal(data)
		err = common.SetByRedis(token, tokenStr, 60*60*time.Second)
	} else {
		sqlstr := "UPDATE user_token SET " +
			" user=? ,num =? ,shop=? ,shop_sn=? " +
			", edit_time = ? ,edit_user=? " +
			"where token = ?"
		line, err := common.UpdateTable(sqlstr,
			data.User, data.Num, data.Shop, data.ShopSn,
			time.Now().Unix(), auth.User.UserId,
			token)
		if err != nil {
			return err
		}
		if line < 1 {
			// err = errors.New("修改失败")
			return nil
		}
	}
	return
}

//添加登陆信息
func AddToken(user *modelUser.User, shop *[]*modelShop.Shop, shopSn string) (token string, err error) {
	id := uuid.NewV4()
	token = id.String()
	userStr, _ := json.Marshal(user)
	shopStr, _ := json.Marshal(shop)
	//然后进行插入
	if common.CacheConfig.Name == "redis" {
		//redis方式进行数据保存登陆信息
		var userToken = modelUser.UserToken{
			Token:  token,
			User:   string(userStr),
			Shop:   string(shopStr),
			ShopSn: shopSn,
			Num:    1,
		}
		tokenStr, _ := json.Marshal(userToken)
		err = common.SetByRedis(token, tokenStr, 60*60*time.Second)
	} else {
		sqlstr := "INSERT INTO user_token (add_time,add_user,token,user,shop,shop_sn) VALUES (?,?,?,?,?,?)"
		_, err = common.InsertTable(sqlstr, time.Now().Unix(), user.UserId, token, string(userStr), string(shopStr), shopSn)
	}
	fmt.Println(shopSn, "shopSn")
	return
}
