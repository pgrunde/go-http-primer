# go-http

### Get it Go-ing
###### Gopher puns!

Get that sweet [installation of Go](https://golang.org/doc/install), and bookmark [Effective Go](https://golang.org/doc/effective_go.html) for later reading and [A Tour of Go](https://tour.golang.org/welcome/1) for later learning. Also the official [Getting started with Go](https://github.com/golang/go/wiki#getting-started-with-go) section of the wiki is a good place to, well, get started.

### Gopath
###### Hey gopath yourself buddy!

The [Gopath](https://golang.org/doc/code.html#GOPATH) environment variable determines the path of your [workspace](https://golang.org/doc/code.html#Workspaces). Right out the gates I have to warn you- Go has an opinion about *everything*. This includes your directory structure, your linting, your formatting, your test file, you name it. Now if you're the kinda folk who will balk at some language waltzing in here and having the **gall** to tell you how to write your code, I implore you to reconsider. Go code is formatted and linted the same nearly everywhere, making code more familiar even when you've never seen it.

Let us [test that our installation of Go is correct](https://golang.org/doc/install#testing) with a quick and simple `hello, world`.

With the path setup and your workspace ready, be sure you understand what is happening in the `bin`, `pkg` and `src` folders- bin is like Unix bin in that it holds your executables, package is where you keep all your installed packages, and source is the source of your code.

### JavaScript differences

Alright you've probably already started to notice some of the differences between JS and Go. Lets talk about the differences.

Thing | Go | JS
--- | --- | ---
[Strongly Typed](https://en.wikipedia.org/wiki/Strong_and_weak_typing) | :heavy_check_mark: | :heavy_multiplication_x:
[Compiled](http://stackoverflow.com/questions/787239/what-is-a-dynamic-language-and-why-doesnt-c-sharp-qualify) | :heavy_check_mark: | :heavy_multiplication_x:
[Language Runs in Browser](http://www.wikihow.com/Change-Your-Browser's-Language) | :heavy_multiplication_x: | :heavy_check_mark:


Simple Server