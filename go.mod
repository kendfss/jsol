module github.com/kendfss/jsol

go 1.18

replace (
	github.com/buger/jsonparser => ../../buger/jsonparser
	github.com/kendfss/but => ../but
	github.com/kendfss/iters => ../iters
	github.com/kendfss/mandy => ../mandy
	github.com/kendfss/oprs => ../oprs
	github.com/kendfss/oracle => ../oracle
	github.com/kendfss/pipe => ../pipe
	github.com/kendfss/rules => ../rules
	golang.org/x/exp => ../exp
)

require (
	github.com/buger/jsonparser v0.0.0-00010101000000-000000000000
	github.com/hokaccha/go-prettyjson v0.0.0-20211117102719-0474bc63780f
	github.com/kendfss/but v0.0.0-00010101000000-000000000000
	github.com/kendfss/mandy v0.0.0-00010101000000-000000000000
	github.com/kendfss/pipe v0.0.0-00010101000000-000000000000
	golang.design/x/clipboard v0.6.1
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/kendfss/iters v0.0.0-00010101000000-000000000000 // indirect
	github.com/kendfss/oprs v0.0.0-00010101000000-000000000000 // indirect
	github.com/kendfss/rules v0.0.0-00010101000000-000000000000 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/exp/shiny v0.0.0-20221002003631-540bb7301a08 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	golang.org/x/mobile v0.0.0-20210716004757-34ab1303b554 // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
)
