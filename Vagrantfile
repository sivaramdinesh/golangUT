$installScript = <<SCRIPT
echo "Downloading anf Settingup golang"
sudo apt-get install -y golang-go
SCRIPT
VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "bento/ubuntu-16.04"
  config.vm.provision "shell", privileged: false, inline: $installScript
end
