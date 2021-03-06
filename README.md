# CLI

Simple command line interfaces.

This is a fork from: https://github.com/robmerrell/comandante 

[Example App](http://github.com/robmerrell/comandante_example)

## What is CLI

Many command line tools are structured to accept subcommands. Like the go tool. When you run the go command you are presented with help 
text listing the available subcommands. Each subcommand has its own command line flags and documentation. CLI makes creating binaries with subcommands easy.

So what does CLI look like? Let's say that you are creating a program called "coolbinary" with two subcommands: "sayhi" and "demo". 
If you were to run coolbinary without a subcommand 

```bash
$ coolbinary
```

you would get this output:

```
Do cool things from the commandline

Usage:
	coolbinary command [arguments]

Available commands:
demo   run a demo
help   get more information about a command
sayhi  say hello

Use "coolbinary help [command]" for more information about a command.
```

And the (over)simplified code that handles the subcommands:

```go
bin := cli.New("coolbinary", "Do cool things from the commandline")
bin.IncludeHelp() // can create a "help" command for you

greetCmd := cli.NewCommand("sayhi", "say hello", helloFunction)
greetCmd.Documentation = "This is longer form documentation for the sayhi command"
bin.RegisterCommand(greetCmd)

demoCmd := cli.NewCommand("demo", "Run a demo", DemoFunction)
demoCmd.Documentation = "longer form documentation of the demo command"
bin.RegisterCommand(demoCmd)

// run the command
if err := bin.Run(); err != nil {
	fmt.Fprintln(os.Stderr, err)
}
```

If you were to invoke the `sayhi` subcommand

```bash
$ coolbinary sayhi
```

`helloFunction` would be called

## Installation

Install to your `GOPATH` with

```bash
$ go get github.com/lucasepe/cli
```

## Usage

First step is to generate a new cli with the name of your program and a short description:

```go
bin := cli.New("coolbinary", "Do cool things from the commandline")
```

We also want to include the autogenerated help command:

```go
bin.IncludeHelp()
```

Now we add subcommands to that cli instance. Creating a new subcommand takes a name and a short description.

```go
greetCmd := cli.NewCommand("sayhi", "say hello", helloFunction)
bin.RegisterCommand(greetCmd)
```

We also should give that subcommand a longer description for the autogenerated help:

```go
greetCmd.Documentation = `
This is longer form documentation that describes this command in much greater detail.
`
```

Now register the subcommand:

```go
bin.RegisterCommand(greetCmd)
```

## Command line options

Subcommands can also handle their own command line options via FlagInit. These command line options
are automatically displayed in the command's help text.

```go
var testFlag string
greeterCmd.FlagInit = func(fs *flag.FlagSet) {
	fs.StringVar(&testFlag, "testing", "", "This is the usage")
}
```

