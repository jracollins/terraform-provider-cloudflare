
	resource "cloudflare_zero_trust_dex_test" "%[1]s" {
		account_id = "%[2]s"
		name = "%[1]s"
		description = "%[1]s"
		interval = "0h30m0s"
		enabled = true
		data = {
  host = "1.1.1.1"
			kind = "traceroute"
}
	}
	