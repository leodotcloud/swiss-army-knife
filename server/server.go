package server

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/leodotcloud/log"
)

const (
	// DefaultServerPort ...
	DefaultServerPort = "80"
	defaultServerPort = 80

)

type alphabet struct {
	Short string
	Long  string
}

var allAlphabets = []alphabet{
	{"A", "Alpha"},
	{"B", "Bravo"},
	{"C", "Charlie"},
	{"D", "Delta"},
	{"E", "Echo"},
	{"F", "Foxtrot"},
	{"G", "Golf"},
	{"H", "Hotel"},
	{"I", "India"},
	{"J", "Juliet"},
	{"K", "Kilo"},
	{"L", "Lima"},
	{"M", "Mike"},
	{"N", "November"},
	{"O", "Oscar"},
	{"P", "Papa"},
	{"Q", "Quebec"},
	{"R", "Romeo"},
	{"S", "Sierra"},
	{"T", "Tango"},
	{"U", "Uniform"},
	{"V", "Victor"},
	{"W", "Whiskey"},
	{"X", "X-ray"},
	{"Y", "Yankee"},
	{"Z", "Zulu"},
}

// TODO: Signal handler

var curlHomePageTmpl = `<html>
  <head>
    <title>Nato Alphabets</title>
  </head>
  <body>
    <div>
    <h1>Nato Alphabets</h1>
      {{- range $index, $alphabet := .Alphabets }}
      {{$alphabet.Short}} for {{$alphabet.Long}}
	  {{- end }}
    </div>
    <div>
      <h2>Server Info</h2>
        Hostname   : {{ .Hostname }}
        IP Address : {{ .IPAddress }}
		{{- if .DockerID }}
        Docker ID  : {{ .DockerID }}
		{{ end }}
		{{- if .RancherID }}
        Rancher ID : {{ .RancherID }}
		{{ end }}
    </div>
  </body>
</html>
`
var homePageTmpl = `<html>
  <head>
    <title>Nato Alphabets</title>
  </head>
  <body>
    <div>
    <h1>Nato Alphabets</h1>
      {{- range $index, $alphabet := .Alphabets }}
      <p>{{$alphabet.Short}} for {{$alphabet.Long}}</p>
	  {{- end }}
    </div>
    <div>
      <h2>Server Info</h2>
        <p><strong>Hostname   : </strong>{{ .Hostname }}</p>
        <p><strong>IP Address : </strong>{{ .IPAddress }}</p>
		{{- if .DockerID }}
        <p><strong>Docker ID  : </strong>{{ .DockerID }}</p>
		{{ end }}
		{{- if .RancherID }}
        <p><strong>Rancher ID : </strong>{{ .RancherID }}</p>
		{{ end }}
    </div>
  </body>
</html>
`

// Server ...
type Server struct {
	port        int
	exitCh      chan int
	l           net.Listener
	alphabets   []alphabet
}

// ErrorResponse ...
type ErrorResponse struct {
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// NewServer ...
func NewServer(portStr, inputAlphabet string) (*Server, error) {
	exitCh := make(chan int)

	return &Server{
		port:        getServerPortToRun(portStr),
		exitCh:      exitCh,
		alphabets:   getAlphabetsToUse(inputAlphabet),
	}, nil
}

// GetPort ...
func (s *Server) GetPort() int {
	return s.port
}

// GetExitChannel ...
func (s *Server) GetExitChannel() chan int {
	return s.exitCh
}

// Close ...
func (s *Server) Close() {
	s.l.Close()
}

// Run ...
func (s *Server) Run() error {
	log.Infof("Starting webserver on port: %v", s.port)
	http.HandleFunc("/", s.homePageHandler)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Errorf("error listening: %v", err)
		return err
	}
	s.l = l
	go func() {
		_ = http.Serve(l, nil)
	}()

	return nil
}

type homePageInfo struct {
	Alphabets []alphabet
	Hostname  string
	IPAddress string
	DockerID  string
	RancherID string
}

func (s *Server) homePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("r: %#v", r)

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	p := &homePageInfo{}
	p.Alphabets = s.alphabets

	hostname, err := os.Hostname()
	if err != nil {
		log.Errorf("Error getting hostname")
	} else {
		p.Hostname = hostname
	}

	localIP := GetLocalIP()
	log.Debugf("localIP: %v", localIP)
	p.IPAddress = localIP

	var t *template.Template

	ua := r.UserAgent()
	log.Debugf("ua: %v", ua)
	if strings.HasPrefix(ua, "curl") {
		t = template.Must(template.New("curlHomePageTmpl").Parse(curlHomePageTmpl))
	} else {
		t = template.Must(template.New("homePageTmpl").Parse(homePageTmpl))
	}
	err = t.Execute(w, p)
	if err != nil {
		log.Errorf("error parsing template: %v", err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}

// GetLocalIP ...
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getServerPortToRun(portStr string) int {

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Errorf("error parsing input port string: %v, using default port: %v", err, defaultServerPort)
		port = defaultServerPort
	}

	if port < 1 || port > 65535 {
		log.Errorf("invalid input port: %v, using default port: %v", port, defaultServerPort)
		port = defaultServerPort
	}
	return port
}

func getRandomAlphabetIndex() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(allAlphabets))
}

func getAlphabetIndex(alphabet string) int {
	index := rune(-1)
	if alphabet != "" {
		chIntIndex := rune(strings.ToUpper(alphabet)[0]) - 65
		if 0 <= chIntIndex && chIntIndex <= 25 {
			index = chIntIndex
		} else {
			log.Errorf("invalid alphabet specified, using all as default")
		}
	}
	return int(index)
}

func getAlphabetsToUse(inputAlphabet string) []alphabet {
	var alphabets []alphabet
	var alphabetIndex int

	if inputAlphabet == "random" || inputAlphabet == "RANDOM" {
		log.Infof("picking a random alphabet")
		alphabetIndex = getRandomAlphabetIndex()
	} else {
		alphabetIndex = getAlphabetIndex(inputAlphabet)
	}
	if alphabetIndex >= 0 {
		alphabets = allAlphabets[alphabetIndex : alphabetIndex+1]
		log.Infof("using alphabet: %v", alphabets)
	} else {
		alphabets = allAlphabets
		log.Infof("using all alphabets")
	}
	return alphabets
}
