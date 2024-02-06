# Install dependencies

sudo apt update -y
sudo apt upgrade -y
sudo apt install golang-go gcc sqlite3 wireguard wireguard-tools -y

# Go configuration

go mod download
export CGO_ENABLED=1
go build app.go

# WireGuard configuration

wg genkey | sudo tee /etc/wireguard/private.key
sudo chmod go= /etc/wireguard/private.key
sudo cat /etc/wireguard/private.key | wg pubkey | sudo tee /etc/wireguard/public.key

sudo systemctl enable wg-quick@wg0.service

# Database creation

cat dump.sql | sqlite3 clients.db