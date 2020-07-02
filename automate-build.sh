sudo /etc/init.d/mysql stop
sudo ./docker-nuke.sh
sudo docker-compose up -d --build