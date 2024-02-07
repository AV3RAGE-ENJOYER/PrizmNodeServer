# Install dependencies

sudo apt update -y
sudo apt upgrade -y
sudo apt install golang-go gcc python3 sqlite3 wireguard -y

# Go configuration

go mod download
export CGO_ENABLED=1
go build app.go

# WireGuard configuration

wg genkey | sudo tee /etc/wireguard/private.key
sudo chmod go= /etc/wireguard/private.key
sudo cat /etc/wireguard/private.key | wg pubkey | sudo tee /etc/wireguard/public.key
export WgPrivateKey=$(cat /etc/wireguard/private.key)
export WgPublicKey=$(cat /etc/wireguard/public.key)
echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.conf
echo -e "[Interface]\nPrivateKey = $WgPrivateKey\nAddress = 10.8.0.1/24\nPostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o ens3 -j MASQUERADE\nPostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o ens3 -j MASQUERADE\nListenPort = 51820\n\n" >> /etc/wireguard/wg0.conf
sudo systemctl enable wg-quick@wg0.service
mkdir /etc/wireguard/clients
echo "1" >> /etc/wireguard/clients/last_id

# Database creation

cat dump.sql | sqlite3 clients.db