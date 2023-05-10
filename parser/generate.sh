
#!/bin/sh

java -Xmx500M -cp "./antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -no-visitor -package parser *.g4