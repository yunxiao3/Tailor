#!/bin/sh

# 注意根据openstack环境实际情况配置

source keystonerc_admin

nova boot --flavor m1.large     \
--image centos7     \
--nic net-id=36601d90-86d1-4fc3-bfb8-ed90ee213cc6   \
--security-groups 9b81353a-824e-48a2-9e75-74b2f0c78433  \
--key-name my-key   \
tailor-client

nova boot --flavor m1.large     \
--image centos7     \
--nic net-id=36601d90-86d1-4fc3-bfb8-ed90ee213cc6   \
--security-groups 9b81353a-824e-48a2-9e75-74b2f0c78433  \
--key-name my-key   \
server001

nova boot --flavor m1.large     \
--image centos7     \
--nic net-id=36601d90-86d1-4fc3-bfb8-ed90ee213cc6   \
--security-groups 9b81353a-824e-48a2-9e75-74b2f0c78433  \
--key-name my-key   \
server002

nova boot --flavor m1.large     \
--image centos7     \
--nic net-id=36601d90-86d1-4fc3-bfb8-ed90ee213cc6   \
--security-groups 9b81353a-824e-48a2-9e75-74b2f0c78433  \
--key-name my-key   \
server003

sleep 5s

#neutron floatingip-list
openstack server add floating ip server001 192.168.6.43
openstack server add floating ip server002 192.168.6.37
openstack server add floating ip server003 192.168.6.48
openstack server add floating ip tailor-client 192.168.6.38


#delete server

# nova delete server001
# nova delete server002
# nova delete server003
