# Copyright (C) 2021  mieru authors
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

# Assume this file is executed with `docker build -f` from the root dir of mieru project.

FROM golang:1.15-buster

WORKDIR /test

# Copy golang source files into container.
COPY . .

# Create mita user and server config directory.
RUN /usr/sbin/useradd --no-create-home --user-group mita && mkdir -p /etc/mita

RUN go build -ldflags="-X 'github.com/enfein/mieru/pkg/kcp.TestOnlySegmentDropRate=10'" cmd/mieru/mieru.go && \
    go build -ldflags="-X 'github.com/enfein/mieru/pkg/kcp.TestOnlySegmentDropRate=10'" cmd/mita/mita.go && \
    go build test/cmd/sockshttpclient/sockshttpclient.go && \
    go build test/cmd/httpserver/httpserver.go

CMD ["/test/test/deploy/httptest/test.sh"]
