package comm

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/kataras/iris"

	"go-lottery/conf"
	"go-lottery/models"
)

func ClientIp(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

func SetLoginUser(writer http.ResponseWriter, loginUser *models.LoginUser) {
	if loginUser == nil || loginUser.Uid < 1 {
		c := &http.Cookie{
			Name:   "lottery_login_user",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}

	if loginUser.Sign == "" {
		loginUser.Sign = CreateLoginUserSign(loginUser)
	}

	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginUser.Uid))
	params.Add("username", loginUser.Username)
	params.Add("ip", loginUser.Ip)
	params.Add("sign", loginUser.Sign)
	params.Add("now", strconv.Itoa(TimeToStamp(loginUser.Now)))
	c := &http.Cookie{
		Name:   "lottery_login_user",
		Value:  params.Encode(),
		Path:   "/",
		MaxAge: 86400 * 30,
	}
	http.SetCookie(writer, c)
}

func GetLoginUser(request *http.Request) *models.LoginUser {
	c, err := request.Cookie("lottery_login_user")
	if err != nil {
		return nil
	}

	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}

	uid, err := strconv.Atoi(params.Get("uid"))

	if err != nil || uid < 1 {
		return nil
	}

	now, err := strconv.Atoi(params.Get("now"))

	// 没有错误，且没有超过三十天
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}

	loginUser := &models.LoginUser{
		Uid:      uid,
		Username: params.Get("username"),
		Now:      StampToTime(now),
		Ip:       ClientIp(request),
		Sign:     params.Get("sign"),
	}

	sign := CreateLoginUserSign(loginUser)
	// 签名是否一致
	if sign != loginUser.Sign {
		log.Println("func_web GetLoginUser createLoginUsrSign not signed",
			sign, loginUser.Sign)
		return nil
	}

	return loginUser
}

func CreateLoginUserSign(loginUser *models.LoginUser) string {
	s := fmt.Sprintf(
		"uid=%d&username=%s&secret=%s&now=%d",
		loginUser.Uid,
		loginUser.Username,
		conf.CookieSecret,
		TimeToStamp(loginUser.Now),
	)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

type Content struct {
	iris.Context
	Result conf.Result
}

func (this *Content) InitResult() {
	this.Result = conf.Result{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
}
