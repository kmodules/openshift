[![Go Report Card](https://goreportcard.com/badge/kmodules/openshift "Go Report Card")](https://goreportcard.com/report/kmodules/openshift)
[![GoDoc](https://godoc.org/kmodules.xyz/openshift?status.svg "GoDoc")](https://godoc.org/kmodules.xyz/openshift)
[![Build Status](https://travis-ci.org/kmodules/openshift.svg?branch=master)](https://travis-ci.org/kmodules/openshift)

# openshift

[OpenShift](https://www.openshift.com) comes with some custom API types, eg, DeploymentConfig and SecurityContextConstraints. To support these API types in AppsCode projects, we have to depend on the following additional projects.

- [github.com/openshift/api](https://github.com/openshift/api)
- [github.com/openshift/client-go](https://github.com/openshift/client-go)
- [github.com/openshift/origin](https://github.com/openshift/origin)

But these projects update their Kubernetes client-go dependency much slowly. It is too complicated to update their k8s.io/client-go dependency in our own fork. Also, the dependency tree of these libraries is just as huge as Kubernetes itself. So, we have "forked" just the API types and used Kubernetes code-generator to generate clients for these custom types.

Caveat:
- Keep the types.go in `apis` package up-to-date.
