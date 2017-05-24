package main // ➊

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

const linhas3Da43 = `
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
`

const linhaLetraA = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;` // ➌

func TestAnalisarLinha(t *testing.T) { // ➍
	runa, nome, palavras := AnalisarLinha(linhaLetraA) // ➎
	if runa != 'A' {                                   // ➏
		t.Errorf("Esperava 'A', veio %q", runa) // ➐
	}
	const nomeA = "LATIN CAPITAL LETTER A"
	if nome != nomeA {
		t.Errorf("Esperava %q veio %q", nomeA, nome)
	}
	palavrasA := []string{"LATIN", "CAPITAL", "LETTER", "A"}
	if !reflect.DeepEqual(palavras, palavrasA) {
		t.Errorf("Esperado: %q\n\tRecebido: %q", palavrasA, palavras)
	}
}

func TestContém(t *testing.T) {
	casos := []struct {
		fatia     []string
		procurado string
		esperado  bool
	}{
		{[]string{"A", "B"}, "B", true},
		{[]string{}, "A", false},
		{[]string{"A", "B"}, "Z", false},
	}
	for _, caso := range casos {
		obtido := contém(caso.fatia, caso.procurado)
		if obtido != caso.esperado {
			t.Errorf("contém(%#v, %#v) esperado: %v, recebido %v", caso.fatia, caso.procurado, caso.esperado, obtido)
		}
	}
}

func TestContémTodos(t *testing.T) {
	casos := []struct {
		fatia     []string
		procurado []string
		esperado  bool
	}{
		{[]string{"A", "B"}, []string{"B"}, true},
		{[]string{"A", "B", "C"}, []string{"A", "C"}, true},
		{[]string{"A"}, []string{}, true},
		{[]string{}, []string{}, true},
		{[]string{}, []string{"A"}, false},
		{[]string{"A", "B"}, []string{"Z"}, false},
		{[]string{"A", "B", "Z"}, []string{"A", "C"}, false},
		{[]string{"A", "B"}, []string{"A", "B", "C"}, false},
	}
	for _, caso := range casos {
		obtido := contémTodos(caso.fatia, caso.procurado)
		if obtido != caso.esperado {
			t.Errorf("contémTodos(%#v, %#v) esperado: %v, recebido %v", caso.fatia, caso.procurado, caso.esperado, obtido)
		}
	}
}

func ExampleListar() {
	texto := strings.NewReader(linhas3Da43)
	Listar(texto, "MARK")
	// Output: U+003F  ?  QUESTION MARK
}

func ExampleListar_doisResultados() {
	texto := strings.NewReader(linhas3Da43)
	Listar(texto, "SIGN")
	// Output:
	// U+003D  =  EQUALS SIGN
	// U+003E  >  GREATER-THAN SIGN
}

func Example() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"", "cruzeiro"}
	main()
	// Output:
	// U+20A2  ₢  CRUZEIRO SIGN
}
