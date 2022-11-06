package model

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var conn *redis.Client

func init() {
	log.Println("create redis client.")
	conn = redis.NewClient(
		&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		},
	)
	log.Println(conn.Info().String())
}

func NewSession(ctx *gin.Context, cookieKey, redisValue string) {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic("ランダムな文字作成時にエラーが発生しました。")
	}

	newRedisKey := base64.URLEncoding.EncodeToString(b)

	if err := conn.Set(newRedisKey, redisValue, 0).Err(); err != nil {
		panic("Session登録時にエラーが発生:" + err.Error())
	}

	ctx.SetCookie(cookieKey, newRedisKey, 0, "/", "localhost", false, false)
}

func GetSession(ctx *gin.Context, cookieKey string) interface{} {
	redisKey, _ := ctx.Cookie(cookieKey)
	redisValue, err := conn.Get(redisKey).Result()
	switch {
	case err == redis.Nil:
		log.Println("SessionKeyが登録されていません")
		return nil
	case err != nil:
		log.Println("session取得時にエラーが発生しました", err.Error())
		return nil
	}
	return redisValue

}

func DeleteSession(ctx *gin.Context, cookieKey string) {
	redisId, _ := ctx.Cookie(cookieKey)
	conn.Del(redisId)
	ctx.SetCookie(cookieKey, "", -1, "/", "localhost", false, false)
}
