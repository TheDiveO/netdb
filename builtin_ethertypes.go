// Code generated by go generate. DO NOT EDIT.

// Generated from Debian project md/netbase at https://salsa.debian.org
// At 2024-02-03T18:47:44Z
// File etc/ethertypes
// Commit 732e166d6709e2899967b948aa6c63bdcd6b8bd9

package netdb

var BuiltinEtherTypes []EtherType = builtinEtherTypes
var builtinEtherTypes = []EtherType{
		{ Name: "IPv4", Number: 2048, Comment: "IP (IPv4)", Aliases: []string{"ip","ip4",} },
		{ Name: "X25", Number: 2053, Comment: "", Aliases: []string{} },
		{ Name: "ARP", Number: 2054, Comment: "Address Resolution Protocol", Aliases: []string{"ether-arp",} },
		{ Name: "FR_ARP", Number: 2056, Comment: "Frame Relay ARP [RFC1701]", Aliases: []string{} },
		{ Name: "BPQ", Number: 2303, Comment: "G8BPQ AX.25 over Ethernet", Aliases: []string{} },
		{ Name: "TRILL", Number: 8947, Comment: "TRILL [RFC6325]", Aliases: []string{} },
		{ Name: "L2-IS-IS", Number: 8948, Comment: "TRILL IS-IS [RFC6325]", Aliases: []string{} },
		{ Name: "TEB", Number: 25944, Comment: "Transparent Ethernet Bridging [RFC1701]", Aliases: []string{} },
		{ Name: "RAW_FR", Number: 25945, Comment: "Raw Frame Relay [RFC1701]", Aliases: []string{} },
		{ Name: "RARP", Number: 32821, Comment: "Reverse ARP [RFC903]", Aliases: []string{} },
		{ Name: "ATALK", Number: 32923, Comment: "Appletalk", Aliases: []string{} },
		{ Name: "AARP", Number: 33011, Comment: "Appletalk Address Resolution Protocol", Aliases: []string{} },
		{ Name: "802_1Q", Number: 33024, Comment: "VLAN tagged frame [802.1q]", Aliases: []string{"8021q","1q","802.1q","dot1q",} },
		{ Name: "IPX", Number: 33079, Comment: "Novell IPX", Aliases: []string{} },
		{ Name: "NetBEUI", Number: 33169, Comment: "NetBEUI", Aliases: []string{} },
		{ Name: "IPv6", Number: 34525, Comment: "IP version 6", Aliases: []string{"ip6",} },
		{ Name: "PPP", Number: 34827, Comment: "Point-to-Point Protocol", Aliases: []string{} },
		{ Name: "MPLS", Number: 34887, Comment: "MPLS [RFC5332]", Aliases: []string{} },
		{ Name: "MPLS_MULTI", Number: 34888, Comment: "MPLS with upstream-assigned label [RFC5332]", Aliases: []string{} },
		{ Name: "ATMMPOA", Number: 34892, Comment: "MultiProtocol over ATM", Aliases: []string{} },
		{ Name: "PPP_DISC", Number: 34915, Comment: "PPP over Ethernet discovery stage", Aliases: []string{} },
		{ Name: "PPP_SES", Number: 34916, Comment: "PPP over Ethernet session stage", Aliases: []string{} },
		{ Name: "ATMFATE", Number: 34948, Comment: "Frame-based ATM Transport over Ethernet", Aliases: []string{} },
		{ Name: "EAPOL", Number: 34958, Comment: "EAP over LAN [802.1x]", Aliases: []string{} },
		{ Name: "S-TAG", Number: 34984, Comment: "QinQ Service VLAN tag identifier [802.1q]", Aliases: []string{} },
		{ Name: "EAP_PREAUTH", Number: 35015, Comment: "EAPOL Pre-Authentication [802.11i]", Aliases: []string{} },
		{ Name: "LLDP", Number: 35020, Comment: "Link Layer Discovery Protocol [802.1ab]", Aliases: []string{} },
		{ Name: "MACSEC", Number: 35045, Comment: "Media Access Control Security [802.1ae]", Aliases: []string{} },
		{ Name: "PBB", Number: 35047, Comment: "Provider Backbone Bridging [802.1ah]", Aliases: []string{"macinmac",} },
		{ Name: "MVRP", Number: 35061, Comment: "Multiple VLAN Registration Protocol [802.1q]", Aliases: []string{} },
		{ Name: "PTP", Number: 35063, Comment: "Precision Time Protocol", Aliases: []string{} },
		{ Name: "FCOE", Number: 35078, Comment: "Fibre Channel over Ethernet", Aliases: []string{} },
		{ Name: "FIP", Number: 35092, Comment: "FCoE Initialization Protocol", Aliases: []string{} },
		{ Name: "ROCE", Number: 35093, Comment: "RDMA over Converged Ethernet", Aliases: []string{} },
		{ Name: "LoWPAN", Number: 41197, Comment: "LoWPAN encapsulation", Aliases: []string{} },
	}
	