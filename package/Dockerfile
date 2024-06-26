#
# build the target application
#
FROM public.ecr.aws/docker/library/golang:1.22.3-alpine3.20 as builder

# update the packages
RUN apk update && apk upgrade && apk add --no-cache make npm

# needed for the build
#ARG FONT_AWESOME_TOKEN

WORKDIR /build
COPY go.mod go.sum Makefile ./
COPY backend ./backend
COPY frontend ./frontend
RUN make linux-full

#
# build the target container
#
FROM public.ecr.aws/docker/library/alpine:3.20

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates curl && rm -rf /var/cache/apk/*

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# set the timezone appropriatly
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /dpg-reporting
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME $APP_HOME/bin $APP_HOME/scripts
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY package/data/container_bash_profile /home/webservice/.profile
COPY package/scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY --from=builder /build/bin/dpg-reporting-svc.linux $APP_HOME/bin/dpg-reporting-svc
COPY --from=builder /build/bin/public $APP_HOME/bin/public

# Ensure permissions are correct
RUN chown webservice:webservice /home/webservice/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/dpg-reporting-svc && chmod 755 /home/webservice/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/dpg-reporting-svc
RUN chown -R webservice:webservice $APP_HOME/bin/public && chmod -R 755 $APP_HOME/bin/public

# Add the build tag
ARG BUILD_TAG
RUN test -n "$BUILD_TAG" && touch $APP_HOME/buildtag.build-$BUILD_TAG || touch $APP_HOME/buildtag.build-0

# Specify the user
USER webservice

#
# end of file
#
