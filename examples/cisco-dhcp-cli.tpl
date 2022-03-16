ip dhcp pool host {{.hostname}}
address {{.ip}} {{.netmask}} hardware-address {{.macid}}
dns-server {{.dns1}} {{.dns2}}
