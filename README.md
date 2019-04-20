# urlshorten

---
tags: Golang
---
# Week 1.3 Assignment: URL Shortener

## Assignment details

The goal of this assignment is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.


## Requirements

The list of redirection should be maintained in a command line tool, what can:
- [x] Manipulate YAML config file. Where the redirection list peristently stored.
- [x] Implement append to the list: `urlshorten configure -a dogs -u www.dogs.com` 
- [x] Implement remove from the list: `urlshorten -d dogs`
- [x] List redirections: `urlshorten -l`
- [x] Run HTTP server on a given port: `urlshorten run -p 8080`
- [x] Prints usage info: `urlshorten -h`


## Bonus

As a bonus exercises you can also...

- [ ] Track number of times each redirection is used. When the user uses `urlshorten -l`, the user should see redirections ranked by how many times they have been used.
- [ ] Provide a default shortening, if no example is given. For example, if `dogs` is not given, generate a random alphanumeric string of length 8.
- [ ] Build a Handler that doesn't read from a map but instead reads from a database. Whether you use BoltDB, SQL, or something else is entirely up to you.