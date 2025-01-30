# Dommie v0.0.4
A TinyGo DOM library I use for a lot of projects. It's very basic, so probably
not what you are looking for.

__Caution: This is still very much WIP__

## Demo
Yup, there's one at <https://ewaldhorn.github.io/dommie/>

## Why though
As a contractor, I do a lot of different things for clients. Of late, I've been
building a lot of dashboards and data visualisations etc. The web is just a great
fit for much of this, but I also enjoy using Go. So Dommie helps me target the
web, but also use mostly the same language I use to work on the API and server side
of things.

Less context-switching for me means I feel more productive. Whether that's really
the case or not is irrelevant to me, I love what I'm doing and this is fun. It's
a big win that clients are happy with what I deliver though!


## How
This uses [TinyGo](https://tinygo.org/) to build small wasm objects that you can
then use in the browser. I often write wasm programs for clients where I need to
interact with the DOM in the browser from Go.  So I decided to combine all the
different little bits and pieces of code I have into a single library so that I
can maintain it in just one place.
