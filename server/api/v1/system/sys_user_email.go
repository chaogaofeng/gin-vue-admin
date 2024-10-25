package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	emailUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"regexp"
	"strings"
	"time"
)

func IsValidEmail(email string) bool {
	// 定义邮箱正则表达式
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// 编译正则表达式
	re := regexp.MustCompile(emailRegex)
	// 返回是否匹配
	return re.MatchString(email)
}
func RemoveEmailSuffix(email string) string {
	// 检查邮箱中是否包含 @ 符号
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return email
	}
	// 返回 @ 符号前的部分
	return email[:atIndex]
}

// EmailCaptcha
// @Tags      Base
// @Summary   生成邮箱验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Param    data  body      systemReq.EmailCaptcha                                             true  "邮箱"
// @Produce   application/json
// @Success   200  {object}  response.Response{data={},msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/emailCaptcha [post]
func (b *BaseApi) EmailCaptcha(c *gin.Context) {
	var info systemReq.EmailCaptcha
	err := c.ShouldBindJSON(&info)
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
		return
	}

	if !IsValidEmail(info.Email) {
		global.GVA_LOG.Error("邮箱地址错误!", zap.Error(err))
		response.FailWithMessage("邮箱地址错误", c)
		return
	}

	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, _, answer, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码生成失败!", zap.Error(err))
		response.FailWithMessage("验证码生成失败", c)
		return
	}
	global.BlackCache.Set(info.Email, id, 5*time.Minute)

	err = emailUtils.Email(info.Email, "您收到的验证码", fmt.Sprintf("%v, 验证码有效期5分钟。", answer))
	if err != nil {
		global.GVA_LOG.Error("验证码发送失败!", zap.Error(err))
		//response.FailWithMessage("验证码发送失败", c)
		//return
	}
	global.GVA_LOG.Info("发送验证码", zap.String("邮箱", info.Email), zap.String("验证码", answer))

	response.OkWithDetailed(map[string]interface{}{}, "验证码发送成功", c)
}

// EmailLogin
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.EmailLogin                                             true  "邮箱, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/emailLogin [post]
func (b *BaseApi) EmailLogin(c *gin.Context) {
	var info systemReq.EmailLogin
	err := c.ShouldBindJSON(&info)
	if err != nil {
		global.GVA_LOG.Error("请求参数错误!", zap.Error(err))
		response.FailWithMessage("请求参数错误", c)
	}

	if !IsValidEmail(info.Email) {
		global.GVA_LOG.Error("邮箱地址错误!", zap.Error(err))
		response.FailWithMessage("邮箱地址错误", c)
		return
	}

	v, ok := global.BlackCache.Get(info.Email)
	if !ok {
		response.FailWithMessage("验证码不存在或已过期", c)
		return
	}

	captchaId := v.(string)
	if store.Verify(captchaId, info.Captcha, true) {
		u := &system.SysUser{
			Username:    RemoveEmailSuffix(info.Email),
			NickName:    RemoveEmailSuffix(info.Email),
			Password:    "123456",
			AuthorityId: 888,
			Email:       info.Email,
		}
		user, err := userService.LoginEmail(u)
		if err != nil {
			global.GVA_LOG.Error("登录失败!", zap.Error(err))
			response.FailWithMessage("登陆失败!", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}

	response.FailWithMessage("验证码错误", c)
}

// GetContactUs
// @Tags      Base
// @Summary   获取联系我们
// @Produce   application/json
// @Success   200  {object}  response.Response{data=object,msg=string}  "获取联系我们"
// @Router    /base/contact [post]
func (b *BaseApi) GetContactUs(c *gin.Context) {
	config, err := systemConfigService.GetSystemConfig()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(config.ContactUs, "获取成功", c)
}
