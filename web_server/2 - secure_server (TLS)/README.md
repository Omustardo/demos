If you're setting up a web server, then the end goal is likely to serve it on a real website.

1. The first thing you'll need is a domain name, which you can buy on a variety of sites. I like https://domains.google.com 

2. Once you've obtained a domain, you need to direct it to your external IP address.
  * "Temporary Redirect" is fine. 
  * Make sure to enable "Forward Paths".
  * In the DNS settings you should enable DNSSEC.

3. At this point if you run the minimal server from the previous demo, you should be able to view it at your domain.

4. Next you need to get a certificate in order to have secure web traffic.
 
 * To get an initial proof of concept you can generate your own.
   * Create a private key with: `openssl genrsa -out server.key 2048`
   * Use the key to generate a certificate: `openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
   * Run the code in this demo after putting the key and crt files in the same directory as the executable
   * You should now be able to view your site at https://www.your_website.com/ although it will likely give a warning about an invalid certificate.

 * Real certificates can't be signed with your own key, they need to be signed by someone your browser is programmed to trust. Luckily this is straightforward and offered by many services. Many cost some money, so I recommend using letsencrypt to get a certificate for free. https://community.letsencrypt.org/t/list-of-client-implementations/2103
   * I decided to go with their browser implementation because it really guides you through it. Note that this requires 
     waiting until your website is available at http://your_website.com/ which may take 24-36 hours (note the lack of 
     'www' that appears to require some DNS propagation). It took about 2 hours when I did it.
    * https://zerossl.com/
    * Use the "FREE SSL Certificate Wizard"
    * Fill in Domain (for me it's "omustardo.com www.omustardo.com")
    * Accept terms of service, select "DNS Verification" and then click next. Save the Certificate Signing Request in a text file on your computer.
    * Click next again and save the RSA Private key somewhere on your computer. It's important never to share this. This key and the CSR will be used to renew your certificate periodically. Then click next again.
    * To verify ownership of the domain, go to domains.google.com and the DNS section. In the "Custom resource records" section you should add the TXT records that ZeroSSL asks for.
     * More info here: https://support.google.com/domains/answer/3290350?hl=en
    * Once you finish that step, you'll get a certificate and a key. Save them to server.crt and server.key files respectively (those names are just a convention and don't matter).
    * Also keep in mind that these certificates are only valid for 90 days so make sure to renew them regularly.
   * Finally modify your server to use the new certificate and private key. https://your_website.com should now work without any warning messages.
