module github.com/thoughtworks/maeve-csms/gateway

go 1.20

require (
	github.com/eclipse/paho.golang v0.11.0
	github.com/go-chi/chi/v5 v5.0.8
	github.com/google/go-cmp v0.5.9
	github.com/mochi-co/mqtt/v2 v2.2.11
	github.com/prometheus/client_golang v1.15.1
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.2
	go.uber.org/goleak v1.2.1
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/rs/xid v1.4.0 // indirect
	github.com/rs/zerolog v1.28.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// upstream version is leaking go routines - fixed in fork, need to submit PR
replace github.com/eclipse/paho.golang v0.11.0 => github.com/subnova/paho.golang v0.0.0-20230606110013-87b4fea2a216
