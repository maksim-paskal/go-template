# Motivation
Modern dynamic infrastructure needs some config files that use environment variables, sometimes applications allow to use environment variable in configuration - some time no. This tool is single binnary application that pretty thin ~8Mb, its reduce docker image size, it's use golang templating engine for text processing

# Example of usage
```bash
cat << EOF > /tmp/go-template-example
{{ env "TEST_ENV" }}
EOF

# creates rendered file 
cat /tmp/go-template-example | TEST_ENV=test go-template > /tmp/go-template-example.txt

# check results
cat /tmp/go-template-example.txt

# for example it can be used for render all files in directory with .template file extension
for file in $(find /etc/nginx/conf.d/ /etc/nginx -name *.template -type f)
do
  dir=$(dirname $file)
  name=$(basename $file .template)
  cat $file | go-template > $dir/$name
done
```

# Installation
MacOS
```bash
brew install maksim-paskal/tap/go-template
```
Linux
```bash
curl -L https://github.com/maksim-paskal/go-template/releases/download/v0.0.8/go-template-linux-amd64 -o /usr/local/bin/go-template
chmod +x /usr/local/bin/go-template
```
Dockerfile
```
ADD https://github.com/maksim-paskal/go-template/releases/download/v0.0.8/go-template-linux-amd64 /usr/local/bin/go-template

RUN chmod +x /usr/local/bin/go-template
```