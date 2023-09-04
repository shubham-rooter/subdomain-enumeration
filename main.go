package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
    "sync"
)

var (
    target      string
    outputfile  string
    verbose     bool
    githublink  string
    tools       = []string{"assetfinder", "amass", "findomain", "subfinder", "crtsh", "sublist3r", "subbrute"}
)

const (
    bold   = "\033[1m"
    normal = "\033[0m"
    green  = "\033[32m"
    yellow = "\033[33m"
    red    = "\033[31m"
)

func usage() {
    fmt.Printf("\n%s#########################################################%s\n", bold+green, normal)
    fmt.Printf("%s#                      - ************ -                  #%s\n", bold+yellow, normal)
    fmt.Printf("%s#              - * Rooter Subdomain Brutforce* -         #%s\n", bold+yellow, normal)
    fmt.Printf("%s#                      - ************ -                  #%s\n", bold+yellow, normal)
    fmt.Printf("%s#                      - by Shubham Rooter -             #%s\n", bold+yellow, normal)
    fmt.Printf("%s#########################################################%s\n\n", bold+green, normal)
    fmt.Printf("Usage: %s -t <target_domain> [-o <output_file>] [-v]\n", os.Args[0])
    fmt.Printf("%sOptions:%s\n", bold, normal)
    fmt.Println("  -t <target_domain>: Specify the target domain (required).")
    fmt.Println("  -o <output_file>: Specify the output file (default: subdomains.txt).")
    fmt.Println("  -v: Enable verbose mode (print tool names and progress messages).")
    fmt.Println("  -h: Display this help message.")
    fmt.Printf("\n%sExample:%s\n", bold, normal)
    fmt.Printf("  %s -t example.com -o results.txt -v\n\n", os.Args[0])
    fmt.Printf("\n%sDescription:%s\n", bold, normal)
    fmt.Println("This script performs subdomain enumeration on a target domain using various tools")
    fmt.Println("and saves the results in an output file.")
    fmt.Printf("\n%sGitHub:%s\n", bold, normal)
    fmt.Printf("GitHub Repository: %s\n", githublink)
    os.Exit(1)
}

func runTool(toolName string, command string) {
    if verbose {
        fmt.Printf("[%s*%s] Running %s%s%s...\n", bold, normal, bold, toolName, normal)
    }
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Printf("[%s-%s] %s%s%s failed.\n", bold, normal, bold, toolName, normal)
    } else {
        if verbose {
            fmt.Printf("[%s+%s] %s%s%s completed.\n", bold, normal, bold, toolName, normal)
        }
    }
}

func main() {
    githublink = "https://github.com/shubham-roote"
    flag.StringVar(&target, "t", "", "Specify the target domain (required).")
    flag.StringVar(&outputfile, "o", "subdomains.txt", "Specify the output file.")
    flag.BoolVar(&verbose, "v", false, "Enable verbose mode.")
    help := flag.Bool("h", false, "Display help message.")
    flag.Parse()

    if *help || target == "" {
        usage()
    }

    outputFile, err := os.Create(outputfile)
    if err != nil {
        log.Fatalf("[%s-%s] Unable to create output file: %v\n", bold, normal, err)
    }
    defer outputFile.Close()

    var wg sync.WaitGroup
    wg.Add(len(tools))

    for _, tool := range tools {
        switch tool {
        case "assetfinder":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("assetfinder --subs-only %s", target))
            }()
        case "amass":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("amass enum -timeout 2 -passive -d %s", target))
            }()
        case "findomain":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("findomain --quiet -t %s", target))
            }()
        case "subfinder":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("subfinder -d %s -o %s", target, outputfile))
            }()
        case "crtsh":
            go func() {
                defer wg.Done()
                response, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%%25.%s", target))
                if err != nil {
                    fmt.Printf("[%s-%s] Failed to fetch crt.sh data: %v\n", bold, normal, err)
                    return
                }
                defer response.Body.Close()

                reader := bufio.NewReader(response.Body)
                for {
                    line, err := reader.ReadString('\n')
                    if err == io.EOF {
                        break
                    }
                    if strings.Contains(line, target) {
                        subdomain := strings.TrimSpace(line)
                        outputFile.WriteString(subdomain + "\n")
                    }
                }
            }()
        case "sublist3r":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("sublist3r -d %s -o %s", target, outputfile))
            }()
        case "subbrute":
            go func() {
                defer wg.Done()
                runTool(tool, fmt.Sprintf("subbrute -t 5 %s", target))
            }()
        }
    }

    wg.Wait()

    // Deduplicate and sort the subdomains
    cmd := exec.Command("sort", "-u", "-o", outputfile, outputfile)
    if err := cmd.Run(); err != nil {
        fmt.Printf("[%s-%s] Error sorting subdomains: %v\n", bold, normal, err)
    }

    fmt.Printf("\n%s[*] Subdomain enumeration completed.%s Results saved to %s\n", bold+green, normal, outputfile)
}
