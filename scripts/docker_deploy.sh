#!/bin/bash

### The following section will push your container image somewhere. The `$NAME` variable is provided from our
### Makefile under 'deploy:' rule, which is set to the name of the component/module/service.
###

if [ "$TRAVIS_BRANCH" == "main" ] && [ "$TRAVIS_PULL_REQUEST" == "false" ]; then
    docker tag ${NAME} ${DOCKER_URI}/${NAME}:latest
    rc=$?; if [ $rc -ne 0 ]; then exit $rc; fi

    docker images
    
    # Uncomment this if you want to push docker images somewhere
    # docker push ${DOCKER_URI}/${NAME}:latest
    # rc=$?; if [ $rc -ne 0 ]; then exit $rc; fi
else
    echo "Skipping deploy because not on main branch"
fi