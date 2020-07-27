# OpenFaaS and Chromedp


```sh
  # build
  faas-cli build -f faas-chrome.yml

  # deploy
  faas-cli deploy -f faas-chrome.yml

  # call
  curl -X POST http://127.0.0.1:8080/function/faas-chrome -d "http://www.eugenio.xyz"

  # output
  # <head><meta charset="utf-8"><meta http-equiv="cleartype" content="on"><meta 
  # name="HandheldFriendly" content="True"><meta name="description"...
```