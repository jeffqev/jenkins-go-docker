# Go API Deploy with Jenkins and github Webhooks

# Jenkins Github integration Plugin
* GitHub project: `https://github.com/jeffqev/jenkins-go-docker/`
* Source Code Management: `https://github.com/jeffqev/jenkins-go-docker/`
* Build Triggers: `github hook trigger for gitscm polling (checked)`
* Build shell: `docker-compose up -d --build`

# Github Webhooks settings 
* Payload URL: `http://jenkinsurl/github-webhook/ `
* Content type: `application/json`
