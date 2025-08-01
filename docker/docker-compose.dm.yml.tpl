services:
  # https://hub.docker.com/_/mysql
  mysql1:
    hostname: mysql1
    networks:
      tiops:
        ipv4_address: __IPPREFIX__.201
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password --log-bin=/var/lib/mysql/mysql-bin.log --server-id=1 --binlog-format=ROW --gtid_mode=ON --enforce-gtid-consistency=true
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: ""
      MYSQL_ALLOW_EMPTY_PASSWORD: 1

  mysql2:
    hostname: mysql2
    networks:
      tiops:
        ipv4_address: __IPPREFIX__.202
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password --log-bin=/var/lib/mysql/mysql-bin.log --server-id=2 --binlog-format=ROW --gtid_mode=ON --enforce-gtid-consistency=true
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: ""
      MYSQL_ALLOW_EMPTY_PASSWORD: 1

  tidb1:
    hostname: tidb1
    networks:
      tiops:
        ipv4_address: __IPPREFIX__.211
    image: pingcap/tidb:v4.0.3
