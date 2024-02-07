import sys, os

def main(client_id: str, ip: str):
    os.system(
        f"wg genkey | tee /etc/wireguard/clients/{client_id}_private.key | wg pubkey > /etc/wireguard/clients/{client_id}_public.key"
    )
    client_private_key = open(f"/etc/wireguard/clients/{client_id}_private.key").read()
    client_public_key = open(f"/etc/wireguard/clients/{client_id}_public.key").read()

    client_id = int(open("/etc/wireguard/clients/last_id").read()) + 1

    with open("/etc/wireguard/wg0.conf", "a") as conf:
        peer = f"[Peer]\nPublicKey = {client_public_key}\nAllowedIPs = 10.8.0.{client_id}"
        conf.write(peer)

    with open("/etc/wireguard/clients/last_id", "") as f:
        f.write(str(client_id))

    server_public_key = open("/etc/wireguard/public.key").read()

    client_conf = f"[Interface]\nAddress = 10.8.0.{client_id}\nPrivateKey = {client_private_key}\nDNS = 1.1.1.1\n\n[Peer]\nPublicKey = {server_public_key}\nEndpoint = {ip}:51820\nAllowedIPs = 0.0.0.0/0, ::/0"
    print(client_conf)

if __name__ == "__main__":
    main(sys.argv[1])