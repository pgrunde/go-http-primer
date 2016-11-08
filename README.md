# go-http

## Get it Go-ing
###### Gopher puns!

Get that sweet [installation of Go](https://golang.org/doc/install), and bookmark [Effective Go](https://golang.org/doc/effective_go.html) for later reading and [A Tour of Go](https://tour.golang.org/welcome/1) for later learning. Also the official [Getting started with Go](https://github.com/golang/go/wiki#getting-started-with-go) section of the wiki is a good place to, well, get started.

## Gopath
###### Hey gopath yourself buddy!

The [Gopath](https://golang.org/doc/code.html#GOPATH) environment variable determines the path of your [workspace](https://golang.org/doc/code.html#Workspaces). Right out the gates I have to warn you- Go has an opinion about *everything*. This includes your directory structure, your linting, your formatting, your test file, you name it. Now if you're the kinda folk who will balk at some language waltzing in here and having the **gall** to tell you how to write your code, I implore you to reconsider. Go code is formatted and linted the same nearly everywhere, making code more familiar even when you've never seen it.

Let us [test that our installation of Go is correct](https://golang.org/doc/install#testing) with a quick and simple `hello, world`.

With the path setup and your workspace ready, be sure you understand what is happening in the `bin`, `pkg` and `src` folders- bin is like Unix bin in that it holds your executables, package is where you keep all your installed packages, and source is the source of your code.

## JavaScript differences

Alright you've probably already started to notice some of the differences between JS and Go. Lets talk about the differences.

Thing | Go | JS
--- | --- | ---
[Strongly Typed](https://en.wikipedia.org/wiki/Strong_and_weak_typing) | Yes | No
[Compiled](http://stackoverflow.com/questions/787239/what-is-a-dynamic-language-and-why-doesnt-c-sharp-qualify) | Yes | No
[Language Runs in Browser](http://www.wikihow.com/Change-Your-Browser's-Language) | No | Yes

###### Did I leave Go with the extra yes as pro-Golang propaganda? You decide, dear reader, but no.

## Simple Server

So one cool thing to know about Go is that is has an [awesome standard library of packages](https://golang.org/pkg/) and a built in [HTTP package](https://golang.org/pkg/net/http/) that includes its own [testing toolkit](https://golang.org/pkg/net/http/httptest/). That toolkit is used with Go's built in [testing framework](https://golang.org/pkg/testing/). Were going to build an example web app with that HTTP package.

### The first response

[Here on this commit](https://github.com/gSchool/go-http/tree/7865f5199c47bce852ff45f01050116796ed5153) we can see the beginning of a baby server, just born into the world and learning to respond to requests. Express users might recognize the format of an incoming request and an outgoing response. Babies might recognize some of the letters (if they're really bright and have good eyes). Either way, [congratulations on your first Go web app!](https://s-media-cache-ak0.pinimg.com/736x/e8/f6/85/e8f68586c05e9c608bf08efa1daeb752.jpg) Execute the program with `go run main.go`.

### Serve up files

Responding with a static template is all well and good, but what if I want to serve up some files? [Check out the directory of served files](https://github.com/gSchool/go-http/tree/6c0de34c4ec6538768679e4c680d82ab08427d53) we added here. Now I've got my HTML, my CSS, and my JavaScript all together. [Hot dog](http://www.kimballstock.com/pix/DOG/05/DOG_05_RK0084_05_P.JPG). 

### Insert your values here

Now [we've added templates as proper html files](https://github.com/gSchool/go-http/tree/37d86a62cf4bb8299dbd4f555c9e33f5f99d9e31). We're still just doing string interpolation, but now the templates live in their own happy home.

### What about those other URL values?

The incoming request has all the goodies you need to bag yourself some URL parameters [in this example](https://github.com/gSchool/go-http/tree/a000ec3b94f3d2fdf8064c1f063c2fe6cc3fbaf1). Now you can specify things like `?color=%234e73c0` or `?color=blue` if you're feeling retro. So far I'm only choosing to render my path and params, but you [REST](https://en.wikipedia.org/wiki/Representational_state_transfer) assured you can make all kinds of delicious decisions out of the combinations of those two things.


### Make Posts, Not War

Now I've gone and [shown you what a form post might look like](https://github.com/gSchool/go-http/tree/a5afcadb95994d04a8dcd943e20048ea4c7ceb20) (also notice I remembered to run [go fmt](https://blog.golang.org/go-fmt-your-code) this time). While I'm only having the server print out the form values, I bet you could imagine a website where they did other things. Think of something else you submit to a website? Which thing you submit is your faaaaavorite thing? I should mention that the pattern you're seeing here where I switch on the request method to determine actions taken is a very simple way to solve the muxing problem. There are tools out there which you can use to solve this problem more elegantly like Naoya Inada's [denco](https://github.com/naoina/denco) or Julien Schmidt's [httprouter](https://github.com/julienschmidt/httprouter).

### JavaScript Object Notation Nation

It is super practical to serve up information from a server that comes in a JSON payload, and [here we can see just that](https://github.com/gSchool/go-http/tree/4161b77a21ba6d3cd444ae26e56abe6ca7ab19ae). Structs and Maps are all hunky-dory when it comes to JSON marshalling. See if you can't swap out the `json.Marshal` call and replace it with a [Decoder](https://golang.org/pkg/encoding/json/#Decoder) type that decodes the JSON.

Well that's it for now- if you need a quick reference I've found [Go By Example](https://gobyexample.com/) to be an invaluable resource for remembering how to do things. 

