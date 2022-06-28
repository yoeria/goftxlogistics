module github.com/yoeria/goftxlogistics

go 1.18

require (
	github.com/charmbracelet/bubbles v0.10.3
	github.com/charmbracelet/bubbletea v0.20.0
	github.com/charmbracelet/lipgloss v0.4.0
	github.com/dustin/go-humanize v1.0.0
	github.com/go-echarts/go-echarts/v2 v2.2.4
	github.com/go-numb/go-ftx v0.0.0-20211208060907-82614a9fd25b
	github.com/go-redis/redis/v8 v8.11.4
	github.com/joho/godotenv v1.4.0
	github.com/json-iterator/go v1.1.12
	github.com/labstack/gommon v0.3.1
	github.com/markcheno/go-talib v0.0.0-20190307022042-cd53a9264d70
	github.com/sdcoffey/big v0.7.0
	github.com/sdcoffey/techan v0.12.1
	github.com/spf13/cobra v1.5.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/andybalholm/brotli v1.0.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.10.7 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/muesli/ansi v0.0.0-20211018074035-2e021307bc4b // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.11.1-0.20220212125758-44cd13922739 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/sahilm/fuzzy v0.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.21.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/term v0.0.0-20210422114643-f5beecf764ed // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace (
	github.com/yoeria/goftxlogistics/cmd => ./cmd
	github.com/yoeria/goftxlogistics/protos => ./protos
	github.com/yoeria/goftxlogistics/strategies => ./strategies
)

exclude github.com/yoeria/goftxlogistics/protos v0.0.0
