package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a / *b
	*b = *a % *b
	*a = tmp
}