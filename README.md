```
$ copilot init

Welcome to the Copilot CLI! We're going to walk you through some questions
to help you get set up with a containerized application on AWS. An application is a collection of
containerized services that operate together.

Workload type: Request-Driven Web Service
Service name: uuid
Manifest file for job uuid already exists. Skipping configuration.
Ok great, we'll set up a Request-Driven Web Service named uuid in application copilot-apprunner-go.

✔ Created the infrastructure to manage services and jobs under application copilot-apprunner-go..

✔ The directory copilot will hold service manifests for application copilot-apprunner-go.

✔ Manifest file for service uuid already exists at copilot/uuid/manifest.yml, skipping writing it.
Your manifest contains configurations like your container size and port.

✔ Created ECR repositories for service uuid..

All right, you're all set for local development.
Deploy: Yes

✔ Linked account 573423333442 and region ap-northeast-1 to application copilot-apprunner-go..

✔ Proposing infrastructure changes for the copilot-apprunner-go-test environment.
✔ Created environment test in region ap-northeast-1 under application copilot-apprunner-go.
Environment test is already on the latest version v1.7.0, skip upgrade.
Building your container image: docker build -t 573423333442.dkr.ecr.ap-northeast-1.amazonaws.com/copilot-apprunner-go/uuid --platform linux/amd64 /Users/ryokosuge/Workspace/go/src/github.com/ryokosuge/aws-copilot-apprunner-go -f /Users/ryokosuge/Workspace/go/src/github.com/ryokosuge/aws-copilot-apprunner-go/Dockerfile
[+] Building 3.4s (18/18) FINISHED
 => [internal] load build definition from Dockerfile                                                                                                                                      0.0s
 => => transferring dockerfile: 37B                                                                                                                                                       0.0s
 => [internal] load .dockerignore                                                                                                                                                         0.0s
 => => transferring context: 2B                                                                                                                                                           0.0s
 => [internal] load metadata for docker.io/library/alpine:3                                                                                                                               2.1s
 => [internal] load metadata for docker.io/library/golang:1.17-buster                                                                                                                     3.1s
 => [bulder 1/7] FROM docker.io/library/golang:1.17-buster@sha256:ad61e922fb389384f00ba98dba8dae56f1dcf9ff731506857ab8f1312c47c212                                                        0.0s
 => [internal] load build context                                                                                                                                                         0.0s
 => => transferring context: 138B                                                                                                                                                         0.0s
 => [stage-1 1/5] FROM docker.io/library/alpine:3@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300                                                                 0.0s
 => CACHED [stage-1 2/5] WORKDIR /app                                                                                                                                                     0.0s
 => CACHED [bulder 2/7] WORKDIR /app                                                                                                                                                      0.0s
 => CACHED [bulder 3/7] COPY go.mod ./                                                                                                                                                    0.0s
 => CACHED [bulder 4/7] COPY go.sum ./                                                                                                                                                    0.0s
 => CACHED [bulder 5/7] RUN go mod download                                                                                                                                               0.0s
 => CACHED [bulder 6/7] COPY cmd/app/main.go ./                                                                                                                                           0.0s
 => CACHED [bulder 7/7] RUN go clean -cache && go build -o /server                                                                                                                        0.0s
 => CACHED [stage-1 3/5] COPY --from=bulder /server /app/server                                                                                                                           0.0s
 => CACHED [stage-1 4/5] RUN set -eux &&  mkdir /lib64 &&   ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2                                                                  0.0s
 => CACHED [stage-1 5/5] RUN set -eux &&  addgroup app &&  adduser -D -G app app &&  chown -R app:app /app &&  chmod 500 /app/server                                                      0.0s
 => exporting to image                                                                                                                                                                    0.0s
 => => exporting layers                                                                                                                                                                   0.0s
 => => writing image sha256:ec62cb328666fbdfd8dc85961054d72821b8facd7d9ed9dd8f1da9c80bc6a36e                                                                                              0.0s
 => => naming to 573423333442.dkr.ecr.ap-northeast-1.amazonaws.com/copilot-apprunner-go/uuid                                                                                              0.0s

Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
WARNING! Your password will be stored unencrypted in /Users/ryokosuge/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
Using default tag: latest
The push refers to repository [573423333442.dkr.ecr.ap-northeast-1.amazonaws.com/copilot-apprunner-go/uuid]
c30d19dbf562: Layer already exists
93ea72d83254: Layer already exists
8997b4223178: Layer already exists
a3d6d38369cf: Layer already exists
8d3ac3489996: Layer already exists
latest: digest: sha256:1f354ef454406e7a2503bdf32cec1dcd4185200727c47f558dd8fd91bdd85a9c size: 1364
✔ Proposing infrastructure changes for stack copilot-apprunner-go-test-uuid
- Creating the infrastructure for stack copilot-apprunner-go-test-uuid            [create complete]  [325.3s]
  - An IAM Role for App Runner to use on your behalf to pull your image from ECR  [create complete]  [25.4s]
  - An IAM role to control permissions for the containers in your service         [create complete]  [25.4s]
  - An App Runner service to run and manage your containers                       [create complete]  [289.7s]
✔ Deployed service uuid.
Recommended follow-up action:
  - You can access your service at https://sjvxwjua9r.ap-northeast-1.awsapprunner.com over the internet.
```
