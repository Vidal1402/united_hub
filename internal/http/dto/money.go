package dto

import (
	"fmt"
	"strconv"
	"strings"
)

// FormatBRLCents converte um valor em centavos para string no formato "R$ 4.800,00".
func FormatBRLCents(cents int64) string {
	sign := ""
	if cents < 0 {
		sign = "-"
		cents = -cents
	}

	reais := cents / 100
	centavos := cents % 100

	reaisStr := strconv.FormatInt(reais, 10)

	// adiciona separador de milhar com ponto
	var parts []string
	for len(reaisStr) > 3 {
		n := len(reaisStr)
		parts = append([]string{reaisStr[n-3:]}, parts...)
		reaisStr = reaisStr[:n-3]
	}
	parts = append([]string{reaisStr}, parts...)

	intPart := strings.Join(parts, ".")

	return fmt.Sprintf("%sR$ %s,%02d", sign, intPart, centavos)
}

