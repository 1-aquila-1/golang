<h1>Golang - variável de ambiente</h1>

Import necessário

<pre>
import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)
</pre>

O pacote <code>github.com/joho/godotenv</code> é o alvo deste breve estudo sobre variáveis de ambiente em <code>golang</code>.

<div>
Em um projeto as variáveis de ambiente são colocada no arquivo <code>.env</code> ou definida no próprio sistema operacional.
</div>
Então, vamos criar o arquivo <code>.env</code> no projeto.

<h2>1 - Buscando valor de derteminada variável</h2>

O método <code>godotenv.Load()</code> é utilizao para carrega o  arquivo <code>.env</code> para utilizamos as variáveis 
definida no arquivo em questão. 
<div>Exemplo</div>
<pre><code>
if erro := godotenv.Load(".env"); erro != nil {
    t.Fatal(erro)
}
</code></pre>

Com o comando <code>os.Getenv("NOME_VARIAVEL")</code> obtermos o valor da variável definida no argumento do método.
<div>
Exemplo:
</div>
<pre><code>
func TestBuscandoValorDaVariavel(t *testing.T) {
	if erro := godotenv.Load(".env"); erro != nil {
		t.Fatal(erro)
	}
	fmt.Println(os.Getenv("EMAIL_SENHA"))
}
</code></pre>
<p>
Caso a variável não existe o método retorna um valor vazio.
</p>

<h2>2 - Retornando todas as variáveis</h2>

O camando <code>godotenv.Read</code> é utilizado para trazer um <code>map</code> de <code>string</code> que contém as 
variáveis definida no arquivo <code>.env</code>.

<div>Exemplo</div>

<pre><code>
func TestRetornandoVariaveisArquivo(t *testing.T) {
	data, erro := godotenv.Read(".env")
	if erro != nil {
		t.Fatal(erro)
	}
	fmt.Println(data["EMAIL_SENHA"])
}
</code></pre>


