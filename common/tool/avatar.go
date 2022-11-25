package tool

import (
	"math/rand"
	"strconv"
	"time"
)

func GetAvatar() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(60)
	return "https://buffet-1306963147.cos.ap-shanghai.myqcloud.com/avatars/" + strconv.Itoa(r) + ".png"
}
