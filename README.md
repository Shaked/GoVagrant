Go Vagrant
=========

### Dependencies

- [Ansible](http://www.ansible.com)
- [Vagrant](vagrantup.com)

### Setup & Run

 - Set up your /etc/hosts to point to


    192.168.66.70 go.dev

 - Start Vagrant


     $ vagrant up

 - Start the Go app


     $ ssh vagrant@go.dev 
     $ cd /vagrant/www/app 
     $ go run app.go 

 - Hit the url http://go.dev/app/   

### Supports 
 
- [gvm](https://github.com/moovweb/gvm)
- [ngnix](http://nginx.org/)


### Refrerences 

- http://mwholt.blogspot.nl/2013/05/writing-go-golang-web-app-with-nginx.html

