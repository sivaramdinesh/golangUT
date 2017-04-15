#The script is assigned to a global variable $installScript. 
#This global variable contains a string which is then passed in as the inline script to the Vagrant configuration.
$installScript = <<SCRIPT
sudo apt-get install -y golang-go
SCRIPT
VAGRANTFILE_API_VERSION = "2"
Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
# Every Vagrant development environment requires a box.
  config.vm.box = "bento/ubuntu-16.04"
# Disable automatic box update checking.
  config.vm.box_check_update = false
# Enable provisioning with a shell script. 
 #For more about provision visit: https://www.vagrantup.com/docs/provisioning/shell.html
  config.vm.provision "shell", privileged: true inline: $installScript
end
