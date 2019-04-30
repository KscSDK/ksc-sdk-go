package utils

type UrlInfo struct {
	UseSSL bool
	Locate bool
}

type ServiceInfo struct {
	Service string
	Region string
}

const (
	url = ".api.ksyun.com"
	http = "http"
	https = "https"
)

func Url(urlInfo *UrlInfo,info ServiceInfo) string{
	p := Protocol(urlInfo.UseSSL)
	if urlInfo.Locate && &info.Region!=nil {
		return p + "://"+info.Service+"."+info.Region+url
	}
	return p + "://"+info.Service+url
}

func Protocol(useSSL bool) string{
	if useSSL{
		return https
	}
	return http
}


