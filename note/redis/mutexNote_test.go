package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/hjr265/redsync.go/redsync"
	"testing"
)

/*
redis为SET命令增加了一系列选项:
EX seconds – 设置键key的过期时间，单位时秒
PX milliseconds – 设置键key的过期时间，单位时毫秒
NX – 只有键key不存在的时候才会设置key的值
XX – 只有键key存在的时候才会设置key的值
*/
func TestMutex(t *testing.T) {
	var pools []redsync.Pool
	m, err := redsync.NewMutexWithGenericPool("FlyingSquirrels", pools)
	if err != nil {
		panic(err)
	}
	//获取锁
	//通过 set nx px
	//仅当 key不存在时进行操作，并设置失效时间，单位时毫秒
	err = m.Lock()
	if err != nil {
		panic(err)
	}

	//通过 lua 脚本释放锁
	defer m.Unlock()

	//延长锁持有时间
	m.Touch()

	t.Logf("delScript : %v \n", delScript)
	t.Logf("touchScript : %v \n", touchScript)
}

//释放锁脚本
//判断对应的value是否正确，正确则释放删除key
var delScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

//延长锁持有时间脚本
// set xx 表示 key 存在才进行操作
// px 设置key的过期时间，单位时毫秒
var touchScript = redis.NewScript(1, `
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("set", KEYS[1], ARGV[1], "xx", "px", ARGV[2])
else
	return "ERR"
end`)
