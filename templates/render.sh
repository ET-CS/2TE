#!/bin/bash

# compile
cheetah compile _layout --quiet
cheetah compile index --quiet

# run
python - <<EOF
import codecs

# compile cheetah
from index import index
#client = { 'surname': 'et', 'firstname': 'et', 'email': 'bla' }
#clients = [client]
#namespace = { 'title': 'Hello', 'contents': 'world', 'clients': clients }
namespace = { 'hello': 'Welcome to' }

# uncomment to debug cheetah
#print index(searchList=[namespace])
# save html version
html = str(index(searchList=[namespace]))

# uncomment to save non-minified version
#file = codecs.open("index.html", "w", "utf-8")
#file.write(html)
#file.close()

# minify using html_minify
from htmlmin.minify import html_minify
# uncomment to debug minified html
#print html_minify(str(index(searchList=[namespace])))
# save minified version
file = codecs.open("../index.min.html", "w", "utf-8")
file.write(html_minify(html))
file.close()
EOF

#python RUN.py

# uncomment to leave created cache files and final html(s) after done
# exit

#clean files
rm *.pyc
rm _layout.py
rm index.py
rm *.bak -f

# comment to remove created files after done
exit

# clean html
rm *.html