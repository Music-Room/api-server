FROM golang:1.10.1

#ENV MIGRATE=${MIGRATE}
#ENV PGHOST=$PGHOST}
#ENV PGPORT=$PGPORT}
#ENV PGNAME=$PGNAME}
#ENV PGUSER=$PGUSER}
#ENV PGPASS=$PGPASS}
#ENV RDSHOST=$RDSHOST}
#ENV RDSPORT=$RDSPORT


# Download and install the latest release of dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

# Set workdir
WORKDIR $GOPATH/src/github.com/music_room

# Dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

# Copy the code from the host
COPY . ./

# Build
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o bin/server .

# Run
#CMD ./bin/server