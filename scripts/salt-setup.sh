#!/bin/sh

#master
sudo yum install https://repo.saltstack.com/yum/redhat/salt-repo-latest.el7.noarch.rpm 
yum install salt-master -y
chkconfig salt-master on
service salt-master start
sudo yum install firewalld -y
systemctl stop firewalld

#minion1
sudo yum install https://repo.saltstack.com/yum/redhat/salt-repo-latest.el7.noarch.rpm -y 
yum install salt-minion -y
chkconfig salt-minion on
echo "master: 10.0.0.15" >> /etc/salt/minion
echo "id: 10.0.0.10" >> /etc/salt/minion
service salt-minion start
sudo yum install firewalld -y
systemctl stop firewalld

#minion2
sudo yum install https://repo.saltstack.com/yum/redhat/salt-repo-latest.el7.noarch.rpm 
yum install salt-minion -y
chkconfig salt-minion on
echo "master: 10.0.0.15" >> /etc/salt/minion
echo "id: 10.0.0.12" >> /etc/salt/minion
service salt-minion start
sudo yum install firewalld -y
systemctl stop firewalld


#minion3
sudo yum install https://repo.saltstack.com/yum/redhat/salt-repo-latest.el7.noarch.rpm 
yum install salt-minion -y
chkconfig salt-minion on
echo "master: 10.0.0.15" >> /etc/salt/minion
echo "id: 10.0.0.23" >> /etc/salt/minion
service salt-minion start
sudo yum install firewalld -y
systemctl stop firewalld
