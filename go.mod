module github.com/yoeria/goftxlogistics

go 1.18

require (
	github.com/charmbracelet/bubbles v0.10.3
	github.com/charmbracelet/bubbletea v0.20.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dustin/go-humanize v1.0.0
	github.com/go-echarts/go-echarts/v2 v2.2.4
	github.com/go-numb/go-ftx v0.0.0-20211208060907-82614a9fd25b
	github.com/go-redis/redis/v8 v8.11.4
	github.com/gookit/color v1.5.0
	github.com/joho/godotenv v1.4.0
	github.com/labstack/gommon v0.3.1
	github.com/markcheno/go-talib v0.0.0-20190307022042-cd53a9264d70
	github.com/sdcoffey/big v0.7.0
	github.com/sdcoffey/techan v0.12.1
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.26.0
)

require (
	github.com/andybalholm/brotli v1.0.0 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/klauspost/compress v1.10.7 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.11.1-0.20220212125758-44cd13922739 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.21.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/xo/terminfo v0.0.0-20210125001918-ca9a967f8778 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/term v0.0.0-20210422114643-f5beecf764ed // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/yoeria/goftxlogistics/cmd => ./cmd
	github.com/yoeria/goftxlogistics/protos => ./protos
	github.com/yoeria/goftxlogistics/strategies => ./strategies
)

exclude github.com/yoeria/goftxlogistics/protos v0.0.0
