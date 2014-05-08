# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "hashicorp/precise64"

  config.vm.provision "shell", inline: <<-EOS
    apt-get update
    apt-get install -y stow build-essential libxrender-dev fontconfig otf-ipafont-gothic otf-ipafont-mincho
    if [ ! `which wkhtmltox` ];then
      mkdir -p /usr/local/stow
      cd /usr/local/stow
      wget http://downloads.sourceforge.net/project/wkhtmltopdf/0.12.0/wkhtmltox-linux-amd64_0.12.0-03c001d.tar.xz
      tar xJf wkhtmltox-linux-amd64_0.12.0-03c001d.tar.xz
      stow wkhtmltox
    fi
  EOS
end
