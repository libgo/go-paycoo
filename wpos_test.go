package paycoo

import "testing"

var (
	privateKey = "MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCekCRTUDXvlo6Sgo4MJ0xUgq4uULcEKXz7Ol71pnAHVG3q22EU/NsCdXnFhXwuyID2tYuXhBggwF/o2wSX+qOLLJiFUEUU543NIHq/7pXMlde72ac1ssWaTJ9Iq3sR5iSssU+CcNiqpnGc1K3cF2xsI7Es3it+xHSIkFbkpUcfBkdqn/ENzCoaqmt9R8Q9boVsQs3QXCT/NE86hMuo9U1YgM9ujGW+yUuFRHCPWmk3VKe10fzuc5PwDuaqEL/fxQ3eRQUkGCEKpKbBXhD9PMkAsUh05Mm0Abr399Z+ezWoOO6UUTTb+7HVBlsvJVsd4tc7IigdIv4r4gHw3QpXLHU3AgMBAAECggEAfM6WWSGcoq5OIJI0vuo70ZVfBsOvms4ROyoxzoeoTIyvTBpLqxpwRNo09D1ihux5Xjd93dAysazakU27qui0y/pZPy30gkTUv64eEG76aCNVwN3MGJ9TnGVNRwT7BRayrV+PZuvuDzyd1sUs6D1jsSbS8a8aa1i1JTz/uE9D0O16bS7dcsL2h4W1epjo2jFfwLX7W4otT0o5LcyNCS4lMra8PX+zemXF4nNqg/EVHJcb0s0hhAEBNQpHHoJvv2HGOqOjirYejmDQBUokq9fWuOZ53GUNxrVM8TKm2qjn5zvUjhyfVCNTO5NWbOnggw0at7iHENpcCNyjggo4wTzqKQKBgQD7EeSrPlgyW8Iq+A0Gz/0cQPH5J/kM+fCalSkfaWe47i5+NVrdch1ZWe+l4RyIjwvDKvr6mXkqVPmxo8b3URJ+SaLHzLMKGsTWZ9q2p+RRTs/+Tcy64SI0sEQGjmhsgB4Uwh9cgB6F0msGGzYz1RCd6lS0NtLp492llTgq5WwkCwKBgQChrTmHt3Ya8Ey051wP6COmm1b219UoKW9QvLGoOEqvlO0wHb7D+atnVQa0PhjUtDGvaEes4kDU+mnf6mUqjnEa4Y0QHV4vZ9DGY5ixTfGVV9xnEOGlbWzz6w0V8JnCDesKmsK6fp6NCoJt+XJk+xuDkXNAh28i4WPvbHlNLo7jBQKBgA79lYKn+UogcmpwXw6+EmXCdvAic9tTfQstxsPKXlgLWJlOhDtemeQA268yzqpUqCYiBdQ6n6JoanQxzCpuG8WQjfiR/0qmKISItOVdBuPX6dFMQeISmSE+OQGPNSrtR16D7K1wNOJWmu9FZqoXmaNkH2SXNWajpVNCopTlF9zbAoGAFvCwUAx9Cpd8GtMVIgCrRlGWRlIboaY3cDpF7vuFxWIZQZGG5uG/K1tKJzsEZPHodt4SBXrY/h9F9Iqari4l6E2GqXmxSoKtgejR2p6Xn7wmvs0LZ2CebAG3CpzoBDvbVNbeuH2M27BZD2esjp1+qu1xb/+Himh2uJsonwXzKk0CgYBNTJ9SPNkxgfVgRNAAaVaxgJqwxaf3EgQ4uW7LEj0iHxk4rQUR9KAZyx8zForWh1QiloHHawC5mDxdqrspGURjINngmWBvsNoGI4XQBCrH+pfUZoYn8yNqGAD5aWa+4HLieRbQfYSn2jo6Da5+QauSizdKcsST5rf0TRjhx0lAPQ=="
)

func TestWpos(t *testing.T) {
	client, err := NewClient("d040aeaf23a44dcd", privateKey, true)
	if err != nil {
		t.Error(err)
	}

	param := &Push2cashier{
		MerchantNo:    "10030461",
		StoreNo:       "166130",
		TerminalNo:    "888888888",
		AcceptCashier: "WIP",
		TransType:     "PURCHASE",
		OutOrderNo:    "6666",
		OrderAmount:   "1.20",
	}
	result, err := client.Push2cashier(param)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
