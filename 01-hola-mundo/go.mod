module holamundo

go 1.22.1

// dependencias
require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.3.0 // indirect
)

require github.com/juanpabotero/greeting v0.0.0-00010101000000-000000000000

replace github.com/juanpabotero/greeting => ../greeting
