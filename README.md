# README

Setup this program using below simple steps:
1. Clone this repo to your server.
`git clone https://github.com/irvanherz/golang-test`
then `cd` to the cloned folder
`cd golang-test`
2. Build docker container
`docker build . -t go-dock`
3. Run container
`docker run -p 5000:5000 go-dock`
  ```