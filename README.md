## Bug Bounty Script

Bug Bounty Script is a powerful and versatile Bash script designed to automate security testing tasks for bug bounty hunting. This script streamlines the process of reconnaissance, port scanning, vulnerability scanning, and more, helping security researchers and bug bounty hunters efficiently identify potential security vulnerabilities in target domains.


## Prerequisites

Rooter Subdomain Enumeration leverages the following subdomain enumeration tools:
- assetfinder
- amass
- findomain
- subfinder
- crt.sh
- sublist3r
- subbrute

Refer to the respective tool documentation for installation instructions.


### Key Features

- **Multi-Tool Integration:** Rooter Subdomain Enumeration seamlessly integrates with popular subdomain enumeration tools, including assetfinder, amass, findomain, subfinder, crt.sh, sublist3r, and subbrute. This comprehensive approach ensures a wide range of subdomains are identified.

- **User-Friendly Command-Line Interface:** The script offers an easy-to-use command-line interface, allowing users to quickly initiate subdomain enumeration tasks with minimal effort.

- **Verbose Mode:** For those who require detailed insights into the subdomain discovery process, Rooter Subdomain Enumeration provides a verbose mode that displays tool names and progress messages during execution.

- **Results in an Output File:** The script saves the discovered subdomains in an output file, streamlining subsequent analysis and enabling users to focus on potential attack vectors.

## Options

    -t <target_domain>: Specify the target domain (required).
    -o <output_file>: Specify the output file (default: subdomains.txt).
    -v: Enable verbose mode.



### Prerequisites

Before you begin, ensure you have met the following requirements:

- Go is installed on your system. If not, download and install it from the [official website](https://golang.org/dl/).


## Usage

1. Clone the repository:

   ```
    git clone https://github.com/shubham-rooter/subdomain-enumeration.git
	cd subdomain-enumeration
	go build
	sudo mv rooter-subdomain /usr/local/bin/
	rooter-subdomain -h
    ```
## RUN :   
 	rooter-subdomain -t example.com -o results.txt -v
  
Wait for the script to complete the subdomain enumeration process.

Once completed, the results will be saved in the specified output file.

## Contribution

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## Maintainers :

`This Repo is maintained by : `

- [Shubham Rooter](https://github.com/shubham-rooter)

### Author :

**Shubham Rooter**

* [Github](https://www.github.com/shubham-rooter)
* [Twitter](https://www.twitter.com/shubhamtiwari_r)
* [Instagram](https://www.instagram.com/shubham_rooter)
* [Linkdin](https://www.linkedin.com/in/shubham-tiwari09/)  

### License :

Copyright © 2023, [Shubham Rooter](https://github.com/Shubham-Rooter).
Released under the [MIT License](LICENSE).

***Thankyou.***
***Happy Hunting..***
