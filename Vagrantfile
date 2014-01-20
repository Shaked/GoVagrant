# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "centos65-x86_64"
  config.vm.box_url = "https://github.com/2creatives/vagrant-centos/releases/download/v6.5.1/centos65-x86_64-20131205.box"

  config.vm.network :forwarded_port, guest: 80, host: 8081
  config.vm.network :private_network, ip: "192.168.66.70"

  nfs = (RUBY_PLATFORM =~ /darwin/ || RUBY_PLATFORM =~ /linux/)
  config.vm.synced_folder ".", "/vagrant", :nfs => nfs

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "1536"] 
    vb.customize ["modifyvm", :id, "--ostype", "Centos65_64"]
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook = "provisioning/playbook.yml"
    ansible.verbose = "vvvv"
    ansible.inventory_path = "provisioning/development"
  end
end
