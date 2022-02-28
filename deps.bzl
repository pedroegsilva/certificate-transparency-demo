load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_anknown_ahocorasick",
        importpath = "github.com/anknown/ahocorasick",
        sum = "h1:onfun1RA+KcxaMk1lfrRnwCd1UUuOjJM/lri5eM1qMs=",
        version = "v0.0.0-20190904063843-d75dbd5169c0",
    )
    go_repository(
        name = "com_github_anknown_darts",
        importpath = "github.com/anknown/darts",
        sum = "h1:HblK3eJHq54yET63qPCTJnks3loDse5xRmmqHgHzwoI=",
        version = "v0.0.0-20151216065714-83ff685239e6",
    )
    go_repository(
        name = "com_github_calidog_certstream_go",
        importpath = "github.com/CaliDog/certstream-go",
        sum = "h1:P2kAob5k67YLNGJg0C6Wg0nTLrUpIxYvpPFVrk1y1PQ=",
        version = "v0.0.0-20200713031452-eca7997412f1",
    )
    go_repository(
        name = "com_github_cloudflare_ahocorasick",
        importpath = "github.com/cloudflare/ahocorasick",
        sum = "h1:8yL+85JpbwrIc6m+7N1iYrjn/22z68jwrTIBOJHNe4k=",
        version = "v0.0.0-20210425175752-730270c3e184",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:D9/bQk5vlXQFZ6Kwuu6zaiXJ9oTPe68++AzAJc1DzSI=",
        version = "v22.3.2",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:ZDRjVQ15GmhC3fiQ8ni8+OwkZQO4DARzQgrnXU1Liz8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
        version = "v5.0.4",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_jmoiron_jsonq",
        importpath = "github.com/jmoiron/jsonq",
        sum = "h1:ZZCvgaRDZg1gC9/1xrsgaJzQUCQgniKtw0xjWywWAOE=",
        version = "v0.0.0-20150511023944-e874b168d07e",
    )
    go_repository(
        name = "com_github_pedroegsilva_ahocorasick",
        importpath = "github.com/pedroegsilva/ahocorasick",
        sum = "h1:N5egH9vhDB1eGp00uZmi8OFzb3cFrE9dG5MfrjRKcGc=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_pedroegsilva_gofindthem",
        importpath = "github.com/pedroegsilva/gofindthem",
        sum = "h1:VzZT3J7w/8Kn3VEyTsImT8QRdlNw7614uJ4AbHQLg9A=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_pedroegsilva_gotagthem",
        importpath = "github.com/pedroegsilva/gotagthem",
        sum = "h1:ovCkf0CEbhsDKdF3UmAOH7JoqdB69a/XvxLsGjQpivQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:6NjYksEUlhurdVehpc7S7dk6DAmcKv8V9gG0FsVN2U4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        importpath = "github.com/rs/zerolog",
        sum = "h1:/ihwxqH+4z8UxyI70wM1z9yCvkWcfz/a3mj48k/Zngc=",
        version = "v1.26.1",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:nwc3DEeHmmLAfoZucVR881uASk0Mfjw8xYJ99tb5CcY=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:OtISOGfH6sOWa1/qXqqAiOIAO6Z5J3AEAE18WAq6BiQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
        version = "v0.0.0-20161208181325-20d25e280405",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:dUUwHk2QECo/6vqA44rthZ8ie2QXMNeKRTHCNY2nXvo=",
        version = "v3.0.0-20200313102051-9f266ea9e77c",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:1SzTfNOXwIS2oWiMF+6qu0OUDKb0dauo6MoDUQyu+yU=",
        version = "v0.0.0-20211215165025-cf75a172585e",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:Gz96sIWK3OalVv/I/qNygP42zyoKp3xptRVCWRFEBvo=",
        version = "v0.4.2",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:20cMwl2fHAzkJMEA+8J4JgqBQcQGzbisXo31MIeenXI=",
        version = "v0.0.0-20210805182204-aaa1db679c0d",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:5KslGYwFpkhGh+Q16bwMP3cOontH8FOep7tGV86Y7SQ=",
        version = "v0.0.0-20210220032951-036812b2e83c",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:WUoyKPm6nCo1BnNUvPGnFG3T5DUVem42yDJZZ4CNxMA=",
        version = "v0.0.0-20210809222454-d867a43fc93e",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
        version = "v0.0.0-20201126162022-7de9c90e9dd1",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:aRYxNxv6iGQlyVaZmk6ZgYEDa+Jg18DxebPSrd6bg1M=",
        version = "v0.3.6",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_extra_args = [
            "-exclude=**/testdata",
            "-exclude=go/packages/packagestest",
        ],
        importpath = "golang.org/x/tools",
        sum = "h1:6j8CgantCy3yc8JGBqkDLMKWqZ0RDU2g1HVgacojGWQ=",
        version = "v0.1.7",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
