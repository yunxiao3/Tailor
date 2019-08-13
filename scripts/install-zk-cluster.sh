#!/bin/bash
# 参考 https://www.cnblogs.com/woshimrf/p/zk-install.html

salt '*' cmd.run 'wget http://good.ncu.edu.cn/mirrors/jdk/jdk-8u161-linux-x64.tar.gz'
salt '*' cmd.run 'tar -zxvf /root/jdk-8u161-linux-x64.tar.gz -C /opt'

salt '*' cmd.run "echo 'export JAVA_HOME=/opt/jdk1.8.0_161' >> /etc/profile"
salt '*' cmd.run "echo 'export PATH=\$PATH:\$JAVA_HOME/bin' >> /etc/profile"
salt '*' cmd.run "echo 'export CLASSPATH=\$JAVA_HOME/jre/lib/ext:\$JAVA_HOME/lib/tools.jar' >> /etc/profile"


salt '*' cmd.run "source /etc/profile" runas='root' shell='/bin/bash'
salt '*' cmd.run "java --version"

salt '*' cmd.run 'wget http://good.ncu.edu.cn/mirrors/zookeeper-3.4.6.tar.gz'
salt '*' cmd.run 'tar -zxvf /root/zookeeper-3.4.6.tar.gz -C /opt'

salt '*' cmd.run 'cp /opt/zookeeper-3.4.6/conf/zoo_sample.cfg /opt/zookeeper-3.4.6/conf/zoo.cfg'
