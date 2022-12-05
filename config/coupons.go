package config

type Coupons struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultCouponsConfig returns default config values
func DefaultCouponsConfig() Coupons {
	return Coupons{
	HelloWorldMessage:
		"hello from the default config",
	}
}
