# Muun

_A simple static site generator, written in Go, built with simple blogs in mind_

## Getting started

Muun can run in two modes, either it starts an http server and serves your content
for during development, or, builds your website and spits out flat html file in
the directory of your choice so you can then deploy to [Divshot](http://divshot.com/)
or [Github Pages](https://pages.github.com/).

**Server**

```bash
muun -serve -port 3000
```

**Build**

```bash
muun -out dist
```

## Configuring

Muun is also configurable using a config file in the root of your project directory

Here's an example config file with the options availble:

```yaml
# Example config file
# The values here are the defaults

# Http port to listen on when serving
port: 8080

# Directory in which to output built website
out: build

# Layout file location
layout: _layout.html

# Pages folder location
pages: pages

# Posts folder location
posts: posts
```

## Contributing

I you feel a feature is missing or find out a bug, feel totally free to submit
a pull request I will make myself a pleasure to read and attend to it, in other
words: contributions are welcomed this is not a pet project with a single selfish
owner.

## License

(MIT)

Copyright (c) 2015 Frederic Gingras (frederic@gingras.cc)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
