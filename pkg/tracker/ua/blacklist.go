package ua

// Blacklist is a list of User-Agents to ignore.
var Blacklist = []string{
	"!(()&&!|*|*|",
	"${",
	"(null)",
	"***",
	"-1 or",
	"-1' or",
	"-1\" or",
	"/etc/passwd",
	"0+0+0+1",
	"06d8d6b03c5d67d530631759888ad064",
	"10kgp7vb2o",
	"12345",
	"15q0ixzylo",
	"1680033504967",
	"1681696517991",
	"1685169227575",
	"1685804726376",
	"1686401621817",
	"1686705713195",
	"1686705749888",
	"1686805689354",
	"1687171999244",
	"1687905700113",
	"1690974233830",
	"1691513620386",
	"180thestrand",
	"19yh2xtdvo",
	"1a94m83v2o",
	"1b",
	"1c",
	"1d",
	"1ftlldrnto",
	"1imornni4o",
	"1k2sv0a6yo",
	"1kb80glvxo",
	"1kmrvmr8to",
	"1lbpifr8yo",
	"1vdxq0ad2o",
	"2b818e3af0a3928ea9c640987bfce08c",
	"496ffe900b0471c45ed62acedd255020",
	"49dba188c47cddc37a5d69084d5102e9",
	"5zloa7rtayxakp8",
	"74eb044d588799dd5f74d1e70c33a661",
	"7e4ff1f7-3a6f-457b-ae4f-ea4a22dc33dc",
	"7siters",
	"89.0.4389.114",
	"://",
	"<default user agent>",
	"<script>",
	"<title>",
	"[fban",
	"a6-indexer",
	"abonti",
	"accountsd",
	"aceexplorer",
	"ackerm",
	"acoon",
	"active",
	"activebookmark",
	"activerefresh",
	"activeworlds",
	"acumatica",
	"ad muncher",
	"adbeat",
	"adbeat.com",
	"addthis",
	"admantx",
	"administrator",
	"admuncher",
	"adobeuxtech",
	"ads.txt",
	"ae/0.1",
	"aezakmi",
	"afengineurl",
	"affiliatewindow",
	"agent orange",
	"aggregage",
	"ahc",
	"alertra",
	"alexa",
	"alfright",
	"aliappiphone",
	"allentrack",
	"almaden",
	"alwayson",
	"amazingtalker",
	"amazon",
	"amazon music podcast",
	"amazon.com",
	"amiga",
	"amiga-aweb",
	"amigavoyager",
	"analyz",
	"anchor",
	"android_bp_app",
	"anglesharp",
	"anonymized",
	"anonymous",
	"anonymous_agent",
	"anthill",
	"anthropic",
	"anyconnect",
	"anyevent-http",
	"aol explorer",
	"apache",
	"appie",
	"appinsights",
	"apple-pubsub",
	"applicationhealthservice",
	"arachmo",
	"arachni",
	"architext",
	"archive",
	"archiveteam",
	"aria2",
	"arks",
	"arora",
	"arquivo-web-crawler",
	"array",
	"asafaweb",
	"ask jeeves/teoma",
	"askjeevesteoma",
	"asterias",
	"astute",
	"astutesrm",
	"asusrouter",
	"asynchttp",
	"ata-z",
	"atomz",
	"atvoice",
	"authorizeua",
	"auto",
	"avant",
	"avsdevicesdk",
	"awesometechstack",
	"awiton",
	"axios",
	"azureus",
	"baidu",
	"barooders",
	"barracuda sentinel",
	"bdfetch",
	"benefithubwebview",
	"betsie",
	"bidtellect",
	"biglotron",
	"bigpointclient",
	"bin/bash",
	"bingpreview",
	"binlar",
	"bit.ly",
	"bitdiscovery",
	"bjaaland",
	"black magic",
	"blackboard",
	"blackboardsafeassign",
	"blaiz-bee",
	"blocker",
	"blocknote.net",
	"bloglines",
	"bloglovin",
	"blogpulse",
	"blogtrottr",
	"bluecoat drtr",
	"bluecoatdrtr",
	"bluefish",
	"boitho",
	"bonecho",
	"bookmark-manager",
	"bordermanager",
	"bot",
	"bp_app",
	"braavos",
	"brandverity",
	"browsershots",
	"browsex",
	"brutus",
	"btwebclient",
	"bubing",
	"buck",
	"builtwith",
	"bunnyoptimizer",
	"burpcollaborator",
	"bwh3_user_agent",
	"c1647ea35c6eb53d6a562328ec47e",
	"cakephp",
	"camino",
	"camo asset proxy",
	"camoassetproxy",
	"candroid",
	"captivenetworksupport",
	"capture",
	"castro",
	"catch",
	"catchpoint",
	"catexplorador",
	"celestial",
	"cfnetwork",
	"chatterino",
	"check",
	"checklink",
	"checkprivacy",
	"chicken laser",
	"china",
	"chrome-lighthouse",
	"chromeframe",
	"chuck norris",
	"cisco",
	"citoid",
	"clamav",
	"clamavs",
	"classify-workers",
	"clickfunnels-ua",
	"client",
	"clinicsoftware",
	"cloakdetect",
	"cloud",
	"cloudflare",
	"cobweb",
	"coccoc",
	"code87",
	"codjiranode",
	"coffee-mug",
	"coldfusion",
	"collect",
	"collectd",
	"collection@infegy.com",
	"com.the-dots",
	"combine",
	"cometbird",
	"commons-httpclient",
	"comodo_dragon",
	"contentmatch",
	"contentsmartz",
	"convera",
	"core",
	"cortana",
	"cortex",
	"coverscout",
	"crawl",
	"crawler",
	"cron",
	"crossingminds",
	"crowsnest",
	"crusty",
	"curb",
	"cursor",
	"custo",
	"custom",
	"cve-2023-35078",
	"cvvs/",
	"cyberdog",
	"cyberwarcon",
	"d038eacf-ab1a-408e-9f28-d05703d0433c",
	"daemon",
	"dalvik",
	"dap",
	"darcyripper",
	"dareboost",
	"dart/",
	"dataaccessd",
	"datacha0s",
	"datadogagent",
	"datajoe",
	"dataminr",
	"datanyze",
	"dataprovider",
	"daum",
	"daums",
	"davclnt",
	"dbpedia",
	"dc0c3uu2ppjwyph",
	"de.cb.netz",
	"dejaclick",
	"delay",
	"deluge",
	"deno",
	"detector",
	"deusu",
	"developer",
	"dexador",
	"digg",
	"dingtalkbot",
	"discourse",
	"dispatch",
	"dispatchd",
	"disqus",
	"dmbrowser",
	"dnstwist",
	"docoloc",
	"docomo",
	"doluxe",
	"domains project",
	"domaner.xyz",
	"donutp",
	"download",
	"doximity-pipeline",
	"dreampassport",
	"drip",
	"drupact",
	"drupal",
	"dsurf",
	"dts agent",
	"duckduckbot",
	"duckduckgo",
	"durston",
	"dynamic-image",
	"e46df615-2dbc-4311-8217-c4e61c4ed1e2",
	"easy-thumb",
	"easybib",
	"easydl",
	"ebsco",
	"ecatch",
	"ecosearch",
	"elinks",
	"elisabot",
	"email",
	"emailsiphon",
	"emailwolf",
	"embedly",
	"empty user agent",
	"enigmabrowser",
	"epicai",
	"eudaimon",
	"evc-batch",
	"evernote clip resolver",
	"evernoteclipresolver",
	"evolution",
	"exedwnloadmnger",
	"extraireliensnomdomaine",
	"f325b9c5-501c-4b1a-ad9e-c688c5bcee59",
	"facebook",
	"facebookexternalhit",
	"facebookplatform",
	"faraday",
	"farm_option_agent",
	"fasthttp",
	"favicon",
	"favorg",
	"fdm",
	"fdmsd",
	"feed",
	"feedbin",
	"feedburner",
	"feedfetcher",
	"feedreader",
	"ferret",
	"fetch",
	"filedown",
	"filetransport",
	"filter",
	"finder",
	"findlink",
	"findlinks",
	"findthatfile",
	"firephp",
	"flashget",
	"flipboardproxy",
	"flutter",
	"foca",
	"foodient",
	"foofox",
	"force-ws05",
	"foregenix",
	"freesafeip",
	"friendica",
	"fulltext",
	"funnelback",
	"fuzz faster u fool",
	"g-i-g-a-b-o-t",
	"galeon",
	"gaspedaal",
	"gayloader",
	"genieo",
	"getlinkinfo",
	"getright",
	"getthit",
	"geturl",
	"geziyor",
	"ggg",
	"ghost",
	"gibgas",
	"giex",
	"gigablastopensource",
	"git/",
	"github.com",
	"gkoudai",
	"globalprotect",
	"gnip",
	"gnu c",
	"go-http-client",
	"golang",
	"goldfire",
	"gomezagent",
	"gonative",
	"gooblog",
	"goodjudge",
	"googal",
	"google",
	"googleanalytics",
	"goose",
	"gosquared",
	"gozilla",
	"grammarly",
	"granparadiso",
	"greatnews",
	"greenbrowser",
	"gregarius",
	"grouphigh",
	"grub",
	"gsak",
	"gtmetrix",
	"guayoyo",
	"gulliver",
	"guzzlehttp",
	"gvfs",
	"gzip",
	"h002375fzr7s",
	"h6c86lti5ifj_2",
	"hackernews",
	"hackney",
	"handshake",
	"harvest",
	"hatena",
	"headlesschrome",
	"headlines",
	"hello world",
	"hello-world",
	"heritrix",
	"hexometer",
	"hjdicks",
	"hobbit",
	"holmes",
	"hotzonu",
	"htdig",
	"htmlparser",
	"http",
	"httpcomponents",
	"httpfetcher",
	"httpget",
	"https:",
	"httpx",
	"httrack",
	"hubspot",
	"hubspot marketing grader",
	"hubspotmarketinggrader",
	"hwcdn",
	"hydra",
	"hyperbeam",
	"hyperhost.ua",
	"ia_archiver",
	"iab-tech-lab",
	"ibisbrowser",
	"ibrowse",
	"ice browser",
	"ichiro",
	"idbte4m",
	"iframely",
	"ignored",
	"iktomi",
	"ilse",
	"images",
	"importer",
	"index",
	"indy library",
	"infestra",
	"infox-wisg",
	"ingrid",
	"ingridd",
	"innersourcecoach",
	"inreachapplication",
	"insomnia",
	"instapaper",
	"integration",
	"integrity",
	"intelx",
	"internal",
	"internetseer",
	"internetwache",
	"intute",
	"investment",
	"ipcamviewer",
	"ips-agent",
	"ipsum",
	"isilox",
	"iskanie",
	"itjuzi",
	"itunes",
	"iubenda",
	"iubenda-radar",
	"ivoox",
	"ixquick",
	"java",
	"javafx",
	"jeeves",
	"jeode",
	"jersey",
	"jetbrains",
	"jetty",
	"jigsaw",
	"jnanagangothri",
	"jobo",
	"jorgee",
	"joxypoxy",
	"jurgendata",
	"jusprogdns",
	"juta6ja22uumnu3",
	"kaspersky",
	"keurig",
	"keycdn-tools",
	"keynote",
	"kith-app",
	"kizie",
	"knowledge",
	"kopeechka",
	"kulturarw3",
	"kundenzone",
	"kyluka",
	"l9explore",
	"l9tcpid",
	"labjs.pro",
	"laks",
	"larbin",
	"lasso",
	"lavf",
	"leadspotting",
	"lenns.io",
	"letsextract",
	"lfgnusp4yvgiisj",
	"libcurl",
	"libhttp",
	"library",
	"libtorrent",
	"libvlc",
	"libwww",
	"liferea",
	"lightspeedsystemscrawler",
	"lilina",
	"link preview",
	"linkanalyser",
	"linkdex",
	"linklint-checkonly",
	"linkparser",
	"links",
	"linksaver",
	"linkscan",
	"linktiger",
	"linkwalker",
	"lipperhey",
	"livejournal",
	"lmslinkanalysis",
	"loadster",
	"lobste",
	"lockss",
	"logstatistic",
	"loilonote",
	"loli_tentacle",
	"longurl.api",
	"lorem",
	"lotus-notes",
	"lscache_runner",
	"ltmetadataservice",
	"ltx71",
	"lua-resty-http",
	"luasocket",
	"lucee",
	"lucidworks-anda",
	"lwp",
	"lwp-",
	"lwp::simple",
	"lychee",
	"lycos",
	"macbaiduyunguanjia",
	"macocu",
	"magic browser",
	"magpierss",
	"mail",
	"mail.ru",
	"mailchimp",
	"mailchimp.com",
	"mailto",
	"manager",
	"manictime",
	"marcedit",
	"margin",
	"marktplaats",
	"matric editor",
	"matrix-media-repo",
	"mattermost",
	"mechanize",
	"mediahubmx",
	"mediapartners-google",
	"megaproxy",
	"megite",
	"meltwaternews",
	"mention",
	"metainspector",
	"metauri",
	"microblogpub",
	"microsoft bits",
	"microsoft data",
	"microsoft office",
	"microsoft office existence",
	"microsoft office protocol discovery",
	"microsoft windows network diagnostics",
	"microsoft-cryptoapi",
	"microsoft-webdav-miniredir",
	"microsoftbits",
	"microsoftdata",
	"microsoftofficeexistence",
	"microsoftofficeprotocoldiscovery",
	"microsoftteamsroom",
	"microsoftwindowsnetworkdiagnostics",
	"mimas",
	"minefield",
	"mingw32",
	"miniflux",
	"mixmax-linkpreview",
	"mixnodecache",
	"mizilla",
	"mjukisbyxor",
	"mmb1_pedrillo",
	"mnogosearch",
	"moget",
	"mojeek",
	"monit",
	"monitor",
	"moreover",
	"motor",
	"movabletype",
	"mowser",
	"mozila",
	"mozillad.d(compatible;?)",
	"mozilliqa",
	"mr.4x3",
	"msie",
	"msoffice",
	"msray",
	"muckrack",
	"mucommander",
	"muscatferre",
	"my browser",
	"my user agent",
	"mybrowser",
	"mykcm",
	"mypxapp",
	"mytool",
	"myweb",
	"nagios",
	"navermailapp",
	"nearsoftware",
	"neoload",
	"nessus",
	"netants",
	"netcraft",
	"netcraftsurveyagent",
	"netdisk",
	"netluchs",
	"netnewswire",
	"netscape",
	"netsurf",
	"nettrack anonymous web statistics",
	"netvibes",
	"neustarwpm",
	"news",
	"newsfox",
	"newsgator",
	"newspaper",
	"nextcloud-news",
	"nga_wp_jw",
	"nginx",
	"nhcapacketmng",
	"nibbler",
	"nikto",
	"ning",
	"nmap scripting engine",
	"no_user_agent",
	"node-superagent",
	"node.js",
	"nokiac3",
	"nomad",
	"normalized",
	"notetextview",
	"notionembedder",
	"nu.nl",
	"nutch",
	"nuzzel",
	"nvd0rz",
	"oadoi",
	"oast.online",
	"object object",
	"object promise",
	"ocelli",
	"octopus",
	"offbyone",
	"offline",
	"offline explorer",
	"offlineexplorer",
	"ogscrper",
	"ohio-development",
	"okhttp",
	"omgili",
	"onetszukaj",
	"opengraph",
	"openvas",
	"openwave",
	"operat",
	"optimize",
	"order",
	"order/**/",
	"orion",
	"ossproxy",
	"other",
	"ourbrowser",
	"outbrain",
	"pa11y",
	"page2rss",
	"pagespeed",
	"pagething",
	"panscient",
	"pantest",
	"parse",
	"parsijoo",
	"pattern",
	"paypal",
	"payrexx",
	"pcore-http",
	"peach",
	"pear",
	"pear http_request",
	"pearltrees",
	"pentest",
	"penthouse",
	"perimeterx",
	"perl",
	"perman",
	"pg_",
	"phantom",
	"photon",
	"php",
	"pidcheck",
	"pierre smith",
	"pika.style",
	"ping.blo.gs",
	"pingadmin",
	"pingdom",
	"pinner",
	"pinterest",
	"pioneer",
	"pirsch",
	"pixalate",
	"placid.app",
	"platform",
	"player",
	"playmusic",
	"playstarmusic",
	"plumanalytics",
	"pocketimagecache",
	"podcast",
	"polycomsoundpointip",
	"portalmmm",
	"postgenomic",
	"postiviidakko",
	"postman",
	"postrank",
	"powermarks",
	"powerpc amigaos",
	"pr-cy.ru",
	"preview",
	"print",
	"privacybrowser",
	"prlog",
	"probe",
	"profiling",
	"project_patchwatch",
	"prolog/ragno",
	"prometheus",
	"proximic",
	"proxy",
	"psamma",
	"ptst",
	"ptstd",
	"puppeteeragent",
	"pycurl",
	"python",
	"qcg33mtl7hoskfh",
	"qmez",
	"qqdownload",
	"qualys",
	"quip",
	"qwantify",
	"raindrop",
	"rambler",
	"ramblermail",
	"ranksonicsiteauditor",
	"rapidload",
	"raynette_httprequest",
	"react",
	"reader",
	"readpaper",
	"readyou",
	"realdownload",
	"rebelmouse",
	"redalert",
	"report runner",
	"request",
	"reqwest",
	"restsharp",
	"returnb2b",
	"riddler",
	"rigor",
	"risseri",
	"rivva",
	"rnps-action-cards",
	"robozilla",
	"roche-waas",
	"rocket/preload",
	"rockmelt",
	"roku",
	"romeo-santos",
	"rss",
	"rssbandit",
	"rssowl",
	"ruby",
	"rupee",
	"rx bar",
	"saashub",
	"safeassign",
	"saws",
	"scamadviserexternalhit",
	"scan",
	"scan4mail",
	"science traveller international",
	"scientificcommons",
	"scirus",
	"scoop.it",
	"scooter",
	"scope (mars+)",
	"scoutjet",
	"scpitspi-rs",
	"scrape",
	"scraper",
	"scrapy",
	"scrutiny",
	"search",
	"search.marginalia.nu",
	"searchbloxintra",
	"secragon",
	"sectigo",
	"select",
	"selenium",
	"sentry",
	"seo",
	"seostats",
	"seranking",
	"serpreputationmanagementagent",
	"server",
	"set:",
	"seznamemailproxy",
	"shareaza",
	"shiretoko",
	"shockwaveflash",
	"shortlinktranslate",
	"shoutcast",
	"shrinktheweb",
	"sincera",
	"sistrix",
	"site",
	"site24x7",
	"sitecheckerbotcrawler",
	"sitescorebot",
	"sixy.ch",
	"skypeuripreview",
	"sleipnir",
	"slurp",
	"smallproxy",
	"snacktory",
	"snap",
	"snapchat",
	"socialbeeagent",
	"sogou",
	"space bison",
	"spacebison",
	"sparkler",
	"specialdetect",
	"speedmode",
	"speedy",
	"spider",
	"spiderweb",
	"splash",
	"spotify",
	"spring",
	"sprinklr",
	"sprout social",
	"spy",
	"sqlmap",
	"srcedamp",
	"sri lanka",
	"ss00",
	"ssample",
	"ssllabs",
	"statping",
	"statuscake",
	"strider",
	"stumbleupon",
	"stumbleupon.com",
	"subjs",
	"summify",
	"sundance",
	"sunrise",
	"supercleaner",
	"svn",
	"swcd",
	"sweepatic",
	"swisscows",
	"sylera",
	"symfony",
	"synapse",
	"syndirella",
	"synthetic",
	"sysdate",
	"sysomos",
	"system(id)",
	"t-h-u-n-d-e-r-s-t-o-n-e",
	"t-online browser",
	"t-onlinebrowser",
	"tailrank",
	"targetapp",
	"taringa",
	"teamcity",
	"teleport",
	"telltale",
	"tendermint",
	"teoma",
	"test certificate info",
	"testagent",
	"testcertificateinfo",
	"testingus",
	"testtest",
	"the knowledge ai",
	"theknowledgeai",
	"thinklab",
	"threatview.app",
	"thumb",
	"timetravelaggregator",
	"tineye",
	"tiny tiny rss",
	"titan",
	"tmems",
	"tomthepeeper",
	"toolbar",
	"topic-clusters-ui",
	"topicaxis",
	"torrent",
	"traackr",
	"traackr.com",
	"tracemyfile",
	"transcoder",
	"transmission",
	"trello",
	"trendsmapresolver",
	"trove",
	"trustoo",
	"tumblr",
	"turnitin",
	"tutorial",
	"tweetedtimes",
	"twiceler",
	"twingly",
	"twingly recon",
	"twinglyrecon",
	"typhoeus",
	"tzzdev",
	"ubuntu apt-http",
	"ucmore",
	"ucsd",
	"ultraseek",
	"um-ln",
	"unbekannt",
	"undefined",
	"unknown",
	"unpaywall",
	"unshortenit",
	"upflow",
	"uptime",
	"urd-magpie",
	"ureq",
	"url",
	"url2file",
	"urlaliasbuilder",
	"urllib",
	"user-agent",
	"user.agent",
	"user_agent",
	"useragent",
	"utorrent",
	"uzbl",
	"vagabondo",
	"valavan_app",
	"valid",
	"validator",
	"vbseo",
	"vbulletin",
	"venus/fedoraplanet",
	"venusfedoraplanet",
	"viber",
	"viewport",
	"virtuoso",
	"virus.detector",
	"virustotal",
	"vkshare",
	"voila",
	"voltron",
	"voyager",
	"vpuls0u6s7",
	"vse",
	"vue-telescope",
	"w3af.org",
	"w3c",
	"w3m",
	"waitfor",
	"wakeletlinkexpander",
	"wanadoo",
	"wapchoi",
	"wappalyzer",
	"we7test",
	"weavr",
	"web app",
	"webbandit",
	"webchk",
	"webcloner",
	"webcollage",
	"webcopier",
	"webcorp",
	"webdatastats",
	"webexteams",
	"webglance",
	"webinator",
	"webkit2png",
	"weblayers",
	"webmetrics",
	"webmirror",
	"webmon",
	"weborama-fetcher",
	"webpage-inspector",
	"webreaper",
	"websitemetadataretriever",
	"webstripper",
	"webtech",
	"webthumbs",
	"webzip",
	"wechatgame",
	"wedocservice",
	"wfuzz",
	"wget",
	"whatcms",
	"whatever",
	"whatsapp",
	"whatweb",
	"whiteboard",
	"wii libnup",
	"willnorris",
	"win10chrome76",
	"windows-rss-platform",
	"windowscommunicationsapps",
	"winhttp",
	"withcabin",
	"wmtips.com",
	"woorankreview",
	"wordfence",
	"wordpress",
	"worm",
	"wp rocket",
	"wp-android-native",
	"wp_is_mobile",
	"wpai scheduler",
	"wpe wp upgrade",
	"www-mechanize",
	"www2pdf.de",
	"wxplr",
	"wxwork",
	"wyzo",
	"xenu",
	"xenu link sleuth",
	"xenulinksleuth",
	"xx032_bo9vs83_2a",
	"xymon",
	"yacy",
	"yahoo",
	"yandex",
	"yeti",
	"yoarcwhatsaps",
	"yodelta",
	"yqh24kes1y",
	"yst1upwya0rmlbx",
	"ywh-commander-crew",
	"zabbix",
	"zalopc",
	"zapier",
	"zdm",
	"zdmd",
	"zend_http_client",
	"zengenti",
	"zerodiumvar_dump",
	"zeus",
	"zeushdthree",
	"zgrab",
	"zjavascript",
	"zmeu",
	"zoho",
	"zoom.mac",
	"ztunnel",
	"zx-80 spectrum",
	"zxing",
	"zyborg",
	"破解后的",
	"脝脝陆芒潞贸碌脛",
}
