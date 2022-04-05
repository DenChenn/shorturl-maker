# Overview
[![Build Status](https://github.com/SwarzChen/shorturl-maker/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/SwarzChen/shorturl-maker/actions?query=branch%3Amain)
![GitHub](https://img.shields.io/github/license/SwarzChen/shorturl-maker)
![GitHub Repo stars](https://img.shields.io/github/stars/SwarzChen/shorturl-maker)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/SwarzChen/shorturl-maker)  

This is a project for applying my first intern job at **Dcard Taiwan Ltd.** 😃 !!!  
[🔗 Backend documentation link 🔗](https://documenter.getpostman.com/view/12176709/UVypycK7)

### Company Requirements

- [x] One POST api for uploading url 
- [x] One GET api for redirecting to original url
- [x] Use one of the relational databases: MySQL, PostgresSQL, SQLite
- [x] Use one of the cache database: Redis, Memcached
- [x] Reasonable error handling
- [x] No need to consider auth
- [x] Simultaneously user access handling
- [x] Non-existent shorten URL access handling

### Tech Stack
* Using **Golang Gin** framework to build api
* Using **postgresSQL** for relational database
* Using **redis** for caching database
* Deploy database and backend server on **Google Kubernetes Engine**
* **Github Actions** for CI / CD
* Implement **semantic versioning** with git

### Features
* Deploy backend service on **GKE 3-Nodes distributed systems**
* Deploy databases on **GKE 3-Nodes distributed systems**
* Handling invalid access and simultaneously access by **caching**
* Automatically **unit testing** in CI/CD workflow using github action
* Improve CI/CD efficiency with **pipeline workflow**
* Automatically **semantic versioning** in CI/CD workflow base on git label
* DNS and proxy server configuration using **cloudflare**

### Detail explanation
[How do I design my backend system architecture ?](https://medium.com/@aaaa102234/crazy-go-day-k8s-system-design-for-go-gin-redis-postgresql-957c74b4b25)  
[How do I handle simultaneously access and non-existed-url access with cache ?](https://medium.com/@aaaa102234/crazy-go-day-access-caching-go-gin-redis-58d0446e9a3a)  
[Why do I choose Gin framework and how do I handle error in it?](https://medium.com/@aaaa102234/crazy-go-day-why-using-gin-for-golang-backend-9ca48ec5d855)  
[How do I integrate both versioning process and unit tests into CI/CD process ?](https://medium.com/@aaaa102234/crazy-go-day-integrate-semantic-versioning-and-unit-tests-into-ci-cd-workflow-827d07495ca)  
[How do I implement unit tests ?](https://medium.com/@aaaa102234/crazy-go-day-simple-golang-unit-test-implementation-73518086496e)

### Future TODO list
- [ ] Build frontend UI with next.js and ts
- [ ] Feature: Upload video and image
- [ ] Feature: User can set password for uploaded resource


