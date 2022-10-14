Readme for challenge.

Used language - Golang 1.17.

Notes:
 1. I first wrote code to get the json directly from provided URL but getting 
```<p>Your access is blocked due to the detection of a potential automated access request. The Bureau of Meteorology website does not support web scraping: if you are trying to access Bureau data through automated means, you should stop. You may like to consider the following options:</p>    <ul> <li>An anonymous FTP channel: <a href="http://www.bom.gov.au/catalogue/anon-ftp.shtml">http://www.bom.gov.au/catalogue/anon-ftp.shtml</a> -  this is free to access, but use is subject to the default terms of the Bureau's copyright notice: <a href="http://www.bom.gov.au/other/copyright.shtml">http://www.bom.gov.au/other/copyright.shtml</a> </li> <li>A Registered User service for continued use of Bureau data if your activity does not comply with the default terms: <a href="http://reg.bom.gov.au/other/charges.shtml">http://reg.bom.gov.au/other/charges.shtml</a> noting charges apply to most data products. Please contact webreg@bom.gov.au to discuss your requirements.</li> </ul>    <p>If you still need assistance in accessing our website, please contact us by filling in your details at <a href="http://reg.bom.gov.au/screenscraper/screenscraper_enquiry_form/">http://reg.bom.gov.au/screenscraper/screenscraper_enquiry_form/</a> and we will get in touch with you. Thank you for your understanding.</p> ```

2. So I downloaded the json and put in repo to use.
3. The code parses the json file and prints the output in json with `data` key as array of results in ascending order of apparent_t. If any error occurs it gives 503 as per given specification.
4. To check with different input change `data.json` in root of the repository and redeploy.
5. I used heroku sample heroku deployment code for Golang to deploy to https://murmuring-island-72312.herokuapp.com/
6. Dockerfile used is sample heroku golang dockerfile with small changes.
7. To deploy code use 
CI/CD:
```
heroku create
git push heroku main
heroku open
```
8. For tests use `go test`


Files:
main.go contains main application code
- The main logic is extractSortedData which extracts data with apparent_t greater than 15 and sorts it. 
main_test.go contains testing code.
I have written 4 tests.
First 2 are unit tests which tests basic functionality of algorithm
Rest 2 are end-to-end tests - commented out 1 which is to be tested by replacing json file with wrong file.

Not included exactly 15.0 apparent_t since it mentions greater than.


Solution to bonus question  - what is kafka? and key terminologies

Kafka is a message broker which has capabilites of storage, messaging and stream processing. 
Various terminologies are:
1. producer - one which produces message
2. consumer - one which consumes message
3. topic - similar to one queue.
4. partition - topic subdivisions.
5. consumer group - many consumers can together subscribe to one topic with at max 1 consumer per partition.
6. kafka lag - offset for each consumer on each partition.
7. retention_time - kafka retention time for message (usually 7 days).
