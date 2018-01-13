The most basic server you can do in golang.

`go run server.go`

Then go to http://localhost:80 and it should say 'Hello World!'

Note that http://localhost:80 is equivalent to http://localhost because port 80 is the default for http.


If you want to be able to view this website from the outside internet rather than just your own computer
(or local network), you'll need to set up port forwarding on your router to whichever port you're serving from.
On my router it's in the Firewall/NAT settings. I added two rules:

  * Original_Port:80 Protocol:Both Forward-to_Address:192.168.1.50 Forward-to_Port:80
  * Original_Port:443 Protocol:Both Forward-to_Address:192.168.1.50 Forward-to_Port:443

The Forward-to address is the local address of the device you are serving from. If your router supports it, you could make this set statically rather than the default DHCP - that way it will never happen to change on you. Otherwise don't worry about it, but be aware you will need to change the port forwarding settings if your IP changes.
  
Once port forwarding is set up you should find your external IP address (https://www.google.com/search?q=my%20ip)
and then can go to http://<your_ip>:<port_that_your_server_is_serving_at>
