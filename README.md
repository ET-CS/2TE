2TE
===

Two Template Engines (in one app) usage example - Cheeta (Python) / SCSS / JS / Golang server

This is a VERY simple/basic/skeleton example of how I work with template
engine in my web development.

[Read more](http://itekblog.com/template-engine-need-two/) about why I use two template engines in one app and about this example.

## Requirements

* Golang

if you want to make changes and render the .tmpl files or .scss files you need:
* SCSS for CSS (Ruby/Sass/Compass required)
* Python + Cheetah templating engine

# Two Template Engines Layers

* First layer - Python - Cheetah
* Second Layer - Golang - html/template

## How to use

run server using: go run main.go and browse http://127.0.0.1:8080 to see result.
if you don't have golang installed you can run the supplied main binary instead

## Edit HTML/CSS/JS

If you want to make changes to the .scss files you need Ruby/Sass/Compass

If you want to make changes to the .tmpl files in the /templates folder 
use the render.sh file to render the updated index.min.html file.
you must run the server again as the min.html cached in memory for performance.

If you want to make changes to the JS you'll need minifyjs

## Deploy
In this example, to deploy you need only the binary and the .min.html files.
(2 files in this example)

## Watch
I've supplied a watch.sh file I love to use inside the templates folder. 
It used for watch for every file change in the directory and rebuild it in 
auto.
 inotifywait required.

## Why all this complexity?
* Performance - Golang caches one minified html with css, js included. better for server/client side performance.
* Cheetah is powerful! you can create a complex inheritance for your files. AND you can use the fastest template engine you'll find in the second layer (like golangs html/template in this example).
* You get to write and split code (html, css and js) in many files as you like. you don't care of the final minified file.
* You can put comments wherever you want freely!! (in the html as well as in the css/js) and they won't be there at the client side (they're removed at the minify process)
* You can use Placeholders as usual in the webserver template engine, as well in the first layer - in the cheetah. for example: application name. It doesn't supposed to be changed so it can be rendered in the first layer.
* JS files are minified and verified in the proccess. you will know about missing ; soon then ever!
* SASS is so sweet... at last you can manage your CSS right!