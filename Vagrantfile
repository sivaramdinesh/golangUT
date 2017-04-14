# -*- mode: ruby -*-
# vi: set ft=ruby :


#Bootstrap script for installing Go and setting correct environments
$bootstrapScript = <<SCRIPT
GO_VERSION="1.8"

echo "Updating packages ..."
sudo apt-get update

# Get the ARCH
ARCH=`uname -m | sed 's|i686|386|' | sed 's|x86_64|amd64|'`

# Install Go
echo "Downloading Golang 1.8"
cd /tmp
wget -q https://storage.googleapis.com/golang/go${GO_VERSION}.linux-${ARCH}.tar.gz
echo "Unpacking golang . . ."
sudo tar -xf go${GO_VERSION}.linux-${ARCH}.tar.gz -C /usr/local
rm go*.tar.gz

echo "Setting up go variables"
# Setup Go Path
SRCROOT="/usr/local/go"
SRCPATH="/home/vagrant/go"
mkdir -p $SRCPATH
mkdir -p $SRCPATH/bin
sudo chown -R vagrant:vagrant $SRCPATH

cat <<EOF >/tmp/gopath.sh
export GOPATH="$SRCPATH"
export GOBIN="$SRCPATH/bin"
export GOROOT="$SRCROOT"
export PATH="$SRCROOT/bin:$SRCPATH/bin:\$PATH"

EOF
sudo mv /tmp/gopath.sh /etc/profile.d/gopath.sh
sudo chmod 0755 /etc/profile.d/gopath.sh
source /etc/profile.d/gopath.sh
SCRIPT

required_plugins = %w(vagrant-cachier)

required_plugins.each do |plugin|
  need_restart = false
  unless Vagrant.has_plugin? plugin
    system "vagrant plugin install #{plugin}"
    need_restart = true
  end
  exec "vagrant #{ARGV.join(' ')}" if need_restart
end


# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  vmName = "golang-UT"
  # Using Ubuntu 16.04
  config.vm.box = "bento/ubuntu-16.04"

 # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
   if Vagrant.has_plugin?("vagrant-vbguest") then
    config.vbguest.auto_update = false
  end

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  config.vm.provider "virtualbox" do |vb|
      #vb.gui = true
      vb.customize ["modifyvm", :id, "--cableconnected1", "on"]
  end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Define a Vagrant Push strategy for pushing to Atlas. Other push strategies
  # such as FTP and Heroku are also available. See the documentation at
  # https://docs.vagrantup.com/v2/push/atlas.html for more information.
  # config.push.define "atlas" do |push|
  #   push.app = "YOUR_ATLAS_USERNAME/YOUR_APPLICATION_NAME"
  # end

  config.vm.define vmName do |vmCfg|
      vmCfg.vm.hostname = vmName
      #Adding Vagrant-cachier
      if Vagrant.has_plugin?("vagrant-cachier")
        vmCfg.cache.scope = :machine
        vmCfg.cache.enable :apt
        vmCfg.cache.enable :gem
      end
  end


  #Calling bootstrap setup
  config.vm.provision "shell", privileged: false, inline: $bootstrapScript
end