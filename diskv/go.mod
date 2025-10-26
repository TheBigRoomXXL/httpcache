module pkg.lovergne.dev/httpcache/diskv

go 1.24.9

replace pkg.lovergne.dev/httpcache => ../

require (
	github.com/peterbourgon/diskv/v3 v3.0.1
	pkg.lovergne.dev/httpcache v0.1.0-beta-2
)

require github.com/google/btree v1.0.0 // indirect
