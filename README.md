Go Vagrant
=========

### Dependencies

- [Ansible](http://www.ansible.com)
- [Vagrant](vagrantup.com)

### Setup & Run

 Set up your /etc/hosts to point to

    192.168.66.70 go.dev

 Start Vagrant

     $ vagrant up

 Start the Go app

     $ ssh vagrant@go.dev 
     $ cd /vagrant/www/app 
     $ go run app.go 

 - Hit the url http://go.dev/app/   

### Supports 
 
- [gvm](https://github.com/moovweb/gvm)
- [ngnix](http://nginx.org/)

### References 

- http://mwholt.blogspot.nl/2013/05/writing-go-golang-web-app-with-nginx.html

#### Known Issues 

- Vagrant Mac [VBoxManage: error: Failed to create the host-only adapter](http://stackoverflow.com/questions/21069908/vboxmanage-error-failed-to-create-the-host-only-adapter)