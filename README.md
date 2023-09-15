# WUMBO: Warrant &amp; Update Management Block Object Format

WUMBO is designed with security in mind. The authenticity, integrity, and origin of each WUMBO'd packet is cryptographically guaranteed, and that guarantee can be verified by any node in a G-Net network. Additionally, a WUMBO'd packet will fail verification if it's processed out of sequence, which protects against accidental misconfiguration.

WUMBO formatting prevents unauthorized modification of a G-Net network's configuration by linking the integrity of each control packet to the integrity of control packet before it. Additionally, the process of WUMBO-ing a packet requires the packet to be signed by an authorized network administrator. This creates a fully transparent change management process by creating a public, irrevocable link between a network administrator and the configuration changes that they implement.