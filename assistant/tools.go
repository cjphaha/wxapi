package assistant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)
//用snowflake生成一个随机id 参数是标识机器的作用，如果在分布式系统中使用此函数，需要设置标识符
func GenerateID(workerid int64) int64{
	snowflake,err := NewSnowflake(workerid)
	if err!=nil{
		fmt.Println(err)
	}
	return snowflake.Generate()
}
//gin框架中将请求信息转化为string，适用于c.bind（）获取不到结构体的情况
func ByteToString(origin *gin.Context) string{
	buf := make([]byte, 2024);//定义一个切片
	n, _ := origin.Request.Body.Read(buf);
	str := string(buf[0:n]);//把传送过来的数据转化成了字符串
	return str
}
//获取随机的字符串（数字+字母组合）
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

//返回数据的类型
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
//保留浮点数至两位
func Decimal(value float64) float64 {//保留两位小数
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64);
	return value;
}
//float32 转 String工具类，保留6位小数
func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}
