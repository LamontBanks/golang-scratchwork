module example.com/hello

go 1.23.1

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0
)
