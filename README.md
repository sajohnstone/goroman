# Roman Numerals Conversion - Written in go

## Installing Go

From the termal

```bash
sudo apt-get update
sudo apt-get -y upgrade
curl -O https://storage.googleapis.com/golang/go1.11.2.linux-amd64.tar.gz
tar -xvf go1.11.2.linux-amd64.tar.gz
sudo mv go /usr/local
sudo nano ~/.profile
```

Add this lines at the end

```bash
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Run this

```bash
source ~/.profile
```

## 