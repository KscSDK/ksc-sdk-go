package utils

type UrlInfo struct {
	UseSSL      bool
	Locate      bool
	UseInternal bool
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	url         = ".api.ksyun.com"
	internalUrl = "internal.api.ksyun.com"
	http        = "http"
	https       = "https"
)

func Url(urlInfo *UrlInfo, info ServiceInfo) string {
	p := Protocol(urlInfo.UseSSL)
	if urlInfo.UseInternal {
		return p + "://" + internalUrl
	}
	if urlInfo.Locate && &info.Region != nil {
		return p + "://" + info.Service + "." + info.Region + url
	}
	return p + "://" + info.Service + url
}

func Protocol(useSSL bool) string {
	if useSSL {
		return https
	}
	return http
}
