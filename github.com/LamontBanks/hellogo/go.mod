module github.com/LamontBanks/hellogo

go 1.23.1

// The proper way to create and depend on modules is to publish them to a remote repository. 
// When you do that, the "replace" keyword can be dropped from the go.mod
replace github.com/LamontBanks/mystrings => ../mystrings

require (
    github.com/LamontBanks/mystrings v0.0.0
)