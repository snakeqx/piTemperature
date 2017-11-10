# piTemperature

## dependency:

* Golang (only compiled in 1.9.1, others unknown, but it should work)
* beego
* mysql

## installation:
### install mysql
* sudo apt install mysql
* sudo mysql
* mysql> create database various;
* mysql> insert into mysql.user(Host, User, Password) values("localhost","defaultuser","default");
* mysql> flush privileges;
* mysql> grant all privileges on various.* to defaultuser@localhost identified by "default";
* mysql> flush privileges;

### clone the code and compile
* go get -u github.com/snakeqx/piTemperature
* cd $GOPATH/src/github.com/snakeqx/piTemperature
* go build


### run
* honup $GOPATH/src/github.com/snakeqx/piTemperature&
* and visit http://localhost:8080 from your browser
* The programe will automatically load cpu/gpu temperature in raspberry pi.




