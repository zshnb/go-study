package pointer

import "testing"

func assertBalance(t *testing.T, expected, actual Bitcoin) {
	if expected != actual {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		expected := Bitcoin(10)
		wallet := Wallet{}
		wallet.deposit(10)
		actual := wallet.getBalance()

		assertBalance(t, expected, actual)
	})

	t.Run("withdraw", func(t *testing.T) {
		expected := Bitcoin(10)
		wallet := Wallet{Bitcoin(20)}
		wallet.withdraw(10)
		actual := wallet.getBalance()

		assertBalance(t, expected, actual)
	})

	t.Run("withdraw insufficent funds", func(t *testing.T) {
		expected := Bitcoin(20)
		wallet := Wallet{Bitcoin(20)}
		error := wallet.withdraw(30)
		actual := wallet.getBalance()
		assertBalance(t, expected, actual)
		if error == nil {
			t.Errorf("withdraw get some error")
		}
	})
}
