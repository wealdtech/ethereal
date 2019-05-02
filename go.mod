module github.com/wealdtech/ethereal

go 1.12

// 1.8.21 broke return values from contract calling.  Can remove this when 1.9 is released
replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.8.20

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/antlr/antlr4 v0.0.0-20190223165740-dade65a895c2
	github.com/ethereum/go-ethereum v1.8.27
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/miekg/dns v1.1.4
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.1 // indirect
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829 // indirect
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90 // indirect
	github.com/prometheus/procfs v0.0.0-20190306233201-d0f344d83b0c // indirect
	github.com/prometheus/prometheus v2.1.0+incompatible // indirect
	github.com/prometheus/tsdb v0.7.0 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.3.1
	github.com/stretchr/testify v1.3.0
	github.com/wealdtech/go-ens/v2 v2.0.12
	github.com/wealdtech/go-erc1820 v1.2.0
	github.com/wealdtech/go-string2eth v1.0.0
	golang.org/x/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 // indirect
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)
