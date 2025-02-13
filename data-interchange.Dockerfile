FROM amd64/fedora:latest

WORKDIR /data-interchange

COPY ./ ./

RUN dnf -y upgrade \ 
dnf -y install ufw iptables firewalld nftables \ 
echo 'export PATH="$PATH:/usr/bin"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/bin"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/local/bin"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/local/bin"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/include"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/include"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/local/include"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/local/include"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/bin/go/bin"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/bin/go/bin"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/lib"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/lib"' >> /etc/skel/.bashrc \ 
echo 'export PATH="$PATH:/usr/local/lib"' >> ~/.bashrc && echo 'export PATH="$PATH:/usr/local/lib"' >> /etc/skel/.bashrc \ 
wget https://go.dev/dl/go1.24.0.linux-amd64.tar.gz && tar --extract --file ./*.gz --verbose \ 
mv ./go1.23.4.linux-amd64/go /usr/bin \ 
go env  \ 
go list ./source \ 
go vet ./source \ 
go fix ./source \ 
go fmt ./source \ 
go vet ./source \ 
GOOS=linux GOARCH=amd64 go build -o ./binary/htdinet ./source/main.go
