#!/bin/bash
#This script install all the packages needed for the Tito Front End (Apache, git)
#It also configure the necessary files
#it download the necessary sources from  Git
#and it start the service
exec &> tito_as.log

#### Variables
HTMLPATH=/var/www/html
GITREPO=https://github.com/Tito-org/tito-part4
HTTPDCONF=/etc/httpd/conf/httpd.conf
SQLSERVER=$1
CODEVERSION=V1
TitoCodePath=/tmp/tito-part4/Vince/TitoFE/

#### disable SE Linux
sed -i --follow-symlinks 's/^SELINUX=.*/SELINUX=disabled/g' /etc/sysconfig/selinux && cat /etc/sysconfig/selinux
setenforce 0

#### Disable firewall 
#echo -e "Open Firewall port 80\n"

#firewall-cmd --zone=public --add-port=80/tcp --permanent
#firewall-cmd --reload

#### Install and configire HTTPD
echo
echo -e "Install Apache HTTPD & PHP\n"

#yum update -y
yum install httpd -y
/usr/sbin/service httpd start
yum install php -y
yum install php-mysql -y
/usr/sbin/chkconfig httpd on



#### Download Tito code and configure HTTPD
git checkout tags/$CODEVERSION

echo
echo -e "Install Tito sources \n"
mv $TitoCodePath/* $HTMLPATH

echo
echo -e "modify SQLSERVER variable to remove not needed characters"
SQLSERVER=$(tr -d []\' <<< $SQLSERVER)
echo -e "SQLSERVER= " $SQLSERVER

echo
echo "LoadModule php5_module modules/libphp5.so" >> $HTTPDCONF
cat <<EOF >> $HTTPDCONF
<IfModule env_module>
    SetEnv TITO-SQL "$SQLSERVER"
</IfModule>
EOF

echo
echo -e "conf php.ini Timezone \n"

echo "date.timezone = \"Europe/Rome\"" >> /etc/php.ini

#### Start HTTPD
echo
echo -e "Restart Apache HTTPD"

/usr/sbin/service httpd restart
