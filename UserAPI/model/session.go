package model

import (
	"encoding/json"
	db "v1/Gin/UserAPI/database"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/pkg/errors"
	"github.com/gomodule/redigo/redis"
)

// 描述：插入一个 session 对象
// session 顶级 key，顶级 key 可以设置过期时间
// <[session]： 要插入的 session 对象
// >[error]：插入失败相关信息

// 保存session
func (s *SessionService) SetSession(session *Session) error {
	// 1、从池里获取连接
	conn := db.Pool.Get()
	if conn == nil {
		log.Error("redis connnection is nil")
		return errors.New("redis connnection is nil")
	}
	// 2、用完后将连接放回连接池
	defer conn.Close()

	// 3、将session转换成json数据，注意：转换后的value是一个byte数组
	value, err := json.Marshal(session)
	if err != nil {
		log.Error("json marshal err,%s", err)
		return err
	}
	log.Info("send data[%s]", session.SessionID, value)

	// 4、将SessionID作为key，json数据转换后的byte数组作为value，写入到Redis
	_, err = conn.Do("SET", session.SessionID, value, "EX", 30)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

//  描述： 删除一个 session 对象
//	session 顶级 key，一般情况下 session 会在用户无操作 30 分钟后自行过期删除
//	但用户登出操作可以提前对 session 进行删除，这就是本方法被调用的地方
//	<[sessionID]： 要删除的 session 对象的 id
//	>[error]：删除失败相关信息

// 删除session
func (s *SessionService) DelSession(SessionID string) error {
	// 1、从连接池获取连接
	conn := db.Pool.Get()
	if conn == nil {
		log.Error("redis connnection is nil")
		return errors.New("redis connnection is nil")
	}

	// 2、用完后将连接放回连接池
	defer conn.Close()

	// 3、从Redis删除对应的SessionID
	_,err := conn.Do("Del",SessionID)
	if err != nil {
		return err
	}
	log.Info("move data[%s]",SessionID)

	return nil
}

//  描述： 查看并返回一个 session 实体
//	session 顶级 key
//	<[sessionID]： 要查看的 session 对象的 id
//	>[error]：查看失败相关信息

// 查看session
func (s *SessionService)GetSession(SessionID string)(session *Session,err error){
	// 1、从连接池获取连接
	conn := db.Pool.Get()
	if conn == nil {
		log.Error("redis connection is nil")
		return nil,errors.New("redis connection is nil")
	}

	// 2、用完后将连接放回连接池
	defer conn.Close()

	// 3、查看该session是否存在
	var ifExists bool
	ifExists,err = s.ExistSession(SessionID)
	if err != nil {
		log.Error("fail to exists one session(%s):%s",SessionID,err)
		return nil,errors.New("session not exists,SessionID:"+SessionID)
	}

	if ifExists {
		// 4、将JSON数据从Redis中的redis.Bytes转换为[]byte
		valueBytes,err := redis.Bytes(conn.Do("Get",SessionID))
		if err != nil {
			return nil,err
		}

		// 5、反序列化json数据到Session结构体
		session = &Session{}
		err = json.Unmarshal(valueBytes,session)
		if err != nil {
			return nil,err
		}
		return session,nil
	} else {
		return nil,errors.New("session not exists,SessionID:"+SessionID)
	}

}

// 判断SessionID是否存在
func (s *SessionService)ExistSession(SessionID string) (bool,error) {
	// 1、从连接池获取连接
	conn := db.Pool.Get()
	if conn == nil {
		log.Error("redis connection is nil")
		return false,errors.New("redis connection is nil")
	}

	// 2、用完后将连接放回连接池
	defer conn.Close()

	// 3、从Redis中获取SessionID
	_,err := redis.Bytes(conn.Do("Get",SessionID))
	if err != nil {
		log.Error("fail to exists one session(%s):%s",SessionID,err)
		return false,errors.New("session not exists,SessionID:"+SessionID)
	}

	return true,nil
}