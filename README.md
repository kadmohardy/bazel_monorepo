
Bazel Golang Monorepo
============

This is a Golang monorepo application example using Bazel (Blzmod)

![Chat Preview](http://i.imgur.com/lgRe8z4.png)

## Setup
It's important to have Bazel installed. For that, follow link:
`https://bazel.build/install?hl=pt-br`

---

#### Usage:
- **Build all projects projects:**  `bazel build //...`
- **Update dependencies with Gazelle:** `bazel run //:gazelle`
- **Run a single app A:** `bazel run //apps/app_a:app_a`

---
### References
https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE
https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/?couponCode=A2836A32A88CD55DED0E
https://www.youtube.com/watch?v=m6nSsCIb0iQ
https://github.com/techschool/simplebank
https://earthly.dev/blog/golang-monorepo/
https://github.com/jankremlacek/go-bazel/blob/main/services/servicea/BUILD.bazel
https://earthly.dev/blog/monorepo-with-bazel/
https://github.com/wissensalt/golang-bazel-monorepo
