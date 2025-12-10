package config

import "time"

const (
	TokenAPI        = "https://api.steamdl.ir/get_user?token=%s"
	DNSListAPI      = "https://files.steamdl.ir/anti_sanction_dns.json"
	CacheDomain     = "dl.steamdl.ir"
	LocalIP         = "127.0.0.1"
	DNSPort         = "53"
	HTTPPort        = "80"
	HTTPSPort       = "443"
	StatsFile       = "rx.txt"
	StatsUpdateFreq = 2 * time.Second
)

var globalDomains = []string{
	"dl.steamdl.ir.",
}

var steamDomains = []string{
	"lancache.steamcontent.com.",
}

var playstationDomains = []string{
	"*.gs2-ww-prod.psn.akadns.net",
	"*.gs2.sonycoment.loris-e.llnwd.net",
	"*.gs2.ww.prod.dl.playstation.net",
	"*.gs2.ww.prod.dl.playstation.net.edgesuite.net",
	"gs-sec.ww.np.dl.playstation.net",
	"gs2-ww-prod.psn.akadns.net",
	"gs2.ww.prod.dl.playstation.net",
	"gs2.ww.prod.dl.playstation.net.edgesuite.net",
	"gst.prod.dl.playstation.net",
	"playstation4.sony.akadns.net",
	"psnobj.prod.dl.playstation.net",
	"sgst.prod.dl.playstation.net",
	"theia.dl.playstation.net",
	"tmdb.np.dl.playstation.net",
	"uef.np.dl.playstation.net",
	"vulcan.dl.playstation.net",
}

var xboxDomains = []string{
	"assets1.xboxlive.com",
	"assets1.xboxlive.com.nsatc.net",
	"assets2.xboxlive.com",
	"d1.xboxlive.com",
	"xbox-mbr.xboxlive.com",
	"xvcf1.xboxlive.com",
	"xvcf2.xboxlive.com",
}

var riotDomains = []string{
	"*.dyn.riotcdn.net",
	"l3cdn.riotgames.com",
	"riotgamespatcher-a.akamaihd.net",
	"riotgamespatcher-a.akamaihd.net.edgesuite.net",
	"worldwide.l3cdn.riotgames.com",
}

var epicDomains = []string{
	"cdn.unrealengine.com",
	"cdn1.epicgames.com",
	"cdn1.unrealengine.com",
	"cdn2.epicgames.com",
	"cdn2.unrealengine.com",
	"cdn3.unrealengine.com",
	"cloudflare.epicgamescdn.com",
	"download.epicgames.com",
	"download2.epicgames.com",
	"download3.epicgames.com",
	"download4.epicgames.com",
	"egdownload.fastly-edge.com",
	"epicgames-download1.akamaized.net",
	"fastly-download.epicgames.com",
}

var AllDomains = append(
	append(
		append(
			append(
				append(steamDomains, playstationDomains...),
				xboxDomains...),
			riotDomains...),
		epicDomains...),
	globalDomains...)

// runtime configuration
type AppState struct {
	UserToken     string
	CacheServerIP string
	DNSIP         string
}
