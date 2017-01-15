#!/bin/bash

set -eux

yum update -y

yum erase ruby20 ruby -y
yum install ruby23-devel ruby23 zlib zlib-devel -y
gem update --system
echo "export PATH=$PATH:/home/ec2-user/bin" >> /home/ec2-user/.bashrc
