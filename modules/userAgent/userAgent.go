package userAgent

import (
	"strings"
	"regexp"
	"log"
)

/**
	// 使用
	ua := userAgent.NewUserAgent()
	ua.SetUseragent(userAgent1)
 */

type UserAgent struct {
	UserAgentStr string
	Platforms    string
	Browser
	Mobile
}

type Mobile struct {
	IsMobile bool
	Name     string
}

type Browser struct {
	Name    string
	Version string
}

var platforms = make(map[string]string)
var browsers = make(map[string]string)
var mobiles = make(map[string]string)
var browsersSlice []string

func init() {
	config()
}

func NewUserAgent() *UserAgent {
	ua := &UserAgent{}
	return ua
}
func config() {
	platforms["windows nt 6.0"] = "Windows Longhorn"
	platforms["windows nt 5.2"] = "Windows 2003"
	platforms["windows nt 5.0"] = "Windows 2000"
	platforms["windows nt 5.1"] = "Windows XP"
	platforms["windows nt 4.0"] = "Windows NT 4.0"
	platforms["winnt4.0"] = "Windows NT 4.0"
	platforms["winnt 4.0"] = "Windows NT"
	platforms["winnt"] = "Windows NT"
	platforms["windows 98"] = "Windows 98"
	platforms["win98"] = "Windows 98"
	platforms["windows 95"] = "Windows 95"
	platforms["win95"] = "Windows 95"
	platforms["windows"] = "Unknown Windows OS"
	platforms["os x"] = "Mac OS X"
	platforms["ppc mac"] = "Power PC Mac"
	platforms["freebsd"] = "FreeBSD"
	platforms["ppc"] = "Macintosh"
	platforms["linux"] = "Linux"
	platforms["debian"] = "Debian"
	platforms["sunos"] = "Sun Solaris"
	platforms["beos"] = "BeOS"
	platforms["apachebench"] = "ApacheBench"
	platforms["aix"] = "AIX"
	platforms["irix"] = "Irix"
	platforms["osf"] = "DEC OSF"
	platforms["hp-ux"] = "HP-UX"
	platforms["netbsd"] = "NetBSD"
	platforms["bsdi"] = "BSDi"
	platforms["openbsd"] = "OpenBSD"
	platforms["gnu"] = "GNU/Linux"
	platforms["unix"] = "Unknown Unix OS"

	mobiles["mobileexplorer"] = "Mobile Explorer"
	mobiles["palmsource"] = "Palm"
	mobiles["palmscape"] = "Palmscape"

	// Phones and Manufacturers
	mobiles["motorola"] = "Motorola"
	mobiles["nokia"] = "Nokia"
	mobiles["palm"] = "Palm"
	mobiles["iphone"] = "Apple iPhone"
	mobiles["ipod"] = "Apple iPod Touch"
	mobiles["sony"] = "Sony Ericsson"
	mobiles["ericsson"] = "Sony Ericsson"
	mobiles["blackberry"] = "BlackBerry"
	mobiles["cocoon"] = "O2 Cocoon"
	mobiles["blazer"] = "Treo"
	mobiles["lg"] = "LG"
	mobiles["amoi"] = "Amoi"
	mobiles["xda"] = "XDA"
	mobiles["mda"] = "MDA"
	mobiles["vario"] = "Vario"
	mobiles["htc"] = "HTC"
	mobiles["samsung"] = "Samsung"
	mobiles["sharp"] = "Sharp"
	mobiles["sie-"] = "Siemens"
	mobiles["alcatel"] = "Alcatel"
	mobiles["benq"] = "BenQ"
	mobiles["ipaq"] = "HP iPaq"
	mobiles["mot-"] = "Motorola"
	mobiles["playstation portable"] = "PlayStation Portable"
	mobiles["hiptop"] = "Danger Hiptop"
	mobiles["nec-"] = "NEC"
	mobiles["panasonic"] = "Panasonic"
	mobiles["philips"] = "Philips"
	mobiles["sagem"] = "Sagem"
	mobiles["sanyo"] = "Sanyo"
	mobiles["spv"] = "SPV"
	mobiles["zte"] = "ZTE"
	mobiles["sendo"] = "Sendo"

	// Operating Systems
	mobiles["symbian"] = "Symbian"
	mobiles["Symbianos"] = "SymbianOS"
	mobiles["elaine"] = "Palm"
	mobiles["palm"] = "Palm"
	mobiles["series60"] = "Symbian S60"
	mobiles["windows ce"] = "Windows CE"

	// browsers
	mobiles["obigo"] = "Obigo"
	mobiles["netfront"] = "Netfront Browser"
	mobiles["openwave"] = "Openwave Browser"
	mobiles["mobilexplorer"] = "Mobile Explorer"
	mobiles["operamini"] = "Opera Mini"
	mobiles["opera mini"] = "Opera Mini"

	// Other
	mobiles["digital paths"] = "Digital Paths"
	mobiles["avantgo"] = "AvantGo"
	mobiles["xiino"] = "Xiino"
	mobiles["novarra"] = "Novarra Transcoder"
	mobiles["vodafone"] = "Vodafone"
	mobiles["docomo"] = "NTT DoCoMo"
	mobiles["o2"] = "O2"

	// Fallback
	// mobiles["mobile"] = "Generic Mobile"
	mobiles["wireless"] = "Generic Mobile"
	mobiles["j2me"] = "Generic Mobile"
	mobiles["midp"] = "Generic Mobile"
	mobiles["cldc"] = "Generic Mobile"
	mobiles["up.link"] = "Generic Mobile"
	mobiles["up.browser"] = "Generic Mobile"
	mobiles["smartphone"] = "Generic Mobile"
	mobiles["cellphone"] = "Generic Mobile"

	browsersSlice = []string{"Flock", "Chrome", "Opera",
		"MSIE", "Internet Explorer", "ipad", "Shiira", "Firefox",
		"Chimera", "Phoenix", "Firebird", "Camino",
		"Netscape", "OmniWeb", "Safari", "Mozilla",
		"Konqueror", "icab", "Lynx", "Links", "hotjava", "amaya", "IBrowse"}

	browsers["Flock"] = "Flock"
	browsers["Chrome"] = "Chrome"
	browsers["Opera"] = "Opera"
	browsers["MSIE"] = "Internet Explorer"
	browsers["Internet Explorer"] = "Internet Explorer"
	browsers["ipad"] = "iPad"
	browsers["Shiira"] = "Shiira"
	browsers["Firefox"] = "Firefox"
	browsers["Chimera"] = "Chimera"
	browsers["Phoenix"] = "Phoenix"
	browsers["Firebird"] = "Firebird"
	browsers["Camino"] = "Camino"
	browsers["Netscape"] = "Netscape"
	browsers["OmniWeb"] = "OmniWeb"
	browsers["Safari"] = "Safari"
	browsers["Mozilla"] = "Mozilla"
	browsers["Konqueror"] = "Konqueror"
	browsers["icab"] = "iCab"
	browsers["Lynx"] = "Lynx"
	browsers["Links"] = "Links"
	browsers["hotjava"] = "HotJava"
	browsers["amaya"] = "Amaya"
	browsers["IBrowse"] = "IBrowse"

}

func (ua *UserAgent) SetUseragent(useragent string) {
	ua.UserAgentStr = useragent
	ua.setBrowser()
	ua.setMobile()
	ua.setPlatform()
}

func (ua *UserAgent) setMobile() {
	userAgentLower := strings.ToLower(ua.UserAgentStr)
	for k, v := range mobiles {
		if strings.Contains(userAgentLower, k) {
			if strings.Index(userAgentLower, k) != -1 {
				ua.Mobile.IsMobile = true
				ua.Mobile.Name = v
				break
			}
		}
	}
}

func (ua *UserAgent) setBrowser() {
	var matchedSlice [][]string
	for _, v := range browsersSlice {

		regexp, err := regexp.Compile(`(?i)` + v + `.*?([0-9\.]+)`)
		if err != nil {
			log.Panic(err)
		}

		matched := regexp.FindStringSubmatch(ua.UserAgentStr)
		if matched != nil {
			matchedSlice = append(matchedSlice, matched)
		}
	}
	if matchedSlice != nil {
		key := strings.Split(matchedSlice[0][0], "/")
		ua.Browser.Name = key[0]
		ua.Browser.Version = matchedSlice[0][1]
	}
}

func (ua *UserAgent) setPlatform() {
	for k := range platforms {
		regexp, err := regexp.Compile(`(?i)` + k)
		if err != nil {
			log.Panic(err)
		}

		platformsStr := strings.ToLower(regexp.FindString(ua.UserAgentStr))
		if platformsStr != "" {
			ua.Platforms = platforms[platformsStr]
		}
	}
}
